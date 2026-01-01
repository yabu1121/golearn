package main

import "fmt"

/*
【学習ポイント】
1. 算術演算子
2. 比較演算子
3. 論理演算子
4. ビット演算子
5. 代入演算子
*/

func main() {
	// ========== 算術演算子 ==========
	fmt.Println("=== 算術演算子 ===")
	a, b := 10, 3

	fmt.Printf("%d + %d = %d (加算)\n", a, b, a+b)
	fmt.Printf("%d - %d = %d (減算)\n", a, b, a-b)
	fmt.Printf("%d * %d = %d (乗算)\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (除算)\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d (剰余)\n", a, b, a%b)

	// インクリメント・デクリメント
	count := 0
	count++ // count = count + 1
	fmt.Printf("インクリメント: %d\n", count)
	count-- // count = count - 1
	fmt.Printf("デクリメント: %d\n", count)

	// 注意: Goでは ++count や --count は使えない
	// また、count++ は文であり、式ではない

	// ========== 比較演算子 ==========
	fmt.Println("\n=== 比較演算子 ===")
	x, y := 5, 10

	fmt.Printf("%d == %d: %t (等しい)\n", x, y, x == y)
	fmt.Printf("%d != %d: %t (等しくない)\n", x, y, x != y)
	fmt.Printf("%d < %d:  %t (より小さい)\n", x, y, x < y)
	fmt.Printf("%d <= %d: %t (以下)\n", x, y, x <= y)
	fmt.Printf("%d > %d:  %t (より大きい)\n", x, y, x > y)
	fmt.Printf("%d >= %d: %t (以上)\n", x, y, x >= y)

	// ========== 論理演算子 ==========
	fmt.Println("\n=== 論理演算子 ===")
	t, f := true, false

	fmt.Printf("true && false: %t (AND - 両方true)\n", t && f)
	fmt.Printf("true || false: %t (OR - どちらかtrue)\n", t || f)
	fmt.Printf("!true:         %t (NOT - 否定)\n", !t)

	// 実用例
	age := 25
	hasLicense := true
	canDrive := age >= 18 && hasLicense
	fmt.Printf("\n年齢%d歳、免許%t → 運転可能: %t\n", age, hasLicense, canDrive)

	// ========== ビット演算子 ==========
	fmt.Println("\n=== ビット演算子 ===")
	p, q := 12, 10 // 1100, 1010 (2進数)

	fmt.Printf("%d & %d  = %d (AND)\n", p, q, p&q)      // 1000 = 8
	fmt.Printf("%d | %d  = %d (OR)\n", p, q, p|q)       // 1110 = 14
	fmt.Printf("%d ^ %d  = %d (XOR)\n", p, q, p^q)      // 0110 = 6
	fmt.Printf("%d &^ %d = %d (AND NOT)\n", p, q, p&^q) // 0100 = 4
	fmt.Printf("%d << 1 = %d (左シフト)\n", p, p<<1)        // 11000 = 24
	fmt.Printf("%d >> 1 = %d (右シフト)\n", p, p>>1)        // 0110 = 6

	// ========== 代入演算子 ==========
	fmt.Println("\n=== 代入演算子 ===")
	num := 10
	fmt.Printf("初期値: %d\n", num)

	num += 5 // num = num + 5
	fmt.Printf("num += 5:  %d\n", num)

	num -= 3 // num = num - 3
	fmt.Printf("num -= 3:  %d\n", num)

	num *= 2 // num = num * 2
	fmt.Printf("num *= 2:  %d\n", num)

	num /= 4 // num = num / 4
	fmt.Printf("num /= 4:  %d\n", num)

	num %= 3 // num = num % 3
	fmt.Printf("num %%= 3:  %d\n", num)

	// ビット演算の代入
	bits := 8
	bits <<= 2 // bits = bits << 2
	fmt.Printf("bits <<= 2: %d\n", bits)

	// ========== 演算子の優先順位 ==========
	fmt.Println("\n=== 演算子の優先順位 ===")
	result1 := 2 + 3*4     // 乗算が先
	result2 := (2 + 3) * 4 // カッコが最優先
	fmt.Printf("2 + 3 * 4 = %d\n", result1)
	fmt.Printf("(2 + 3) * 4 = %d\n", result2)

	// 複雑な式
	result3 := 10 > 5 && 20 < 30 || false
	fmt.Printf("10 > 5 && 20 < 30 || false = %t\n", result3)
}

/*
【実行方法】
go run 04_operators.go

【演算子の優先順位（高→低）】
1. *, /, %, <<, >>, &, &^
2. +, -, |, ^
3. ==, !=, <, <=, >, >=
4. &&
5. ||

【重要な注意点】
1. Goには三項演算子がない（? : は使えない）
2. ++, -- は文であり、式ではない（x = y++ はエラー）
3. 整数の除算は整数を返す（10 / 3 = 3）
4. 異なる型同士の演算はエラー

【次のステップ】
05_control_flow.go で制御構文を学びましょう
*/
