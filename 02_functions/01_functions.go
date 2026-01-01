package main

import "fmt"

/*
【学習ポイント】
1. 関数の基本構文
2. 引数と戻り値
3. 複数の戻り値
4. 名前付き戻り値
5. 可変長引数
*/

func main() {
	// ========== 基本的な関数呼び出し ==========
	fmt.Println("=== 基本的な関数 ===")
	greet()
	greetWithName("太郎")

	// ========== 戻り値のある関数 ==========
	fmt.Println("\n=== 戻り値のある関数 ===")
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	// ========== 複数の戻り値 ==========
	fmt.Println("\n=== 複数の戻り値 ===")
	sum, diff := calculate(10, 3)
	fmt.Printf("和: %d, 差: %d\n", sum, diff)

	// 一部の戻り値を無視
	sum, _ = calculate(20, 5)
	fmt.Printf("和のみ: %d\n", sum)

	// ========== エラーハンドリング ==========
	fmt.Println("\n=== エラーハンドリング ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== 名前付き戻り値 ==========
	fmt.Println("\n=== 名前付き戻り値 ===")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("面積: %d, 周囲: %d\n", area, perimeter)

	// ========== 可変長引数 ==========
	fmt.Println("\n=== 可変長引数 ===")
	fmt.Printf("合計: %d\n", sumAll(1, 2, 3))
	fmt.Printf("合計: %d\n", sumAll(1, 2, 3, 4, 5))

	// スライスを展開して渡す
	numbers := []int{10, 20, 30}
	fmt.Printf("合計: %d\n", sumAll(numbers...))

	// ========== 関数を変数に代入 ==========
	fmt.Println("\n=== 関数を変数に代入 ===")
	operation := add
	fmt.Printf("関数変数: %d\n", operation(5, 3))

	// ========== 関数を引数として渡す ==========
	fmt.Println("\n=== 関数を引数として渡す ===")
	applyOperation(10, 5, add)
	applyOperation(10, 5, multiply)

	// ========== 無名関数 ==========
	fmt.Println("\n=== 無名関数 ===")

	// 無名関数を変数に代入
	double := func(n int) int {
		return n * 2
	}
	fmt.Printf("double(5) = %d\n", double(5))

	// 即座に実行
	func(msg string) {
		fmt.Println("即座に実行:", msg)
	}("Hello!")

	// ========== クロージャ ==========
	fmt.Println("\n=== クロージャ ===")

	counter := makeCounter()
	fmt.Printf("カウント: %d\n", counter()) // 1
	fmt.Printf("カウント: %d\n", counter()) // 2
	fmt.Printf("カウント: %d\n", counter()) // 3

	// 新しいカウンター
	counter2 := makeCounter()
	fmt.Printf("新しいカウンター: %d\n", counter2()) // 1

	// ========== 再帰関数 ==========
	fmt.Println("\n=== 再帰関数 ===")
	fmt.Printf("5の階乗: %d\n", factorial(5))
	fmt.Printf("フィボナッチ(7): %d\n", fibonacci(7))

	// ========== defer ==========
	fmt.Println("\n=== defer ===")
	deferExample()

	// ========== panic と recover ==========
	fmt.Println("\n=== panic と recover ===")
	safeDivide(10, 0)
	fmt.Println("プログラムは続行します")
}

// 引数なし、戻り値なし
func greet() {
	fmt.Println("こんにちは！")
}

// 引数あり、戻り値なし
func greetWithName(name string) {
	fmt.Printf("こんにちは、%sさん！\n", name)
}

// 引数あり、戻り値あり
func add(a, b int) int {
	return a + b
}

// 複数の戻り値
func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// エラーを返す関数
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0で割ることはできません")
	}
	return a / b, nil
}

// 名前付き戻り値
func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 名前付き戻り値は自動的に返される
}

// 可変長引数
func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// 関数を引数として受け取る
func applyOperation(a, b int, op func(int, int) int) {
	result := op(a, b)
	fmt.Printf("結果: %d\n", result)
}

func multiply(a, b int) int {
	return a * b
}

// クロージャ（関数を返す関数）
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 再帰関数：階乗
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 再帰関数：フィボナッチ数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// defer の例
func deferExample() {
	fmt.Println("開始")
	defer fmt.Println("defer 1") // 最後に実行される
	defer fmt.Println("defer 2") // defer は LIFO（後入れ先出し）
	defer fmt.Println("defer 3")
	fmt.Println("終了")
}

// panic と recover の例
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("パニックを回復: %v\n", r)
		}
	}()

	if b == 0 {
		panic("0で割ろうとしました")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

/*
【実行方法】
go run 01_functions.go

【重要な概念】
1. 関数は第一級オブジェクト（変数に代入、引数として渡せる）
2. 複数の戻り値が可能（エラーハンドリングでよく使う）
3. defer は関数終了時に実行される（LIFO順）
4. panic/recover は例外処理に相当

【関数のシグネチャ】
func 関数名(引数名 型, ...) (戻り値の型, ...) {
    // 処理
    return 値, ...
}

【ベストプラクティス】
1. エラーは最後の戻り値として返す
2. 名前付き戻り値は複雑な関数で使う
3. defer はリソースのクリーンアップに使う
4. panic は回復不可能なエラーのみ

【次のステップ】
02_structs.go で構造体を学びましょう
*/
