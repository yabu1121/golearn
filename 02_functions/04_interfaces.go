package main

import (
	"fmt"
	"math"
)

/*
【学習ポイント】
1. インターフェースの定義と実装
2. 暗黙的な実装
3. 空インターフェース
4. 型アサーション
5. 型スイッチ
6. 標準インターフェース
*/

func main() {
	// ========== 基本的なインターフェース ==========
	fmt.Println("=== 基本的なインターフェース ===")

	var s Shape

	// Circleを代入
	s = Circle{Radius: 5}
	printShapeInfo(s)

	// Rectangleを代入
	s = Rectangle{Width: 10, Height: 5}
	printShapeInfo(s)

	// ========== インターフェースのスライス ==========
	fmt.Println("\n=== インターフェースのスライス ===")

	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 5, Height: 8},
	}

	totalArea := 0.0
	for i, shape := range shapes {
		area := shape.Area()
		fmt.Printf("%d: 面積 = %.2f\n", i+1, area)
		totalArea += area
	}
	fmt.Printf("合計面積: %.2f\n", totalArea)

	// ========== 複数のインターフェース ==========
	fmt.Println("\n=== 複数のインターフェース ===")

	dog := Dog{Name: "ポチ"}
	cat := Cat{Name: "タマ"}

	makeSound(dog)
	makeSound(cat)

	// ========== インターフェースの埋め込み ==========
	fmt.Println("\n=== インターフェースの埋め込み ===")

	file := File{Name: "document.txt", Size: 1024}
	printReadWriterInfo(file)

	// ========== 空インターフェース ==========
	fmt.Println("\n=== 空インターフェース ===")

	printAnything(42)
	printAnything("Hello")
	printAnything(3.14)
	printAnything([]int{1, 2, 3})
	printAnything(Circle{Radius: 5})

	// ========== 型アサーション ==========
	fmt.Println("\n=== 型アサーション ===")

	var i interface{} = "Hello"

	// 型アサーション（成功）
	s1, ok := i.(string)
	if ok {
		fmt.Printf("文字列: %s\n", s1)
	}

	// 型アサーション（失敗）
	n, ok := i.(int)
	if !ok {
		fmt.Println("int型ではありません")
	} else {
		fmt.Printf("整数: %d\n", n)
	}

	// panicになる例（ok を使わない場合）
	// n := i.(int) // panic!

	// ========== 型スイッチ ==========
	fmt.Println("\n=== 型スイッチ ===")

	checkType(42)
	checkType("Hello")
	checkType(3.14)
	checkType(true)
	checkType(Circle{Radius: 5})

	// ========== 標準インターフェース：Stringer ==========
	fmt.Println("\n=== Stringer インターフェース ===")

	person := Person{Name: "太郎", Age: 25}
	fmt.Println(person) // String()が自動的に呼ばれる

	// ========== 標準インターフェース：Error ==========
	fmt.Println("\n=== Error インターフェース ===")

	err := divide(10, 0)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	result, err := divide(10, 2)
	if err == nil {
		fmt.Printf("結果: %d\n", result)
	}

	// ========== インターフェースの実用例 ==========
	fmt.Println("\n=== 実用例：データベース ===")

	var db Database

	// MySQLを使用
	db = &MySQL{Host: "localhost"}
	db.Connect()
	db.Query("SELECT * FROM users")
	db.Close()

	fmt.Println()

	// PostgreSQLを使用
	db = &PostgreSQL{Host: "localhost"}
	db.Connect()
	db.Query("SELECT * FROM users")
	db.Close()

	// ========== インターフェースの値の比較 ==========
	fmt.Println("\n=== インターフェースの比較 ===")

	var s2, s3 Shape
	s2 = Circle{Radius: 5}
	s3 = Circle{Radius: 5}

	fmt.Printf("s2 == s3: %t\n", s2 == s3)

	// nilとの比較
	var s4 Shape
	fmt.Printf("s4 == nil: %t\n", s4 == nil)
}

// ========== インターフェースの定義 ==========

