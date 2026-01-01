package main

// 不要なインポートがあるともしかしたら実行できない？？
import (
	"fmt"
)

func main(){
	var num int = 1
	var str string = "太郎"
	fmt.Println(num, str)

	// 静的型つけ
	var name string = "吉田"
	var old int = 25

	// 型推論
	// var name = "吉田"

	// 省略
	// name := "吉田"

	// fmt.Sprint() s => ストリングビルダー？？
	// fmt.Fprint() f => file操作？？
	var message string = fmt.Sprintf("%v、%v歳、独身", name, old)
	fmt.Printf(message)

	fmt.Println()
	// 同時宣言

	// var x, y, z int = 1, 2, 3
	
	// 省略ならこうでも書ける
	x, y, z  := 1, 2, 3 
	
	// var (
	// 	x int = 1
	// 	y int = 2
	// 	z int = 3
	// )
	fmt.Printf("x, y, z = ( %d, %d, %d )", x, y, z)

}
