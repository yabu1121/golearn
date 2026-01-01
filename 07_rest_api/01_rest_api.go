package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
【学習ポイント】
1. RESTful API の設計
2. CRUD 操作
3. ルーティング
4. エラーハンドリング
5. バリデーション
*/

// ========== データモデル ==========

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// ========== インメモリデータストア ==========

type UserStore struct {
	mu     sync.RWMutex
	users  map[int]*User
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (s *UserStore) Create(name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users[s.nextID] = user
	s.nextID++

	return user, nil
}

func (s *UserStore) GetAll() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *UserStore) GetByID(id int) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("ユーザーが見つかりません")
	}
	return user, nil
}

func (s *UserStore) Update(id int, name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("ユーザーが見つかりません")
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	user.UpdatedAt = time.Now()

	return user, nil
}

func (s *UserStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("ユーザーが見つかりません")
	}

	delete(s.users, id)
	return nil
}

// ========== API ハンドラー ==========

type API struct {
	store *UserStore
}

func NewAPI() *API {
	return &API{
		store: NewUserStore(),
	}
}

// ルーター
func (api *API) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/users", api.usersHandler)
	mux.HandleFunc("/api/users/", api.userHandler)

	return loggingMiddleware(corsMiddleware(mux))
}

// GET /api/users - 全ユーザー取得
// POST /api/users - ユーザー作成
func (api *API) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getAllUsers(w, r)
	case http.MethodPost:
		api.createUser(w, r)
	default:
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /api/users/{id} - ユーザー取得
// PUT /api/users/{id} - ユーザー更新
// DELETE /api/users/{id} - ユーザー削除
func (api *API) userHandler(w http.ResponseWriter, r *http.Request) {
	// IDを抽出
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		api.getUser(w, r, id)
	case http.MethodPut:
		api.updateUser(w, r, id)
	case http.MethodDelete:
		api.deleteUser(w, r, id)
	default:
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// 全ユーザー取得
func (api *API) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := api.store.GetAll()
	respondJSON(w, Response{
		Success: true,
		Data:    users,
	}, http.StatusOK)
}

// ユーザー作成
func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// バリデーション
	if req.Name == "" || req.Email == "" {
		respondError(w, "名前とメールアドレスは必須です", http.StatusBadRequest)
		return
	}

	user, err := api.store.Create(req.Name, req.Email)
	if err != nil {
		respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, Response{
		Success: true,
		Message: "ユーザー作成成功",
		Data:    user,
	}, http.StatusCreated)
}

// ユーザー取得
func (api *API) getUser(w http.ResponseWriter, r *http.Request, id int) {
	user, err := api.store.GetByID(id)
	if err != nil {
		respondError(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJSON(w, Response{
		Success: true,
		Data:    user,
	}, http.StatusOK)
}

// ユーザー更新
func (api *API) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := api.store.Update(id, req.Name, req.Email)
	if err != nil {
		respondError(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJSON(w, Response{
		Success: true,
		Message: "ユーザー更新成功",
		Data:    user,
	}, http.StatusOK)
}

// ユーザー削除
func (api *API) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	err := api.store.Delete(id)
	if err != nil {
		respondError(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJSON(w, Response{
		Success: true,
		Message: "ユーザー削除成功",
	}, http.StatusOK)
}

// ========== ヘルパー関数 ==========

func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		return 0, fmt.Errorf("invalid path")
	}
	return strconv.Atoi(parts[3])
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, Response{
		Success: false,
		Error:   message,
	}, status)
}

// ========== ミドルウェア ==========

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("処理時間: %v", time.Since(start))
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ========== メイン関数 ==========

func main() {
	api := NewAPI()

	// サンプルデータを追加
	api.store.Create("太郎", "taro@example.com")
	api.store.Create("花子", "hanako@example.com")
	api.store.Create("次郎", "jiro@example.com")

	fmt.Println("REST API サーバー起動: http://localhost:8080")
	fmt.Println("\nエンドポイント:")
	fmt.Println("  GET    /api/users      - 全ユーザー取得")
	fmt.Println("  POST   /api/users      - ユーザー作成")
	fmt.Println("  GET    /api/users/{id} - ユーザー取得")
	fmt.Println("  PUT    /api/users/{id} - ユーザー更新")
	fmt.Println("  DELETE /api/users/{id} - ユーザー削除")
	fmt.Println("\n使用例:")
	fmt.Println("  curl http://localhost:8080/api/users")
	fmt.Println("  curl -X POST http://localhost:8080/api/users -d '{\"name\":\"四郎\",\"email\":\"shiro@example.com\"}' -H 'Content-Type: application/json'")
	fmt.Println("  curl http://localhost:8080/api/users/1")
	fmt.Println("  curl -X PUT http://localhost:8080/api/users/1 -d '{\"name\":\"太郎2\"}' -H 'Content-Type: application/json'")
	fmt.Println("  curl -X DELETE http://localhost:8080/api/users/1")

	log.Fatal(http.ListenAndServe(":8080", api.Router()))
}

/*
【実行方法】
go run 01_rest_api.go

【REST API の原則】
1. リソース指向（/api/users）
2. HTTPメソッドで操作を表現
3. ステータスコードで結果を返す
4. JSON形式でデータをやり取り

【HTTPメソッドとCRUD】
- POST:   Create（作成）
- GET:    Read（読み取り）
- PUT:    Update（更新）
- DELETE: Delete（削除）

【ステータスコード】
- 200 OK: 成功
- 201 Created: 作成成功
- 400 Bad Request: 不正なリクエスト
- 404 Not Found: リソースが見つからない
- 500 Internal Server Error: サーバーエラー

【ベストプラクティス】
1. 一貫性のあるURL設計
2. 適切なHTTPメソッドを使用
3. 適切なステータスコードを返す
4. バリデーションを実装
5. エラーメッセージを明確に

【次のステップ】
08_project/ で完全なWebアプリケーションを学びましょう
*/