// Shapeインターフェース
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle（Shapeを実装）
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle（Shapeを実装）
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle（Shapeを実装）
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func (t Triangle) Perimeter() float64 {
	// 簡略化のため固定値
	return t.Base * 3
}

// インターフェースを受け取る関数
func printShapeInfo(s Shape) {
	fmt.Printf("面積: %.2f, 周囲: %.2f\n", s.Area(), s.Perimeter())
}

// ========== 複数のインターフェース ==========

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return "ワンワン"
}

type Cat struct {
	Name string
}

func (c Cat) Sound() string {
	return "ニャー"
}

func makeSound(a Animal) {
	fmt.Printf("%s\n", a.Sound())
}

// ========== インターフェースの埋め込み ==========

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReadWriterはReaderとWriterを埋め込む
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Name string
	Size int
}

func (f File) Read() string {
	return fmt.Sprintf("Reading from %s", f.Name)
}

func (f File) Write(data string) {
	fmt.Printf("Writing to %s: %s\n", f.Name, data)
}

func printReadWriterInfo(rw ReadWriter) {
	fmt.Println(rw.Read())
	rw.Write("Hello")
}

// ========== 空インターフェース ==========

func printAnything(v interface{}) {
	fmt.Printf("値: %v, 型: %T\n", v, v)
}

// ========== 型スイッチ ==========

func checkType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("文字列: %s\n", v)
	case float64:
		fmt.Printf("浮動小数点: %.2f\n", v)
	case bool:
		fmt.Printf("真偽値: %t\n", v)
	case Circle:
		fmt.Printf("円: 半径 %.2f\n", v.Radius)
	default:
		fmt.Printf("不明な型: %T\n", v)
	}
}

// ========== Stringerインターフェース ==========

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d歳)", p.Name, p.Age)
}

// ========== Errorインターフェース ==========

type DivisionError struct {
	Dividend int
	Divisor  int
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("0で割ることはできません: %d / %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

// ========== 実用例：データベース ==========

type Database interface {
	Connect() error
	Query(sql string) ([]string, error)
	Close() error
}

type MySQL struct {
	Host string
}

func (m *MySQL) Connect() error {
	fmt.Printf("MySQLに接続: %s\n", m.Host)
	return nil
}

func (m *MySQL) Query(sql string) ([]string, error) {
	fmt.Printf("MySQLクエリ: %s\n", sql)
	return []string{"result1", "result2"}, nil
}

func (m *MySQL) Close() error {
	fmt.Println("MySQL接続を閉じる")
	return nil
}

type PostgreSQL struct {
	Host string
}

func (p *PostgreSQL) Connect() error {
	fmt.Printf("PostgreSQLに接続: %s\n", p.Host)
	return nil
}

func (p *PostgreSQL) Query(sql string) ([]string, error) {
	fmt.Printf("PostgreSQLクエリ: %s\n", sql)
	return []string{"result1", "result2"}, nil
}

func (p *PostgreSQL) Close() error {
	fmt.Println("PostgreSQL接続を閉じる")
	return nil
}

/*
【実行方法】
go run 04_interfaces.go

【重要な概念】
1. インターフェースはメソッドの集合
2. 実装は暗黙的（implements キーワード不要）
3. 空インターフェース interface{} は任意の型を受け入れる
4. 型アサーションで具体的な型に変換

【インターフェースの定義】
type インターフェース名 interface {
    メソッド名(引数) 戻り値
    ...
}

【暗黙的な実装】
- 明示的な宣言不要
- メソッドを実装すれば自動的にインターフェースを満たす

【標準インターフェース】
- fmt.Stringer: String() string
- error: Error() string
- io.Reader: Read(p []byte) (n int, err error)
- io.Writer: Write(p []byte) (n int, err error)

【ベストプラクティス】
1. インターフェースは小さく保つ
2. 使う側で定義する（提供側ではない）
3. 空インターフェースは最小限に
4. 型アサーションは ok を使ってチェック

【次のステップ】
05_errors.go でエラーハンドリングを学びましょう
*/
