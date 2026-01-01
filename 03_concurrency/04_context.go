package main

import (
	"context"
	"fmt"
	"time"
)

/*
【学習ポイント】
1. Context の基本
2. キャンセル処理
3. タイムアウト
4. 値の受け渡し
5. 実用例
*/

func main() {
	// ========== 基本的な Context ==========
	fmt.Println("=== 基本的な Context ===")

	// Background context（ルートコンテキスト）
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// TODO context（実装予定の場所で使用）
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)

	// ========== WithCancel（キャンセル可能） ==========
	fmt.Println("\n=== WithCancel ===")

	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1() // 必ず呼び出す

	go func() {
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("Goroutine: キャンセルされました")
				return
			default:
				fmt.Println("Goroutine: 処理中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	cancel1() // キャンセル
	time.Sleep(100 * time.Millisecond)

	// ========== WithTimeout（タイムアウト） ==========
	fmt.Println("\n=== WithTimeout ===")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()

	go longRunningTask(ctx2, "タスク1")

	<-ctx2.Done()
	fmt.Printf("タイムアウト理由: %v\n", ctx2.Err())
	time.Sleep(100 * time.Millisecond)

	// ========== WithDeadline（期限指定） ==========
	fmt.Println("\n=== WithDeadline ===")

	deadline := time.Now().Add(800 * time.Millisecond)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()

	go longRunningTask(ctx3, "タスク2")

	<-ctx3.Done()
	fmt.Printf("期限切れ理由: %v\n", ctx3.Err())
	time.Sleep(100 * time.Millisecond)

	// ========== WithValue（値の受け渡し） ==========
	fmt.Println("\n=== WithValue ===")

	type key string
	const userIDKey key = "userID"
	const requestIDKey key = "requestID"

	ctx4 := context.WithValue(context.Background(), userIDKey, 12345)
	ctx4 = context.WithValue(ctx4, requestIDKey, "req-abc-123")

	processRequest(ctx4)

	// ========== 実用例：HTTPリクエストのタイムアウト ==========
	fmt.Println("\n=== 実用例：HTTPリクエストのタイムアウト ===")

	ctx5, cancel5 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel5()

	result, err := fetchData(ctx5, "https://example.com/api/data")
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("結果: %s\n", result)
	}

	// ========== 実用例：複数のGoroutineの制御 ==========
	fmt.Println("\n=== 実用例：複数のGoroutineの制御 ===")

	ctx6, cancel6 := context.WithCancel(context.Background())

	// 複数のワーカーを起動
	for i := 1; i <= 3; i++ {
		go worker(ctx6, i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("全ワーカーをキャンセル...")
	cancel6() // 全ワーカーをキャンセル
	time.Sleep(500 * time.Millisecond)

	// ========== 実用例：パイプライン処理 ==========
	fmt.Println("\n=== 実用例：パイプライン処理 ===")

	ctx7, cancel7 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel7()

	// データ生成
	numbers := generate(ctx7, 1, 2, 3, 4, 5)

	// 2倍にする
	doubled := multiply(ctx7, numbers, 2)

	// 結果を出力
	for result := range doubled {
		fmt.Printf("結果: %d\n", result)
	}

	// ========== 実用例：データベースクエリ ==========
	fmt.Println("\n=== 実用例：データベースクエリ ===")

	ctx8, cancel8 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel8()

	users, err := queryUsers(ctx8)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("ユーザー数: %d\n", len(users))
	}
}

// ========== ヘルパー関数 ==========

func longRunningTask(ctx context.Context, name string) {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 中断されました（反復 %d）\n", name, i)
			return
		default:
			fmt.Printf("%s: 処理中... %d/10\n", name, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("%s: 完了\n", name)
}

func processRequest(ctx context.Context) {
	type key string
	const userIDKey key = "userID"
	const requestIDKey key = "requestID"

	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)

	fmt.Printf("リクエスト処理: UserID=%v, RequestID=%v\n", userID, requestID)
}

func fetchData(ctx context.Context, url string) (string, error) {
	// データ取得をシミュレート
	resultCh := make(chan string, 1)
	errCh := make(chan error, 1)

	go func() {
		// 実際のHTTPリクエストの代わり
		time.Sleep(500 * time.Millisecond)
		resultCh <- "データ取得成功"
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return "", err
	}
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ワーカー %d: 停止\n", id)
			return
		default:
			fmt.Printf("ワーカー %d: 処理中\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// パイプライン: 数値生成
func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	return out
}

// パイプライン: 乗算
func multiply(ctx context.Context, in <-chan int, factor int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * factor:
			}
		}
	}()
	return out
}

// データベースクエリのシミュレート
func queryUsers(ctx context.Context) ([]string, error) {
	resultCh := make(chan []string, 1)

	go func() {
		// クエリ実行をシミュレート
		time.Sleep(500 * time.Millisecond)
		resultCh <- []string{"太郎", "花子", "次郎"}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case users := <-resultCh:
		return users, nil
	}
}

/*
【実行方法】
go run 04_context.go

【重要な概念】
1. Context はキャンセル、タイムアウト、値の受け渡しに使う
2. 関数の第一引数として渡す慣習
3. 必ず cancel() を呼ぶ（defer で）
4. Done() チャネルで完了を検知

【Context の種類】
context.Background()  // ルートコンテキスト
context.TODO()        // 実装予定の場所で使用
context.WithCancel()  // キャンセル可能
context.WithTimeout() // タイムアウト付き
context.WithDeadline() // 期限付き
context.WithValue()   // 値を持つ

【使用パターン】
func doSomething(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // 処理
    }
}

【ベストプラクティス】
1. 関数の第一引数は ctx context.Context
2. 必ず cancel() を defer で呼ぶ
3. WithValue は最小限に（リクエストスコープの値のみ）
4. nil context は渡さない

【エラーの種類】
context.Canceled       // キャンセルされた
context.DeadlineExceeded // タイムアウト/期限切れ

【実用例】
- HTTPリクエストのタイムアウト
- データベースクエリのキャンセル
- Goroutineの制御
- パイプライン処理の中断

【次のステップ】
04_stdlib/ でファイルIOを学びましょう
*/
