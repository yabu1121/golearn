package main

import (
	"fmt"
	"sync"
	"time"
)

/*
【学習ポイント】
1. ファンアウト/ファンイン
2. ワーカープール
3. パイプライン
4. セマフォパターン
5. 実用的な並行処理パターン
*/

func main() {
	// ========== ファンアウト/ファンイン ==========
	fmt.Println("=== ファンアウト/ファンイン ===")

	// データソース
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// ファンアウト: 複数のワーカーで並列処理
	numWorkers := 3
	jobs := make(chan int, len(numbers))
	results := make(chan int, len(numbers))

	// ワーカー起動
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for num := range jobs {
				fmt.Printf("ワーカー %d: %d を処理中\n", id, num)
				time.Sleep(100 * time.Millisecond)
				results <- num * num
			}
		}(i)
	}

	// ジョブを送信
	for _, num := range numbers {
		jobs <- num
	}
	close(jobs)

	// ファンイン: 結果を集約
	go func() {
		wg.Wait()
		close(results)
	}()

	// 結果を収集
	sum := 0
	for result := range results {
		sum += result
	}
	fmt.Printf("合計: %d\n", sum)

	// ========== 動的ワーカープール ==========
	fmt.Println("\n=== 動的ワーカープール ===")

	pool := NewWorkerPool(3)
	pool.Start()

	// タスクを追加
	for i := 1; i <= 10; i++ {
		taskID := i
		pool.AddTask(func() {
			fmt.Printf("タスク %d 実行中\n", taskID)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("タスク %d 完了\n", taskID)
		})
	}

	pool.Stop()
	fmt.Println("全タスク完了")

	// ========== パイプラインパターン ==========
	fmt.Println("\n=== パイプラインパターン ===")

	// ステージ1: 数値生成
	gen := generator(1, 2, 3, 4, 5)

	// ステージ2: 2倍
	doubled := doubler(gen)

	// ステージ3: フィルタ（10以上）
	filtered := filter(doubled, func(n int) bool {
		return n >= 10
	})

	// 結果出力
	for result := range filtered {
		fmt.Printf("結果: %d\n", result)
	}

	// ========== セマフォパターン（同時実行数制限） ==========
	fmt.Println("\n=== セマフォパターン ===")

	maxConcurrent := 2
	sem := make(chan struct{}, maxConcurrent)

	var wg2 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()

			sem <- struct{}{} // セマフォ取得
			fmt.Printf("タスク %d 開始\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("タスク %d 完了\n", id)
			<-sem // セマフォ解放
		}(i)
	}

	wg2.Wait()

	// ========== レート制限パターン ==========
	fmt.Println("\n=== レート制限パターン ===")

	rateLimiter := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		<-rateLimiter // レート制限
		fmt.Printf("リクエスト %d 送信\n", i)
	}

	// ========== タイムアウト付き処理 ==========
	fmt.Println("\n=== タイムアウト付き処理 ===")

	result := make(chan string, 1)

	go func() {
		time.Sleep(500 * time.Millisecond)
		result <- "処理完了"
	}()

	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("タイムアウト")
	}

	// ========== 実用例：並行ダウンロード ==========
	fmt.Println("\n=== 実用例：並行ダウンロード ===")

	urls := []string{
		"https://example.com/file1.txt",
		"https://example.com/file2.txt",
		"https://example.com/file3.txt",
		"https://example.com/file4.txt",
		"https://example.com/file5.txt",
	}

	downloader := NewConcurrentDownloader(3)
	results2 := downloader.DownloadAll(urls)

	for result := range results2 {
		if result.Error != nil {
			fmt.Printf("エラー [%s]: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("成功 [%s]: %s\n", result.URL, result.Data)
		}
	}

	// ========== 実用例：並行マップ処理 ==========
	fmt.Println("\n=== 実用例：並行マップ処理 ===")

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := ParallelMap(input, func(n int) int {
		time.Sleep(50 * time.Millisecond)
		return n * n
	}, 3)

	fmt.Printf("入力: %v\n", input)
	fmt.Printf("出力: %v\n", output)
}

// ========== ワーカープール ==========

type WorkerPool struct {
	maxWorkers int
	tasks      chan func()
	wg         sync.WaitGroup
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		tasks:      make(chan func(), 100),
	}
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		p.wg.Add(1)
		go func(workerID int) {
			defer p.wg.Done()
			for task := range p.tasks {
				task()
			}
		}(i)
	}
}

func (p *WorkerPool) AddTask(task func()) {
	p.tasks <- task
}

func (p *WorkerPool) Stop() {
	close(p.tasks)
	p.wg.Wait()
}

// ========== パイプライン関数 ==========

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return out
}

func doubler(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

func filter(in <-chan int, predicate func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if predicate(n) {
				out <- n
			}
		}
	}()
	return out
}

// ========== 並行ダウンローダー ==========

type DownloadResult struct {
	URL   string
	Data  string
	Error error
}

type ConcurrentDownloader struct {
	maxConcurrent int
}

func NewConcurrentDownloader(maxConcurrent int) *ConcurrentDownloader {
	return &ConcurrentDownloader{maxConcurrent: maxConcurrent}
}

func (d *ConcurrentDownloader) DownloadAll(urls []string) <-chan DownloadResult {
	results := make(chan DownloadResult, len(urls))
	sem := make(chan struct{}, d.maxConcurrent)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			// ダウンロードをシミュレート
			time.Sleep(200 * time.Millisecond)
			results <- DownloadResult{
				URL:  u,
				Data: fmt.Sprintf("データ from %s", u),
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// ========== 並行マップ処理 ==========

func ParallelMap(input []int, fn func(int) int, workers int) []int {
	results := make([]int, len(input))
	jobs := make(chan struct {
		index int
		value int
	}, len(input))

	var wg sync.WaitGroup

	// ワーカー起動
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results[job.index] = fn(job.value)
			}
		}()
	}

	// ジョブ送信
	for i, v := range input {
		jobs <- struct {
			index int
			value int
		}{i, v}
	}
	close(jobs)

	wg.Wait()
	return results
}

/*
【実行方法】
go run 05_patterns.go

【並行処理パターン】

1. ファンアウト/ファンイン
   - 複数のワーカーで並列処理
   - 結果を1つのチャネルに集約

2. ワーカープール
   - 固定数のワーカーでタスクを処理
   - リソースの効率的な利用

3. パイプライン
   - データを段階的に処理
   - 各ステージが並行実行

4. セマフォ
   - 同時実行数を制限
   - リソース保護

5. レート制限
   - 一定間隔で処理
   - API呼び出し等に使用

【ベストプラクティス】
1. バッファ付きチャネルで効率化
2. WaitGroup で完了を待つ
3. defer で確実にリソース解放
4. エラーハンドリングを忘れない

【パフォーマンス】
- ワーカー数は CPU コア数を目安に
- I/O バウンドな処理は多めに
- CPU バウンドな処理は少なめに

【注意点】
1. Goroutine リークに注意
2. チャネルは必ず close
3. デッドロックに注意
4. 競合状態を避ける

【次のステップ】
04_stdlib/03_file_io.go でファイル操作を学びましょう
*/
