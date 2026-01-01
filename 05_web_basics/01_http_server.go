package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
【学習ポイント】
1. HTTP サーバーの基本
2. ハンドラー関数
3. ルーティング
4. リクエストとレスポンス
5. JSON レスポンス
*/

func main() {
	// ========== 基本的なハンドラー ==========
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)

	// ========== JSON レスポンス ==========
	http.HandleFunc("/api/user", userHandler)
	http.HandleFunc("/api/users", usersHandler)

	// ========== POST リクエスト ==========
	http.HandleFunc("/api/create", createHandler)

	// ========== クエリパラメータ ==========
	http.HandleFunc("/api/search", searchHandler)

	// ========== パスパラメータ（簡易版） ==========
	http.HandleFunc("/api/posts/", postHandler)

	fmt.Println("サーバー起動: http://localhost:8080")
	fmt.Println("エンドポイント:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /hello")
	fmt.Println("  GET  /time")
	fmt.Println("  GET  /api/user")
	fmt.Println("  GET  /api/users")
	fmt.Println("  POST /api/create")
	fmt.Println("  GET  /api/search?q=keyword")
	fmt.Println("  GET  /api/posts/1")

	// サーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ========== ハンドラー関数 ==========

// ホームページ
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Go Web Server!</h1>")
	fmt.Fprintf(w, "<p>Method: %s</p>", r.Method)
	fmt.Fprintf(w, "<p>URL: %s</p>", r.URL.Path)
}

// Hello ハンドラー
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// 現在時刻
func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current Time: %s", now)
}

// ========== JSON レスポンス ==========

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 単一ユーザー
func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Name:  "太郎",
		Email: "taro@example.com",
	}

	// JSON レスポンス
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// ユーザー一覧
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "太郎", Email: "taro@example.com"},
		{ID: 2, Name: "花子", Email: "hanako@example.com"},
		{ID: 3, Name: "次郎", Email: "jiro@example.com"},
	}

	resp := Response{
		Success: true,
		Message: "ユーザー一覧取得成功",
		Data:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ========== POST リクエスト ==========

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	// POST メソッドのみ許可
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディをデコード
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// バリデーション
	if req.Name == "" || req.Email == "" {
		resp := Response{
			Success: false,
			Message: "名前とメールアドレスは必須です",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// ユーザー作成（実際にはDBに保存）
	newUser := User{
		ID:    100,
		Name:  req.Name,
		Email: req.Email,
	}

	resp := Response{
		Success: true,
		Message: "ユーザー作成成功",
		Data:    newUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// ========== クエリパラメータ ==========

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	page := r.URL.Query().Get("page")

	if query == "" {
		http.Error(w, "検索キーワードが必要です", http.StatusBadRequest)
		return
	}

	if page == "" {
		page = "1"
	}

	resp := Response{
		Success: true,
		Message: "検索成功",
		Data: map[string]string{
			"query": query,
			"page":  page,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ========== パスパラメータ（簡易版） ==========

func postHandler(w http.ResponseWriter, r *http.Request) {
	// /api/posts/ の後のIDを取得
	id := r.URL.Path[len("/api/posts/"):]

	if id == "" {
		http.Error(w, "投稿IDが必要です", http.StatusBadRequest)
		return
	}

	post := map[string]interface{}{
		"id":    id,
		"title": "サンプル投稿",
		"body":  "これはサンプルの投稿内容です",
	}

	resp := Response{
		Success: true,
		Message: "投稿取得成功",
		Data:    post,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

/*
【実行方法】
go run 01_http_server.go

ブラウザまたはcurlでアクセス:
curl http://localhost:8080/
curl http://localhost:8080/hello?name=太郎
curl http://localhost:8080/api/user
curl http://localhost:8080/api/users
curl -X POST http://localhost:8080/api/create -d '{"name":"太郎","email":"taro@example.com"}' -H "Content-Type: application/json"
curl http://localhost:8080/api/search?q=golang&page=1
curl http://localhost:8080/api/posts/123

【重要な概念】
1. http.HandleFunc でルートを登録
2. http.ListenAndServe でサーバー起動
3. ResponseWriter でレスポンスを書く
4. Request からリクエスト情報を取得

【ハンドラー関数のシグネチャ】
func handler(w http.ResponseWriter, r *http.Request)

【よく使うメソッド】
w.Header().Set(key, value)  // ヘッダー設定
w.WriteHeader(statusCode)   // ステータスコード設定
w.Write([]byte)             // レスポンス書き込み
fmt.Fprintf(w, format, ...)  // フォーマット付き書き込み
json.NewEncoder(w).Encode(v) // JSON エンコード

【リクエスト情報】
r.Method           // HTTPメソッド
r.URL.Path         // パス
r.URL.Query()      // クエリパラメータ
r.Body             // リクエストボディ
r.Header           // ヘッダー

【ステータスコード】
http.StatusOK                  // 200
http.StatusCreated             // 201
http.StatusBadRequest          // 400
http.StatusNotFound            // 404
http.StatusMethodNotAllowed    // 405
http.StatusInternalServerError // 500

【ベストプラクティス】
1. Content-Type ヘッダーを設定
2. エラーハンドリングを適切に行う
3. メソッドをチェックする
4. バリデーションを実装する
5. 構造化されたレスポンスを返す

【次のステップ】
02_middleware.go でミドルウェアを学びましょう
*/
