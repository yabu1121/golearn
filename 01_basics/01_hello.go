package main

import "fmt"

/*
【学習ポイント】
1. package main: 実行可能なプログラムは必ずmainパッケージに属する
2. import: 他のパッケージを読み込む
3. func main(): プログラムのエントリーポイント
4. fmt.Println: 標準出力に文字列を出力する関数
*/

func main() {
	// 基本的な出力
	fmt.Println("Hello, Go!")
	
	// 複数の値を出力
	fmt.Println("Go", "is", "awesome!")
	
	// フォーマット付き出力
	name := "太郎"
	age := 25
	fmt.Printf("私の名前は%sで、%d歳です。\n", name, age)
	
	// Sprintfで文字列を作成
	message := fmt.Sprintf("%sさんは%d歳です", name, age)
	fmt.Println(message)
}

/*
【実行方法】
go run 01_hello.go

【出力例】
Hello, Go!
Go is awesome!
私の名前は太郎で、25歳です。
太郎さんは25歳です

【次のステップ】
02_variables.go で変数の宣言方法を学びましょう
*/
