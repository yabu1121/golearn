package main

import (
	"fmt"
	"time"
)

/*
【学習ポイント】
1. Goroutine の基本
2. 並行実行
3. WaitGroup
4. 実用例
*/

func main() {
	// ========== 基本的な Goroutine ==========
	fmt.Println("=== 基本的な Goroutine ===")

	// 通常の関数呼び出し（同期）
	sayHello("同期")
	sayHello("実行")

	fmt.Println("\n--- Goroutine で実行 ---")

	// Goroutine で実行（非同期）
	go sayHello("非同期1")
	go sayHello("非同期2")
	go sayHello("非同期3")

	// Goroutine が実行される前にmainが終了しないように待機
	time.Sleep(1 * time.Second)

	// ========== 無名関数の Goroutine ==========
	fmt.Println("\n=== 無名関数の Goroutine ===")

	go func() {
		fmt.Println("無名関数のGoroutine")
	}()

	go func(msg string) {
		fmt.Printf("引数付き: %s\n", msg)
	}("Hello")

	time.Sleep(500 * time.Millisecond)

	// ========== 複数の Goroutine ==========
	fmt.Println("\n=== 複数の Goroutine ===")

	for i := 1; i <= 5; i++ {
		go printNumber(i)
	}

	time.Sleep(1 * time.Second)

	// ========== Goroutine の実用例：並行処理 ==========
	fmt.Println("\n=== 並行処理の例 ===")

	start := time.Now()

	// 同期実行
	fmt.Println("同期実行:")
	task("タスク1", 500)
	task("タスク2", 500)
	task("タスク3", 500)
	fmt.Printf("同期実行時間: %v\n", time.Since(start))

	// 並行実行
	fmt.Println("\n並行実行:")
	start = time.Now()
	go task("タスクA", 500)
	go task("タスクB", 500)
	go task("タスクC", 500)
	time.Sleep(600 * time.Millisecond)
	fmt.Printf("並行実行時間: %v\n", time.Since(start))

	// ========== クロージャと Goroutine ==========
	fmt.Println("\n=== クロージャと Goroutine ===")

	// 間違った例（ループ変数をキャプチャ）
	fmt.Println("間違った例:")
	for i := 1; i <= 3; i++ {
		go func() {
			fmt.Printf("i = %d\n", i) // 全て同じ値になる可能性
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// 正しい例（引数として渡す）
	fmt.Println("\n正しい例:")
	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Printf("n = %d\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// ========== Goroutine のリーク ==========
	fmt.Println("\n=== Goroutine のリーク（注意） ===")

	// これはGoroutineリークの例（実際には避けるべき）
	go func() {
		for {
			// 無限ループ - 終了条件がない
			// time.Sleep(1 * time.Second)
			break // デモのため即座に終了
		}
	}()

	fmt.Println("Goroutineリークに注意！")
}

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printNumber(n int) {
	fmt.Printf("番号: %d\n", n)
}

func task(name string, ms int) {
	fmt.Printf("%s 開始\n", name)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	fmt.Printf("%s 完了\n", name)
}

/*
【実行方法】
go run 01_goroutines.go

【重要な概念】
1. Goroutine は軽量スレッド
2. go キーワードで起動
3. main関数が終了すると全Goroutineも終了
4. 実行順序は保証されない

【Goroutine の特徴】
- 非常に軽量（数KB）
- 数千〜数万のGoroutineを起動可能
- Goランタイムが自動的にスケジューリング

【注意点】
1. main関数が終了するとGoroutineも終了
2. ループ変数のキャプチャに注意
3. Goroutineリークに注意（終了しないGoroutine）
4. 共有メモリへのアクセスは同期が必要

【ベストプラクティス】
1. Goroutineの終了を管理する
2. WaitGroupやChannelで同期
3. ループ変数は引数として渡す
4. context でキャンセル可能にする

【次のステップ】
02_channels.go でChannelを学びましょう
*/
