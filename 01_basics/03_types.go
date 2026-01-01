package main

import "fmt"

/*
【学習ポイント】
1. Goの基本型
2. 数値型の種類とサイズ
3. 文字列とルーン
4. 型のサイズと範囲
*/

func main() {
	// ========== 整数型 ==========
	var i8 int8 = 127                   // -128 ~ 127
	var i16 int16 = 32767               // -32768 ~ 32767
	var i32 int32 = 2147483647          // -2^31 ~ 2^31-1
	var i64 int64 = 9223372036854775807 // -2^63 ~ 2^63-1

	fmt.Println("=== 符号付き整数型 ===")
	fmt.Printf("int8:  %d\n", i8)
	fmt.Printf("int16: %d\n", i16)
	fmt.Printf("int32: %d\n", i32)
	fmt.Printf("int64: %d\n", i64)

	// 符号なし整数型
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	fmt.Println("\n=== 符号なし整数型 ===")
	fmt.Printf("uint8:  %d\n", ui8)
	fmt.Printf("uint16: %d\n", ui16)
	fmt.Printf("uint32: %d\n", ui32)
	fmt.Printf("uint64: %d\n", ui64)

	// プラットフォーム依存の整数型
	var i int = 42    // 32bit or 64bit（環境依存）
	var ui uint = 100 // 32bit or 64bit（環境依存）
	fmt.Println("\n=== プラットフォーム依存型 ===")
	fmt.Printf("int:  %d\n", i)
	fmt.Printf("uint: %d\n", ui)

	// ========== 浮動小数点型 ==========
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	fmt.Println("\n=== 浮動小数点型 ===")
	fmt.Printf("float32: %.2f\n", f32)
	fmt.Printf("float64: %.15f\n", f64)

	// ========== 複素数型 ==========
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	fmt.Println("\n=== 複素数型 ===")
	fmt.Printf("complex64:  %v\n", c64)
	fmt.Printf("complex128: %v\n", c128)

	// ========== 真偽値型 ==========
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println("\n=== 真偽値型 ===")
	fmt.Printf("isTrue:  %t\n", isTrue)
	fmt.Printf("isFalse: %t\n", isFalse)

	// ========== 文字列型 ==========
	var str string = "こんにちは、Go!"
	var multiLine string = `これは
複数行の
文字列です`

	fmt.Println("\n=== 文字列型 ===")
	fmt.Printf("文字列: %s\n", str)
	fmt.Printf("文字列長: %d バイト\n", len(str))
	fmt.Println(multiLine)

	// ========== ルーン型（文字） ==========
	var r rune = 'あ' // rune は int32 のエイリアス
	var b byte = 'A' // byte は uint8 のエイリアス

	fmt.Println("\n=== ルーン・バイト型 ===")
	fmt.Printf("rune: %c (Unicode: %U)\n", r, r)
	fmt.Printf("byte: %c (ASCII: %d)\n", b, b)

	// ========== 型のエイリアス ==========
	// byte = uint8
	// rune = int32

	// ========== 型の確認 ==========
	fmt.Println("\n=== 型の確認 ===")
	fmt.Printf("i の型: %T\n", i)
	fmt.Printf("f64 の型: %T\n", f64)
	fmt.Printf("str の型: %T\n", str)
	fmt.Printf("isTrue の型: %T\n", isTrue)

	// ========== 型変換の例 ==========
	var a int = 10
	var b float64 = 3.14
	// var c = a + b // エラー！異なる型同士は演算できない
	var c = float64(a) + b // 明示的な型変換が必要

	fmt.Println("\n=== 型変換 ===")
	fmt.Printf("int(%d) + float64(%.2f) = %.2f\n", a, b, c)
}

/*
【実行方法】
go run 03_types.go

【重要な概念】
1. Goは静的型付け言語 - 型は厳密に管理される
2. 異なる型同士の演算には明示的な型変換が必要
3. int, uint のサイズは環境依存（通常は64bit）
4. 文字列はUTF-8エンコード
5. rune は1文字のUnicode、byte は1バイトを表す

【よく使う型】
- int: 整数（通常はこれを使う）
- float64: 浮動小数点数
- string: 文字列
- bool: 真偽値

【次のステップ】
04_operators.go で演算子を学びましょう
*/
