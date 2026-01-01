package main

import "fmt"

/*
【学習ポイント】
1. if文（条件分岐）
2. for文（ループ）
3. switch文（多分岐）
4. break, continue
*/

func main() {
	// ========== if文 ==========
	fmt.Println("=== if文 ===")
	age := 20

	if age >= 20 {
		fmt.Println("成人です")
	}

	// if-else
	score := 75
	if score >= 80 {
		fmt.Println("優秀です")
	} else {
		fmt.Println("もう少し頑張りましょう")
	}

	// if-else if-else
	temperature := 25
	if temperature >= 30 {
		fmt.Println("暑い")
	} else if temperature >= 20 {
		fmt.Println("快適")
	} else {
		fmt.Println("寒い")
	}

	// 初期化文付きif（Goの特徴的な機能）
	if num := 10; num%2 == 0 {
		fmt.Printf("%dは偶数です\n", num)
	}
	// numはif文の外では使えない

	// ========== for文（基本形） ==========
	fmt.Println("\n=== for文（基本形） ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while風のfor（条件のみ）
	fmt.Println("\n=== while風のfor ===")
	count := 0
	for count < 3 {
		fmt.Printf("count: %d\n", count)
		count++
	}

	// 無限ループ
	fmt.Println("\n=== 無限ループ（breakで抜ける） ===")
	n := 0
	for {
		if n >= 3 {
			break // ループを抜ける
		}
		fmt.Printf("n: %d\n", n)
		n++
	}

	// continue（次の反復へ）
	fmt.Println("\n=== continue ===")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // i==2の時はスキップ
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ========== range（配列・スライスの反復） ==========
	fmt.Println("\n=== range（スライス） ===")
	fruits := []string{"りんご", "バナナ", "オレンジ"}

	// インデックスと値の両方
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// 値のみ（インデックスを_で無視）
	fmt.Println("\n値のみ:")
	for _, fruit := range fruits {
		fmt.Printf("- %s\n", fruit)
	}

	// インデックスのみ
	fmt.Println("\nインデックスのみ:")
	for i := range fruits {
		fmt.Printf("インデックス: %d\n", i)
	}

	// ========== range（マップ） ==========
	fmt.Println("\n=== range（マップ） ===")
	scores := map[string]int{
		"太郎": 85,
		"花子": 92,
		"次郎": 78,
	}

	for name, score := range scores {
		fmt.Printf("%s: %d点\n", name, score)
	}

	// ========== switch文 ==========
	fmt.Println("\n=== switch文 ===")
	day := "月曜日"

	switch day {
	case "月曜日":
		fmt.Println("週の始まりです")
	case "金曜日":
		fmt.Println("週末が近いです")
	case "土曜日", "日曜日":
		fmt.Println("週末です！")
	default:
		fmt.Println("平日です")
	}

	// 条件式を使ったswitch
	fmt.Println("\n=== 条件式switch ===")
	point := 85

	switch {
	case point >= 90:
		fmt.Println("A評価")
	case point >= 80:
		fmt.Println("B評価")
	case point >= 70:
		fmt.Println("C評価")
	default:
		fmt.Println("D評価")
	}

	// 初期化文付きswitch
	fmt.Println("\n=== 初期化文付きswitch ===")
	switch result := 10 % 3; result {
	case 0:
		fmt.Println("割り切れます")
	case 1:
		fmt.Println("余りは1です")
	case 2:
		fmt.Println("余りは2です")
	}

	// fallthroughの使用（次のcaseも実行）
	fmt.Println("\n=== fallthrough ===")
	num := 1
	switch num {
	case 1:
		fmt.Println("1です")
		fallthrough
	case 2:
		fmt.Println("2以下です")
	}

	// ========== ネストしたループ ==========
	fmt.Println("\n=== ネストしたループ ===")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// ラベル付きbreak（外側のループを抜ける）
	fmt.Println("\n=== ラベル付きbreak ===")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break OuterLoop // 外側のループを抜ける
			}
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println("ループ終了")
}

/*
【実行方法】
go run 05_control_flow.go

【重要な概念】
1. Goのif文は条件式に()が不要（むしろ付けない）
2. forはGoの唯一のループ構文（whileはない）
3. switchは自動的にbreakする（fallthroughで継続）
4. rangeは配列、スライス、マップ、文字列に使える

【他の言語との違い】
- while文がない → for文で代用
- do-while文がない
- 三項演算子がない → if-elseを使う
- switchは自動break（Cとは逆）

【次のステップ】
06_arrays_slices.go で配列とスライスを学びましょう
*/
