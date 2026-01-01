package main

import (
	"fmt"
	"time"
)

/*
【学習ポイント】
1. Channel の基本
2. バッファ付き/なし Channel
3. Channel の方向性
4. select 文
5. close と range
*/

func main() {
	// ========== 基本的な Channel ==========
	fmt.Println("=== 基本的な Channel ===")

	// Channel の作成
	ch := make(chan string)

	// Goroutine でデータを送信
	go func() {
		ch <- "Hello" // Channel に送信
	}()

	// Channel からデータを受信
	msg := <-ch
	fmt.Printf("受信: %s\n", msg)

	// ========== バッファなし Channel ==========
	fmt.Println("\n=== バッファなし Channel ===")

	unbuffered := make(chan int)

	go func() {
		fmt.Println("送信開始")
		unbuffered <- 42
		fmt.Println("送信完了")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("受信開始")
	value := <-unbuffered
	fmt.Printf("受信: %d\n", value)

	// ========== バッファ付き Channel ==========
	fmt.Println("\n=== バッファ付き Channel ===")

	buffered := make(chan int, 3) // バッファサイズ3

	// 受信者がいなくても送信可能（バッファまで）
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("3つ送信完了")

	fmt.Printf("受信: %d\n", <-buffered)
	fmt.Printf("受信: %d\n", <-buffered)
	fmt.Printf("受信: %d\n", <-buffered)

	// ========== Channel のクローズ ==========
	fmt.Println("\n=== Channel のクローズ ===")

	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2) // Channel をクローズ

	// クローズ後も受信可能
	fmt.Printf("受信: %d\n", <-ch2)
	fmt.Printf("受信: %d\n", <-ch2)
	fmt.Printf("受信: %d\n", <-ch2)

	// クローズされた Channel から受信するとゼロ値
	val, ok := <-ch2
	fmt.Printf("受信: %d, open: %t\n", val, ok)

	// ========== range で Channel を読む ==========
	fmt.Println("\n=== range で Channel を読む ===")

	ch3 := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch3 <- i
		}
		close(ch3) // 送信完了後にクローズ
	}()

	for num := range ch3 {
		fmt.Printf("受信: %d\n", num)
	}

	// ========== Channel の方向性 ==========
	fmt.Println("\n=== Channel の方向性 ===")

	ch4 := make(chan string)
	go sender(ch4)
	receiver(ch4)

	// ========== select 文 ==========
	fmt.Println("\n=== select 文 ===")

	ch5 := make(chan string)
	ch6 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch5 <- "Channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch6 <- "Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch5:
			fmt.Printf("受信: %s\n", msg1)
		case msg2 := <-ch6:
			fmt.Printf("受信: %s\n", msg2)
		}
	}

	// ========== select with default ==========
	fmt.Println("\n=== select with default ===")

	ch7 := make(chan int)

	select {
	case val := <-ch7:
		fmt.Printf("受信: %d\n", val)
	default:
		fmt.Println("受信できるデータがありません")
	}

	// ========== タイムアウト ==========
	fmt.Println("\n=== タイムアウト ===")

	ch8 := make(chan string)

	select {
	case msg := <-ch8:
		fmt.Printf("受信: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("タイムアウト")
	}

	// ========== 実用例：ワーカープール ==========
	fmt.Println("\n=== ワーカープール ===")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 3つのワーカーを起動
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5つのジョブを送信
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 結果を受信
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("結果: %d\n", result)
	}

	// ========== 実用例：パイプライン ==========
	fmt.Println("\n=== パイプライン ===")

	// ステージ1: 数値を生成
	nums := generate(1, 2, 3, 4, 5)

	// ステージ2: 2倍にする
	doubled := double(nums)

	// ステージ3: 結果を出力
	for result := range doubled {
		fmt.Printf("結果: %d\n", result)
	}
}

// 送信専用 Channel
func sender(ch chan<- string) {
	ch <- "メッセージ"
	close(ch)
}

// 受信専用 Channel
func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Printf("受信: %s\n", msg)
}

// ワーカー関数
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("ワーカー %d: ジョブ %d 開始\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
		fmt.Printf("ワーカー %d: ジョブ %d 完了\n", id, j)
	}
}

// パイプライン: 数値を生成
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// パイプライン: 2倍にする
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

/*
【実行方法】
go run 02_channels.go

【重要な概念】
1. Channel は Goroutine 間の通信手段
2. バッファなし: 送受信が同期
3. バッファ付き: 非同期（バッファまで）
4. close() で Channel をクローズ

【Channel の作成】
ch := make(chan 型)           // バッファなし
ch := make(chan 型, サイズ)    // バッファ付き

【Channel の操作】
ch <- value    // 送信
value := <-ch  // 受信
close(ch)      // クローズ

【Channel の方向性】
chan<- 型  // 送信専用
<-chan 型  // 受信専用

【select 文】
- 複数の Channel を同時に待機
- 準備できた Channel から受信
- default で非ブロッキング

【ベストプラクティス】
1. 送信側が close() する
2. range でループする場合は close() 必須
3. クローズ済み Channel への送信は panic
4. nil Channel への送受信は永久ブロック

【デザインパターン】
1. ワーカープール: 並行処理
2. パイプライン: データの段階的処理
3. ファンアウト/ファンイン: 並列処理と集約

【次のステップ】
03_sync.go で sync パッケージを学びましょう
*/
