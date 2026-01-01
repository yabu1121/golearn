package main

import "fmt"

/*
【学習ポイント】
1. 変数宣言の3つの方法
2. 型推論
3. 定数の宣言
4. ゼロ値の概念
*/

func main() {
	// ========== 変数宣言の方法 ==========

	// 方法1: var キーワードで型を明示
	var name string = "太郎"
	var age int = 25
	fmt.Printf("方法1: %s, %d歳\n", name, age)

	// 方法2: 型推論（型を省略）
	var city = "東京"
	var temperature = 20.5
	fmt.Printf("方法2: %s, %.1f度\n", city, temperature)

	// 方法3: 短縮宣言（関数内でのみ使用可能）
	country := "日本"
	population := 125000000
	fmt.Printf("方法3: %s, %d人\n", country, population)

	// ========== 複数変数の同時宣言 ==========
	var (
		firstName = "太郎"
		lastName  = "山田"
	)
	fmt.Printf("名前: %s %s\n", lastName, firstName)

	// 短縮宣言での複数変数
	x, y, z := 1, 2, 3
	fmt.Printf("座標: (%d, %d, %d)\n", x, y, z)

	// ========== ゼロ値 ==========
	// 変数を宣言だけして初期化しない場合、ゼロ値が入る
	var count int      // 0
	var price float64  // 0.0
	var message string // ""
	var isActive bool  // false
	fmt.Printf("ゼロ値: count=%d, price=%.1f, message='%s', isActive=%t\n",
		count, price, message, isActive)

	// ========== 定数 ==========
	const Pi = 3.14159
	const AppName = "MyApp"
	const MaxUsers = 100

	fmt.Printf("定数: π=%.5f, アプリ名=%s, 最大ユーザー数=%d\n",
		Pi, AppName, MaxUsers)

	// 定数は再代入できない（以下はエラー）
	// Pi = 3.14 // コンパイルエラー!

	// ========== 型変換 ==========
	var integer int = 42
	var floating float64 = float64(integer) // 明示的な型変換が必要
	fmt.Printf("型変換: %d → %.1f\n", integer, floating)

	// ========== 変数の再代入 ==========
	score := 80
	fmt.Printf("初期スコア: %d\n", score)
	score = 95 // 再代入
	fmt.Printf("更新後スコア: %d\n", score)
}

/*
【実行方法】
go run 02_variables.go

【重要な概念】
1. := は関数内でのみ使用可能（パッケージレベルでは var を使う）
2. Goは型推論が強力だが、明示的な型宣言も可能
3. 未使用の変数があるとコンパイルエラーになる
4. 定数は const で宣言し、再代入不可

【次のステップ】
03_types.go でGoの型システムを学びましょう
*/
