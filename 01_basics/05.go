package main

import "fmt"

func main(){
	// 初期化文付きif（Goの特徴的な機能）
	// ローカルの変数をif分の中で宣言してそれだけでしか利用できない.
	if num := 10; num%2 == 0 {
		fmt.Printf("%dは偶数です\n", num)
	}
	// numはif文の外では使えない

// cではこう書くものを
// for(int i = 0; i < 5; i++){
// 	print("%d", i);
// }

// goでは
// for i := 0; i < 5; i++ {
// 	fmt.Printf("%d\n", i)
// }
// とかく。 ()をつけないのが特徴


// 配列の.length的な機能だと思うけど range というのがある。
// - 三項演算子がない → if-elseを使う

}