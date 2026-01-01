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
1. ミドルウェアの概念
2. ミドルウェアの実装
3. ミドルウェアのチェーン
4. 実用的なミドルウェア
*/

func main() {
	// ========== ミドルウェアなし ==========
	http.HandleFunc("/without-middleware", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ミドルウェアなし")
	})

	// ========== ミドルウェア付き ==========
	http.Handle("/with-middleware",
		loggingMiddleware(
			authMiddleware(
				http.HandlerFunc(protectedHandler),
			),
		),
	)

	// ========== API エンドポイント ==========
	http.Handle("/api/data",
		chain(
			loggingMiddleware,
			corsMiddleware,
			authMiddleware,
		)(http.HandlerFunc(dataHandler)),
	)

	// ========== 公開エンドポイント ==========
	http.Handle("/api/public",
		chain(
			loggingMiddleware,
			corsMiddleware,
		)(http.HandlerFunc(publicHandler)),
	)

	fmt.Println("サーバー起動: http://localhost:8080")
	fmt.Println("エンドポイント:")
	fmt.Println("  GET /without-middleware")
	fmt.Println("  GET /with-middleware (要認証)")
	fmt.Println("  GET /api/data (要認証)")
	fmt.Println("  GET /api/public")
	fmt.Println("\n認証ヘッダー例:")
	fmt.Println("  curl -H \"Authorization: Bearer secret-token\" http://localhost:8080/with-middleware")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ========== ミドルウェアの型定義 ==========

type Middleware func(http.Handler) http.Handler

// ========== ロギングミドルウェア ==========

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// リクエスト情報をログ出力
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// 次のハンドラーを実行
		next.ServeHTTP(w, r)

		// 処理時間をログ出力
		log.Printf("処理時間: %v", time.Since(start))
	})
}

// ========== 認証ミドルウェア ==========

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorization ヘッダーをチェック
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "認証が必要です", http.StatusUnauthorized)
			return
		}

		// トークンを検証（実際にはJWT等を使用）
		if token != "Bearer secret-token" {
			http.Error(w, "無効なトークンです", http.StatusUnauthorized)
			return
		}

		log.Println("認証成功")

		// 次のハンドラーを実行
		next.ServeHTTP(w, r)
	})
}

// ========== CORS ミドルウェア ==========

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS ヘッダーを設定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// プリフライトリクエストの処理
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 次のハンドラーを実行
		next.ServeHTTP(w, r)
	})
}

// ========== レート制限ミドルウェア（簡易版） ==========

var requestCount = make(map[string]int)

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		// リクエスト数をカウント
		requestCount[ip]++

		// 10リクエストまで許可
		if requestCount[ip] > 10 {
			http.Error(w, "レート制限を超えました", http.StatusTooManyRequests)
			return
		}

		// 次のハンドラーを実行
		next.ServeHTTP(w, r)
	})
}

// ========== リカバリーミドルウェア ==========

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("パニック発生: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// 次のハンドラーを実行
		next.ServeHTTP(w, r)
	})
}

// ========== ミドルウェアチェーン ==========

func chain(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		// 逆順で適用
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

// ========== ハンドラー関数 ==========

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "認証済みユーザーのみアクセス可能")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "保護されたデータ",
		"user_id": 123,
		"data":    []string{"item1", "item2", "item3"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "公開データ",
		"version": "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

/*
【実行方法】
go run 02_middleware.go

テスト用コマンド:
# ミドルウェアなし
curl http://localhost:8080/without-middleware

# 認証なし（エラー）
curl http://localhost:8080/with-middleware

# 認証あり（成功）
curl -H "Authorization: Bearer secret-token" http://localhost:8080/with-middleware

# API エンドポイント
curl -H "Authorization: Bearer secret-token" http://localhost:8080/api/data

# 公開エンドポイント
curl http://localhost:8080/api/public

【重要な概念】
1. ミドルウェアは http.Handler を受け取り http.Handler を返す
2. 複数のミドルウェアをチェーンできる
3. リクエストの前後で処理を実行できる
4. 認証、ログ、CORS等に使用

【ミドルウェアのパターン】
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // リクエスト前の処理

        next.ServeHTTP(w, r) // 次のハンドラーを実行

        // リクエスト後の処理
    })
}

【実用的なミドルウェア】
1. ロギング: リクエスト情報を記録
2. 認証: トークンを検証
3. CORS: クロスオリジンリクエストを許可
4. レート制限: リクエスト数を制限
5. リカバリー: パニックから回復

【ミドルウェアの順序】
1. リカバリー（最外層）
2. ロギング
3. CORS
4. 認証
5. レート制限
6. ハンドラー（最内層）

【ベストプラクティス】
1. ミドルウェアは小さく保つ
2. 再利用可能にする
3. 順序に注意する
4. エラーハンドリングを適切に行う

【次のステップ】
07_rest_api/ で REST API を学びましょう
*/
