package main

import "fmt"

func main(){
	fmt.Println("hello")
	var x, y int = 3, 10
	fmt.Printf("%d\n", x + y)
	fmt.Printf("%d\n", x - y)
	fmt.Printf("%d\n", x / y)
	fmt.Printf("%d\n", x * y)

	var count int = 0
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)

	// 比較、論理、否定はjsと同じ

	// ビット演算 できるとしても何に利用できるのか理解していない。
  var p, q int = 12, 10
	fmt.Printf("and %d & %d = %d \n", p, q, p&q)
	fmt.Printf(" or %d | %d = %d \n", p, q, p|q)
	fmt.Printf("xor %d ^ %d = %d \n", p, q, p^q)
	fmt.Printf("and not %d &^ %d = %d \n", p, q, p&^q)
	fmt.Printf("ssl %d << 1 = %d \n", p, p << 1)
	fmt.Printf("srl %d >> 1 = %d \n", p, p >> 1)

}