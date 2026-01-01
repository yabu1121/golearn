// 主な実行関数はmain
package main

// fmtライブラリは汎用ライブラリ。
// print, println, printf
// sprint, sprintln, sprintf
// fprint, fprintln, fprintfの9つがある。
// scan, errorもある。
import "fmt"

// エントリーポイント関数
func main(){
	// これは空白が表示されるくらい、改行があるわけではないから気を付ける。
	fmt.Print("hello", "the", "world");
	fmt.Println("hello", "the", "world");

	// 代入は　:= で記録する。(型を自動でつけてくれる？？)
	name := "太郎"
	// % ~~ は　書式指定子　いろいろあるので確認してください。
	fmt.Printf("%v", name)

// Hello, Go!
// Go is awesome!
// 私の名前は太郎で、25歳です。
// 太郎さんは25歳です
	fmt.Println("Hello, Go!")
	fmt.Println("Go", "is", "awesome!")
	my_name := "太郎" 
	my_yo := 25
	fmt.Printf("私の名前は%vで、%v歳です。", my_name, my_yo)
	fmt.Println()
	message := fmt.Sprintf("%vさんは%v歳です", my_name, my_yo)
	fmt.Printf(message)
}