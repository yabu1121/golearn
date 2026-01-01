package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
【学習ポイント】
1. JWT認証の実装
2. バリデーション
3. ページネーション
4. フィルタリング
5. エラーレスポンス
*/

// ========== データモデル ==========

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // JSONに含めない
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// リクエスト/レスポンス型
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type CreatePostRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

type UpdatePostRequest struct {
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
	Published *bool   `json:"published,omitempty"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
}

type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

// ========== インメモリストア ==========

var (
	users = []User{
		{ID: 1, Username: "太郎", Email: "taro@example.com", Password: "password123", CreatedAt: time.Now()},
		{ID: 2, Username: "花子", Email: "hanako@example.com", Password: "password123", CreatedAt: time.Now()},
	}
	posts = []Post{
		{ID: 1, UserID: 1, Title: "最初の投稿", Content: "これは最初の投稿です", Published: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, UserID: 1, Title: "2番目の投稿", Content: "これは2番目の投稿です", Published: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, UserID: 2, Title: "花子の投稿", Content: "花子の投稿内容", Published: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	nextUserID = 3
	nextPostID = 4
)

func main() {
	// ========== ルーティング ==========

	// 認証
	http.HandleFunc("/api/auth/login", loginHandler)
	http.HandleFunc("/api/auth/register", registerHandler)

	// ユーザー
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/users/", userHandler)

	// 投稿
	http.HandleFunc("/api/posts", postsHandler)
	http.HandleFunc("/api/posts/", postHandler)

	// ヘルスチェック
	http.HandleFunc("/health", healthHandler)

	fmt.Println("高度なREST APIサーバー起動: http://localhost:8080")
	fmt.Println("\nエンドポイント:")
	fmt.Println("  POST   /api/auth/login     - ログイン")
	fmt.Println("  POST   /api/auth/register  - ユーザー登録")
	fmt.Println("  GET    /api/users          - ユーザー一覧（ページネーション）")
	fmt.Println("  GET    /api/users/{id}     - ユーザー詳細")
	fmt.Println("  GET    /api/posts          - 投稿一覧（フィルタ、ページネーション）")
	fmt.Println("  POST   /api/posts          - 投稿作成")
	fmt.Println("  GET    /api/posts/{id}     - 投稿詳細")
	fmt.Println("  PUT    /api/posts/{id}     - 投稿更新")
	fmt.Println("  DELETE /api/posts/{id}     - 投稿削除")
	fmt.Println("  GET    /health             - ヘルスチェック")

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux)))
}

// ========== 認証ハンドラー ==========

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	// バリデーション
	if err := validateLoginRequest(req); err != nil {
		respondError(w, "Validation failed", http.StatusBadRequest, err)
		return
	}

	// ユーザー検索
	var foundUser *User
	for i := range users {
		if users[i].Email == req.Email && users[i].Password == req.Password {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil {
		respondError(w, "Invalid credentials", http.StatusUnauthorized, nil)
		return
	}

	// トークン生成（簡易版）
	token := fmt.Sprintf("token-%d-%d", foundUser.ID, time.Now().Unix())

	respondJSON(w, LoginResponse{
		Token: token,
		User:  *foundUser,
	}, http.StatusOK)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	// バリデーション
	if err := validateUser(user); err != nil {
		respondError(w, "Validation failed", http.StatusBadRequest, err)
		return
	}

	// 重複チェック
	for _, u := range users {
		if u.Email == user.Email {
			respondError(w, "Email already exists", http.StatusConflict, nil)
			return
		}
	}

	// ユーザー作成
	user.ID = nextUserID
	nextUserID++
	user.CreatedAt = time.Now()
	users = append(users, user)

	respondJSON(w, user, http.StatusCreated)
}

// ========== ユーザーハンドラー ==========

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	// ページネーション
	page, perPage := getPagination(r)

	start := (page - 1) * perPage
	end := start + perPage

	if start >= len(users) {
		respondJSON(w, PaginatedResponse{
			Data:       []User{},
			Page:       page,
			PerPage:    perPage,
			Total:      len(users),
			TotalPages: (len(users) + perPage - 1) / perPage,
		}, http.StatusOK)
		return
	}

	if end > len(users) {
		end = len(users)
	}

	respondJSON(w, PaginatedResponse{
		Data:       users[start:end],
		Page:       page,
		PerPage:    perPage,
		Total:      len(users),
		TotalPages: (len(users) + perPage - 1) / perPage,
	}, http.StatusOK)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path, "/api/users/")
	if err != nil {
		respondError(w, "Invalid user ID", http.StatusBadRequest, nil)
		return
	}

	if r.Method != http.MethodGet {
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	for _, user := range users {
		if user.ID == id {
			respondJSON(w, user, http.StatusOK)
			return
		}
	}

	respondError(w, "User not found", http.StatusNotFound, nil)
}

// ========== 投稿ハンドラー ==========

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPostsHandler(w, r)
	case http.MethodPost:
		createPostHandler(w, r)
	default:
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
	}
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	// フィルタリング
	userIDStr := r.URL.Query().Get("user_id")
	publishedStr := r.URL.Query().Get("published")

	filtered := posts

	// ユーザーIDでフィルタ
	if userIDStr != "" {
		userID, _ := strconv.Atoi(userIDStr)
		var temp []Post
		for _, p := range filtered {
			if p.UserID == userID {
				temp = append(temp, p)
			}
		}
		filtered = temp
	}

	// 公開状態でフィルタ
	if publishedStr != "" {
		published := publishedStr == "true"
		var temp []Post
		for _, p := range filtered {
			if p.Published == published {
				temp = append(temp, p)
			}
		}
		filtered = temp
	}

	// ページネーション
	page, perPage := getPagination(r)

	start := (page - 1) * perPage
	end := start + perPage

	if start >= len(filtered) {
		respondJSON(w, PaginatedResponse{
			Data:       []Post{},
			Page:       page,
			PerPage:    perPage,
			Total:      len(filtered),
			TotalPages: (len(filtered) + perPage - 1) / perPage,
		}, http.StatusOK)
		return
	}

	if end > len(filtered) {
		end = len(filtered)
	}

	respondJSON(w, PaginatedResponse{
		Data:       filtered[start:end],
		Page:       page,
		PerPage:    perPage,
		Total:      len(filtered),
		TotalPages: (len(filtered) + perPage - 1) / perPage,
	}, http.StatusOK)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	// バリデーション
	if err := validateCreatePost(req); err != nil {
		respondError(w, "Validation failed", http.StatusBadRequest, err)
		return
	}

	post := Post{
		ID:        nextPostID,
		UserID:    1, // 実際は認証から取得
		Title:     req.Title,
		Content:   req.Content,
		Published: req.Published,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	nextPostID++

	posts = append(posts, post)

	respondJSON(w, post, http.StatusCreated)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path, "/api/posts/")
	if err != nil {
		respondError(w, "Invalid post ID", http.StatusBadRequest, nil)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getPostHandler(w, r, id)
	case http.MethodPut:
		updatePostHandler(w, r, id)
	case http.MethodDelete:
		deletePostHandler(w, r, id)
	default:
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
	}
}

func getPostHandler(w http.ResponseWriter, r *http.Request, id int) {
	for _, post := range posts {
		if post.ID == id {
			respondJSON(w, post, http.StatusOK)
			return
		}
	}
	respondError(w, "Post not found", http.StatusNotFound, nil)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request, id int) {
	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	for i := range posts {
		if posts[i].ID == id {
			if req.Title != nil {
				posts[i].Title = *req.Title
			}
			if req.Content != nil {
				posts[i].Content = *req.Content
			}
			if req.Published != nil {
				posts[i].Published = *req.Published
			}
			posts[i].UpdatedAt = time.Now()

			respondJSON(w, posts[i], http.StatusOK)
			return
		}
	}

	respondError(w, "Post not found", http.StatusNotFound, nil)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request, id int) {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	respondError(w, "Post not found", http.StatusNotFound, nil)
}

// ========== ヘルスチェック ==========

func healthHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, map[string]interface{}{
		"status": "ok",
		"time":   time.Now(),
	}, http.StatusOK)
}

// ========== バリデーション ==========

func validateLoginRequest(req LoginRequest) map[string]string {
	errors := make(map[string]string)

	if req.Email == "" {
		errors["email"] = "Email is required"
	}
	if req.Password == "" {
		errors["password"] = "Password is required"
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func validateUser(user User) map[string]string {
	errors := make(map[string]string)

	if user.Username == "" {
		errors["username"] = "Username is required"
	}
	if user.Email == "" {
		errors["email"] = "Email is required"
	}
	if user.Password == "" {
		errors["password"] = "Password is required"
	} else if len(user.Password) < 6 {
		errors["password"] = "Password must be at least 6 characters"
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func validateCreatePost(req CreatePostRequest) map[string]string {
	errors := make(map[string]string)

	if req.Title == "" {
		errors["title"] = "Title is required"
	} else if len(req.Title) > 200 {
		errors["title"] = "Title must be less than 200 characters"
	}

	if req.Content == "" {
		errors["content"] = "Content is required"
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// ========== ヘルパー関数 ==========

func extractID(path, prefix string) (int, error) {
	idStr := strings.TrimPrefix(path, prefix)
	return strconv.Atoi(idStr)
}

func getPagination(r *http.Request) (page, perPage int) {
	page = 1
	perPage = 10

	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if pp := r.URL.Query().Get("per_page"); pp != "" {
		if parsed, err := strconv.Atoi(pp); err == nil && parsed > 0 && parsed <= 100 {
			perPage = parsed
		}
	}

	return
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, message string, status int, details map[string]string) {
	respondJSON(w, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
		Details: details,
	}, status)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

/*
【実行方法】
go run 02_advanced_api.go

【テスト用コマンド】
# ログイン
curl -X POST http://localhost:8080/api/auth/login -d '{"email":"taro@example.com","password":"password123"}' -H "Content-Type: application/json"

# ユーザー一覧（ページネーション）
curl "http://localhost:8080/api/users?page=1&per_page=10"

# 投稿一覧（フィルタ）
curl "http://localhost:8080/api/posts?user_id=1&published=true&page=1"

# 投稿作成
curl -X POST http://localhost:8080/api/posts -d '{"title":"新しい投稿","content":"内容","published":true}' -H "Content-Type: application/json"

# 投稿更新
curl -X PUT http://localhost:8080/api/posts/1 -d '{"title":"更新されたタイトル"}' -H "Content-Type: application/json"

# 投稿削除
curl -X DELETE http://localhost:8080/api/posts/1

【学習ポイント】
1. バリデーション - 入力チェック
2. ページネーション - 大量データの分割
3. フィルタリング - 条件による絞り込み
4. エラーレスポンス - 統一されたエラー形式
5. CORS対応 - クロスオリジンリクエスト

【次のステップ】
実際のプロジェクトでこれらの技術を組み合わせましょう!
*/
