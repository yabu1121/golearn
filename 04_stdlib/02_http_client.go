package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/*
【学習ポイント】
1. HTTP GET リクエスト
2. HTTP POST リクエスト
3. カスタムヘッダー
4. タイムアウト
5. エラーハンドリング
*/

func main() {
	// ========== 基本的な GET リクエスト ==========
	fmt.Println("=== 基本的な GET リクエスト ===")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Printf("エラー: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("ステータスコード: %d\n", resp.StatusCode)
	fmt.Printf("ステータス: %s\n", resp.Status)

	// レスポンスボディを読む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("レスポンス: %s\n", string(body))

	// ========== JSON レスポンスのデコード ==========
	fmt.Println("\n=== JSON レスポンスのデコード ===")

	resp2, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Printf("エラー: %v\n", err)
		return
	}
	defer resp2.Body.Close()

	var user User
	err = json.NewDecoder(resp2.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ユーザー: %+v\n", user)

	// ========== カスタムリクエスト ==========
	fmt.Println("\n=== カスタムリクエスト ===")

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		log.Fatal(err)
	}

	// ヘッダーを追加
	req.Header.Set("User-Agent", "Go-HTTP-Client")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp3, err := client.Do(req)
	if err != nil {
		log.Printf("エラー: %v\n", err)
		return
	}
	defer resp3.Body.Close()

	fmt.Printf("ステータス: %s\n", resp3.Status)

	// ========== タイムアウト設定 ==========
	fmt.Println("\n=== タイムアウト設定 ===")

	clientWithTimeout := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp4, err := clientWithTimeout.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Printf("エラー: %v\n", err)
		return
	}
	defer resp4.Body.Close()

	fmt.Printf("ステータス: %s\n", resp4.Status)

	// ========== クエリパラメータ ==========
	fmt.Println("\n=== クエリパラメータ ===")

	req2, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)

	// クエリパラメータを追加
	q := req2.URL.Query()
	q.Add("userId", "1")
	req2.URL.RawQuery = q.Encode()

	fmt.Printf("URL: %s\n", req2.URL.String())

	resp5, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Printf("エラー: %v\n", err)
		return
	}
	defer resp5.Body.Close()

	var posts []Post
	json.NewDecoder(resp5.Body).Decode(&posts)
	fmt.Printf("投稿数: %d\n", len(posts))

	// ========== ステータスコードのチェック ==========
	fmt.Println("\n=== ステータスコードのチェック ===")

	resp6, err := http.Get("https://jsonplaceholder.typicode.com/posts/999999")
	if err != nil {
		log.Printf("エラー: %v\n", err)
		return
	}
	defer resp6.Body.Close()

	if resp6.StatusCode != http.StatusOK {
		fmt.Printf("エラー: ステータスコード %d\n", resp6.StatusCode)
	} else {
		fmt.Println("成功")
	}

	fmt.Println("\n=== 実用例：API クライアント ===")

	apiClient := NewAPIClient("https://jsonplaceholder.typicode.com")

	// ユーザー取得
	user2, err := apiClient.GetUser(1)
	if err != nil {
		log.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("ユーザー: %s (%s)\n", user2.Name, user2.Email)
	}

	// 投稿一覧取得
	posts2, err := apiClient.GetPosts(1)
	if err != nil {
		log.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("投稿数: %d\n", len(posts2))
		if len(posts2) > 0 {
			fmt.Printf("最初の投稿: %s\n", posts2[0].Title)
		}
	}
}

// ========== 構造体の定義 ==========

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ========== API クライアント ==========

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *APIClient) GetUser(id int) (*User, error) {
	url := fmt.Sprintf("%s/users/%d", c.BaseURL, id)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ステータスコード: %d", resp.StatusCode)
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *APIClient) GetPosts(userID int) ([]Post, error) {
	url := fmt.Sprintf("%s/posts?userId=%d", c.BaseURL, userID)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ステータスコード: %d", resp.StatusCode)
	}

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

/*
【実行方法】
go run 02_http_client.go

【重要な概念】
1. http.Get: 簡単な GET リクエスト
2. http.NewRequest: カスタムリクエスト
3. http.Client: クライアント設定
4. defer resp.Body.Close(): リソース解放

【基本的な流れ】
1. リクエストを作成
2. クライアントで送信
3. レスポンスを受信
4. ボディを読む
5. Close() でリソース解放

【よく使う関数】
http.Get(url)                    // GET リクエスト
http.Post(url, contentType, body) // POST リクエスト
http.NewRequest(method, url, body) // カスタムリクエスト
client.Do(req)                   // リクエスト送信

【ステータスコード】
http.StatusOK           // 200
http.StatusCreated      // 201
http.StatusBadRequest   // 400
http.StatusNotFound     // 404
http.StatusInternalServerError // 500

【ベストプラクティス】
1. defer resp.Body.Close() を必ず書く
2. タイムアウトを設定する
3. ステータスコードをチェックする
4. エラーハンドリングを適切に行う
5. 構造体でクライアントをラップする

【次のステップ】
05_web_basics/ で HTTP サーバーを学びましょう
*/
