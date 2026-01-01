package main

import "fmt"

/*
【学習ポイント】
1. ポインタとは
2. ポインタの宣言と使用
3. 値渡しと参照渡し
4. ポインタを使う理由
*/

func main() {
	// ========== ポインタの基本 ==========
	fmt.Println("=== ポインタの基本 ===")

	// 通常の変数
	x := 10
	fmt.Printf("x の値: %d\n", x)
	fmt.Printf("x のアドレス: %p\n", &x) // &演算子でアドレスを取得

	// ポインタ変数の宣言
	var p *int // int型へのポインタ
	p = &x     // xのアドレスをpに代入
	fmt.Printf("p の値（アドレス）: %p\n", p)
	fmt.Printf("p が指す値: %d\n", *p) // *演算子で間接参照

	// ポインタ経由で値を変更
	*p = 20
	fmt.Printf("変更後の x: %d\n", x)

	// ========== ポインタのゼロ値 ==========
	fmt.Println("\n=== ポインタのゼロ値 ===")

	var ptr *int
	fmt.Printf("ptr: %v (nil: %t)\n", ptr, ptr == nil)

	// nilポインタの間接参照はpanic
	// fmt.Println(*ptr) // panic!

	// nilチェック
	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("ポインタはnilです")
	}

	// ========== 値渡し vs ポインタ渡し ==========
	fmt.Println("\n=== 値渡し vs ポインタ渡し ===")

	// 値渡し（コピーが渡される）
	num := 10
	fmt.Printf("変更前: %d\n", num)
	modifyValue(num)
	fmt.Printf("変更後: %d (変わらない)\n", num)

	// ポインタ渡し（アドレスが渡される）
	fmt.Printf("\n変更前: %d\n", num)
	modifyPointer(&num)
	fmt.Printf("変更後: %d (変わる)\n", num)

	// ========== 構造体とポインタ ==========
	fmt.Println("\n=== 構造体とポインタ ===")

	type Person struct {
		Name string
		Age  int
	}

	// 値渡し
	person1 := Person{Name: "太郎", Age: 25}
	fmt.Printf("変更前: %+v\n", person1)
	updatePersonValue(person1)
	fmt.Printf("変更後: %+v (変わらない)\n", person1)

	// ポインタ渡し
	fmt.Printf("\n変更前: %+v\n", person1)
	updatePersonPointer(&person1)
	fmt.Printf("変更後: %+v (変わる)\n", person1)

	// ========== new関数 ==========
	fmt.Println("\n=== new関数 ===")

	// new()はゼロ値で初期化されたポインタを返す
	p2 := new(int)
	fmt.Printf("p2: %p, *p2: %d\n", p2, *p2)
	*p2 = 100
	fmt.Printf("変更後 *p2: %d\n", *p2)

	// 構造体のポインタ
	p3 := new(Person)
	p3.Name = "花子" // (*p3).Name の省略形
	p3.Age = 23
	fmt.Printf("p3: %+v\n", *p3)

	// ========== スライスとマップはポインタ不要 ==========
	fmt.Println("\n=== スライスとマップ ===")

	// スライスは参照型なのでポインタ不要
	slice := []int{1, 2, 3}
	fmt.Printf("変更前: %v\n", slice)
	modifySlice(slice)
	fmt.Printf("変更後: %v (変わる)\n", slice)

	// マップも参照型
	m := map[string]int{"a": 1}
	fmt.Printf("\n変更前: %v\n", m)
	modifyMap(m)
	fmt.Printf("変更後: %v (変わる)\n", m)

	// ========== ポインタの配列 vs 配列のポインタ ==========
	fmt.Println("\n=== ポインタの配列 vs 配列のポインタ ===")

	// ポインタの配列
	a, b, c := 1, 2, 3
	ptrArray := [3]*int{&a, &b, &c}
	fmt.Printf("ポインタの配列: %v\n", ptrArray)
	for i, ptr := range ptrArray {
		fmt.Printf("  [%d]: %d\n", i, *ptr)
	}

	// 配列のポインタ
	arr := [3]int{10, 20, 30}
	arrPtr := &arr
	fmt.Printf("配列のポインタ: %v\n", *arrPtr)
	arrPtr[0] = 100 // (*arrPtr)[0] の省略形
	fmt.Printf("変更後: %v\n", arr)

	// ========== 実用例：関数から複数の値を返す ==========
	fmt.Println("\n=== 実用例：複数の値を返す ===")

	var min, max int
	findMinMax([]int{5, 2, 8, 1, 9}, &min, &max)
	fmt.Printf("最小値: %d, 最大値: %d\n", min, max)

	// ========== ポインタのポインタ ==========
	fmt.Println("\n=== ポインタのポインタ ===")

	value := 42
	ptr1 := &value
	ptr2 := &ptr1

	fmt.Printf("value: %d\n", value)
	fmt.Printf("*ptr1: %d\n", *ptr1)
	fmt.Printf("**ptr2: %d\n", **ptr2)

	**ptr2 = 100
	fmt.Printf("変更後 value: %d\n", value)
}

// 値渡しの関数
func modifyValue(n int) {
	n = 100 // コピーを変更（元の値は変わらない）
}

// ポインタ渡しの関数
func modifyPointer(n *int) {
	*n = 100 // 元の値を変更
}

// 構造体の値渡し
func updatePersonValue(p Person) {
	p.Age = 30 // コピーを変更
}

// 構造体のポインタ渡し
func updatePersonPointer(p *Person) {
	p.Age = 30 // 元の値を変更
}

// スライスの変更
func modifySlice(s []int) {
	s[0] = 100 // スライスは参照型なので元の値が変わる
}

// マップの変更
func modifyMap(m map[string]int) {
	m["a"] = 100 // マップは参照型なので元の値が変わる
}

// 複数の値をポインタで返す例
func findMinMax(nums []int, min *int, max *int) {
	if len(nums) == 0 {
		return
	}
	*min = nums[0]
	*max = nums[0]
	for _, n := range nums {
		if n < *min {
			*min = n
		}
		if n > *max {
			*max = n
		}
	}
}

/*
【実行方法】
go run 08_pointers.go

【重要な概念】
1. ポインタはメモリアドレスを格納する
2. &演算子: アドレスを取得
3. *演算子: 間接参照（ポインタが指す値を取得）
4. nilポインタの間接参照はpanic

【値型 vs 参照型】
値型（コピーされる）:
- int, float, bool, string
- 配列
- 構造体

参照型（参照が渡される）:
- スライス
- マップ
- チャネル

【ポインタを使う理由】
1. 大きな構造体のコピーを避ける（効率的）
2. 関数内で元の値を変更したい
3. nilを表現したい（値がない状態）

【ベストプラクティス】
- 小さな値（int, bool等）: 値渡し
- 大きな構造体: ポインタ渡し
- スライス、マップ: そのまま渡す（ポインタ不要）
- 変更が必要: ポインタ渡し

【次のステップ】
これで基礎編は完了です！
02_functions/ で関数とメソッドを学びましょう
*/
