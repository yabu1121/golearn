package main

import (
	"fmt"
	"sync"
	"time"
)

/*
【学習ポイント】
1. WaitGroup
2. Mutex
3. RWMutex
4. Once
5. 実用例
*/

func main() {
	// ========== WaitGroup ==========
	fmt.Println("=== WaitGroup ===")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // カウンターを増やす
		go func(id int) {
			defer wg.Done() // 完了時にカウンターを減らす
			fmt.Printf("Goroutine %d 開始\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d 完了\n", id)
		}(i)
	}

	wg.Wait() // 全Goroutineの完了を待つ
	fmt.Println("全て完了")

	// ========== Mutex（排他制御） ==========
	fmt.Println("\n=== Mutex ===")

	counter := &Counter{value: 0}
	var wg2 sync.WaitGroup

	// 100個のGoroutineが同時にインクリメント
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter.Increment()
		}()
	}

	wg2.Wait()
	fmt.Printf("最終カウント: %d\n", counter.Value())

	// ========== Mutex なしの問題 ==========
	fmt.Println("\n=== Mutex なしの問題 ===")

	unsafeCounter := &UnsafeCounter{value: 0}
	var wg3 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			unsafeCounter.Increment()
		}()
	}

	wg3.Wait()
	fmt.Printf("Mutexなしのカウント: %d (100にならない可能性)\n", unsafeCounter.Value())

	// ========== RWMutex（読み書きロック） ==========
	fmt.Println("\n=== RWMutex ===")

	cache := &Cache{data: make(map[string]string)}
	var wg4 sync.WaitGroup

	// 書き込み
	wg4.Add(1)
	go func() {
		defer wg4.Done()
		cache.Set("key1", "value1")
		cache.Set("key2", "value2")
	}()

	// 複数の読み込み（並行実行可能）
	for i := 0; i < 5; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			time.Sleep(10 * time.Millisecond)
			val := cache.Get("key1")
			fmt.Printf("読み込み %d: %s\n", id, val)
		}(i)
	}

	wg4.Wait()

	// ========== sync.Once ==========
	fmt.Println("\n=== sync.Once ===")

	config := &Config{}
	var wg5 sync.WaitGroup

	// 複数のGoroutineから初期化を試みる
	for i := 0; i < 5; i++ {
		wg5.Add(1)
		go func(id int) {
			defer wg5.Done()
			config.Initialize(id)
		}(i)
	}

	wg5.Wait()
	fmt.Printf("設定値: %s\n", config.Value())

	// ========== 実用例：並行ダウンロード ==========
	fmt.Println("\n=== 実用例：並行ダウンロード ===")

	urls := []string{
		"https://example.com/file1",
		"https://example.com/file2",
		"https://example.com/file3",
	}

	downloader := &Downloader{}
	downloader.DownloadAll(urls)

	// ========== 実用例：並行マップ処理 ==========
	fmt.Println("\n=== 実用例：並行マップ処理 ===")

	numbers := []int{1, 2, 3, 4, 5}
	results := parallelMap(numbers, func(n int) int {
		time.Sleep(100 * time.Millisecond)
		return n * n
	})

	fmt.Printf("結果: %v\n", results)
}

// ========== Mutex を使った安全なカウンター ==========

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// ========== Mutex なしの安全でないカウンター ==========

type UnsafeCounter struct {
	value int
}

func (c *UnsafeCounter) Increment() {
	c.value++ // 競合状態（race condition）
}

func (c *UnsafeCounter) Value() int {
	return c.value
}

// ========== RWMutex を使ったキャッシュ ==========

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	fmt.Printf("書き込み: %s = %s\n", key, value)
}

func (c *Cache) Get(key string) string {
	c.mu.RLock() // 読み込みロック（複数同時可能）
	defer c.mu.RUnlock()
	return c.data[key]
}

// ========== sync.Once を使った初期化 ==========

type Config struct {
	once  sync.Once
	value string
}

func (c *Config) Initialize(id int) {
	c.once.Do(func() {
		fmt.Printf("初期化実行（ID: %d）\n", id)
		time.Sleep(100 * time.Millisecond)
		c.value = fmt.Sprintf("initialized by %d", id)
	})
}

func (c *Config) Value() string {
	return c.value
}

// ========== 並行ダウンロード ==========

type Downloader struct{}

func (d *Downloader) DownloadAll(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			d.download(u)
		}(url)
	}

	wg.Wait()
	fmt.Println("全ダウンロード完了")
}

func (d *Downloader) download(url string) {
	fmt.Printf("ダウンロード開始: %s\n", url)
	time.Sleep(200 * time.Millisecond) // ダウンロードをシミュレート
	fmt.Printf("ダウンロード完了: %s\n", url)
}

// ========== 並行マップ処理 ==========

func parallelMap(nums []int, fn func(int) int) []int {
	results := make([]int, len(nums))
	var wg sync.WaitGroup

	for i, num := range nums {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			results[index] = fn(value)
		}(i, num)
	}

	wg.Wait()
	return results
}

/*
【実行方法】
go run 03_sync.go

【重要な概念】
1. WaitGroup: Goroutineの完了を待つ
2. Mutex: 排他制御（1つずつアクセス）
3. RWMutex: 読み書きロック
4. Once: 1回だけ実行

【sync.WaitGroup】
wg.Add(1)   // カウンター増加
wg.Done()   // カウンター減少
wg.Wait()   // カウンターが0になるまで待機

【sync.Mutex】
mu.Lock()   // ロック取得
mu.Unlock() // ロック解放

【sync.RWMutex】
mu.RLock()   // 読み込みロック
mu.RUnlock() // 読み込みロック解放
mu.Lock()    // 書き込みロック
mu.Unlock()  // 書き込みロック解放

【sync.Once】
once.Do(func() { ... }) // 1回だけ実行

【競合状態（Race Condition）】
複数のGoroutineが同時に同じメモリにアクセスすると、
予期しない結果になる可能性がある。

【検出方法】
go run -race program.go

【ベストプラクティス】
1. defer で Unlock を確実に実行
2. ロックの範囲は最小限に
3. デッドロックに注意
4. Channel を優先的に検討

【次のステップ】
04_stdlib/ で標準ライブラリを学びましょう
*/
