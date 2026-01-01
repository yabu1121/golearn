package main

import "fmt"

/*
【学習ポイント】
1. 配列（固定長）
2. スライス（可変長）
3. スライスの操作
4. make関数
5. append関数
*/

func main() {
	// ========== 配列（固定長） ==========
	fmt.Println("=== 配列 ===")

	// 配列の宣言と初期化
	var arr1 [5]int // ゼロ値で初期化: [0 0 0 0 0]
	fmt.Printf("arr1: %v\n", arr1)

	// 値を指定して初期化
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2)

	// 一部だけ初期化
	arr3 := [5]int{1, 2} // 残りは0
	fmt.Printf("arr3: %v\n", arr3)

	// 長さを自動推論
	arr4 := [...]int{10, 20, 30}
	fmt.Printf("arr4: %v (長さ: %d)\n", arr4, len(arr4))

	// 配列の要素にアクセス
	fmt.Printf("arr2[0]: %d\n", arr2[0])
	arr2[0] = 100
	fmt.Printf("変更後 arr2: %v\n", arr2)

	// 配列のループ
	fmt.Println("\n配列のループ:")
	for i, v := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", i, v)
	}

	// ========== スライス（可変長） ==========
	fmt.Println("\n=== スライス ===")

	// スライスの宣言（配列と違い、長さを指定しない）
	var slice1 []int
	fmt.Printf("slice1: %v (長さ: %d, 容量: %d)\n", slice1, len(slice1), cap(slice1))

	// リテラルで初期化
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice2: %v (長さ: %d, 容量: %d)\n", slice2, len(slice2), cap(slice2))

	// make関数でスライスを作成
	slice3 := make([]int, 3)    // 長さ3、容量3
	slice4 := make([]int, 3, 5) // 長さ3、容量5
	fmt.Printf("slice3: %v (長さ: %d, 容量: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %v (長さ: %d, 容量: %d)\n", slice4, len(slice4), cap(slice4))

	// ========== スライスの操作 ==========
	fmt.Println("\n=== スライスの操作 ===")

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("元のスライス: %v\n", numbers)

	// スライス演算子 [start:end]
	fmt.Printf("numbers[2:5]: %v\n", numbers[2:5]) // [2 3 4]
	fmt.Printf("numbers[:3]:  %v\n", numbers[:3])  // [0 1 2]
	fmt.Printf("numbers[7:]:  %v\n", numbers[7:])  // [7 8 9]
	fmt.Printf("numbers[:]:   %v\n", numbers[:])   // 全体

	// ========== append（要素の追加） ==========
	fmt.Println("\n=== append ===")

	fruits := []string{"りんご", "バナナ"}
	fmt.Printf("初期: %v (長さ: %d, 容量: %d)\n", fruits, len(fruits), cap(fruits))

	fruits = append(fruits, "オレンジ")
	fmt.Printf("追加後: %v (長さ: %d, 容量: %d)\n", fruits, len(fruits), cap(fruits))

	// 複数要素を追加
	fruits = append(fruits, "ぶどう", "メロン")
	fmt.Printf("複数追加: %v (長さ: %d, 容量: %d)\n", fruits, len(fruits), cap(fruits))

	// スライス同士を結合
	moreFruits := []string{"いちご", "もも"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("結合後: %v\n", fruits)

	// ========== copy（スライスのコピー） ==========
	fmt.Println("\n=== copy ===")

	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)

	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)

	// dstを変更してもsrcは影響を受けない
	dst[0] = 100
	fmt.Printf("変更後 src: %v\n", src)
	fmt.Printf("変更後 dst: %v\n", dst)

	// ========== スライスの注意点 ==========
	fmt.Println("\n=== スライスの参照 ===")

	original := []int{1, 2, 3, 4, 5}
	reference := original // 参照のコピー（同じ配列を指す）

	reference[0] = 100
	fmt.Printf("original:  %v\n", original)  // [100 2 3 4 5]
	fmt.Printf("reference: %v\n", reference) // [100 2 3 4 5]

	// ========== 多次元スライス ==========
	fmt.Println("\n=== 多次元スライス ===")

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d ", i, j, val)
		}
		fmt.Println()
	}

	// ========== 実用例：スライスの削除 ==========
	fmt.Println("\n=== 要素の削除 ===")

	items := []string{"A", "B", "C", "D", "E"}
	fmt.Printf("元: %v\n", items)

	// インデックス2の要素を削除
	index := 2
	items = append(items[:index], items[index+1:]...)
	fmt.Printf("削除後: %v\n", items)

	// ========== 実用例：スライスの挿入 ==========
	fmt.Println("\n=== 要素の挿入 ===")

	nums := []int{1, 2, 4, 5}
	fmt.Printf("元: %v\n", nums)

	// インデックス2に3を挿入
	insertIndex := 2
	insertValue := 3
	nums = append(nums[:insertIndex], append([]int{insertValue}, nums[insertIndex:]...)...)
	fmt.Printf("挿入後: %v\n", nums)
}

/*
【実行方法】
go run 06_arrays_slices.go

【配列 vs スライス】
配列:
- 固定長
- 値型（コピーされる）
- [5]int と [10]int は異なる型

スライス:
- 可変長
- 参照型（同じ配列を指す）
- 最もよく使われる

【重要な概念】
1. スライスは内部的に配列への参照を持つ
2. len(): 現在の要素数
3. cap(): 容量（再割り当てなしで追加できる要素数）
4. append()は新しいスライスを返す（元のスライスは変更されない）

【ベストプラクティス】
- 通常はスライスを使う（配列はほとんど使わない）
- 容量が分かっている場合はmakeで事前確保
- スライスのコピーが必要な場合はcopyを使う

【次のステップ】
07_maps.go でマップ（連想配列）を学びましょう
*/
