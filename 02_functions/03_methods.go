package main

import (
	"fmt"
	"math"
)

/*
【学習ポイント】
1. メソッドの定義
2. 値レシーバ vs ポインタレシーバ
3. メソッドチェーン
4. 型にメソッドを追加
*/

func main() {
	// ========== 基本的なメソッド ==========
	fmt.Println("=== 基本的なメソッド ===")

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("長方形: %+v\n", rect)
	fmt.Printf("面積: %.2f\n", rect.Area())
	fmt.Printf("周囲: %.2f\n", rect.Perimeter())

	// ========== 値レシーバ vs ポインタレシーバ ==========
	fmt.Println("\n=== 値レシーバ vs ポインタレシーバ ===")

	counter := Counter{Count: 0}
	fmt.Printf("初期値: %d\n", counter.Count)

	// 値レシーバ（コピーを変更するので元の値は変わらない）
	counter.IncrementValue()
	fmt.Printf("値レシーバ後: %d (変わらない)\n", counter.Count)

	// ポインタレシーバ（元の値を変更）
	counter.IncrementPointer()
	fmt.Printf("ポインタレシーバ後: %d (変わる)\n", counter.Count)

	// ========== メソッドチェーン ==========
	fmt.Println("\n=== メソッドチェーン ===")

	builder := &StringBuilder{}
	result := builder.
		Append("Hello").
		Append(" ").
		Append("World").
		Append("!").
		String()

	fmt.Printf("結果: %s\n", result)

	// ========== 異なる型に同じメソッド名 ==========
	fmt.Println("\n=== 異なる型に同じメソッド名 ===")

	circle := Circle{Radius: 5}
	square := Square{Side: 4}

	fmt.Printf("円の面積: %.2f\n", circle.Area())
	fmt.Printf("正方形の面積: %.2f\n", square.Area())

	// ========== 基本型にメソッドを追加 ==========
	fmt.Println("\n=== 基本型にメソッドを追加 ===")

	var temp Temperature = 25.5
	fmt.Printf("%.1f°C = %.1f°F\n", temp, temp.ToFahrenheit())
	fmt.Printf("%.1f°C = %.1fK\n", temp, temp.ToKelvin())

	// ========== スライス型にメソッドを追加 ==========
	fmt.Println("\n=== スライス型にメソッドを追加 ===")

	numbers := IntSlice{1, 2, 3, 4, 5}
	fmt.Printf("数値: %v\n", numbers)
	fmt.Printf("合計: %d\n", numbers.Sum())
	fmt.Printf("平均: %.2f\n", numbers.Average())
	fmt.Printf("最大: %d\n", numbers.Max())

	// ========== 構造体のメソッド実用例 ==========
	fmt.Println("\n=== 実用例：銀行口座 ===")

	account := NewBankAccount("太郎", 1000)
	fmt.Printf("初期残高: %d円\n", account.Balance())

	account.Deposit(500)
	fmt.Printf("入金後: %d円\n", account.Balance())

	err := account.Withdraw(300)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("出金後: %d円\n", account.Balance())
	}

	err = account.Withdraw(2000)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== Stringerインターフェースの実装 ==========
	fmt.Println("\n=== Stringerインターフェース ===")

	person := Person{Name: "太郎", Age: 25}
	fmt.Println(person) // String()メソッドが自動的に呼ばれる

	// ========== ゲッター・セッター ==========
	fmt.Println("\n=== ゲッター・セッター ===")

	user := NewUser("taro", "taro@example.com")
	fmt.Printf("ユーザー: %s (%s)\n", user.Name(), user.Email())

	user.SetEmail("newemail@example.com")
	fmt.Printf("更新後: %s\n", user.Email())
}

// ========== 構造体の定義 ==========

type Rectangle struct {
	Width  float64
	Height float64
}

// 値レシーバのメソッド
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ========== 値レシーバ vs ポインタレシーバ ==========

type Counter struct {
	Count int
}

// 値レシーバ（コピーを変更）
func (c Counter) IncrementValue() {
	c.Count++ // コピーを変更するので元の値は変わらない
}

// ポインタレシーバ（元の値を変更）
func (c *Counter) IncrementPointer() {
	c.Count++ // 元の値を変更
}

// ========== メソッドチェーン ==========

type StringBuilder struct {
	data string
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
	sb.data += s
	return sb // 自分自身を返す
}

func (sb *StringBuilder) String() string {
	return sb.data
}

// ========== 異なる型に同じメソッド名 ==========

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// ========== 基本型にメソッドを追加 ==========

type Temperature float64

func (t Temperature) ToFahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperature) ToKelvin() float64 {
	return float64(t) + 273.15
}

// ========== スライス型にメソッドを追加 ==========

type IntSlice []int

func (is IntSlice) Sum() int {
	total := 0
	for _, v := range is {
		total += v
	}
	return total
}

func (is IntSlice) Average() float64 {
	if len(is) == 0 {
		return 0
	}
	return float64(is.Sum()) / float64(len(is))
}

func (is IntSlice) Max() int {
	if len(is) == 0 {
		return 0
	}
	max := is[0]
	for _, v := range is {
		if v > max {
			max = v
		}
	}
	return max
}

// ========== 実用例：銀行口座 ==========

type BankAccount struct {
	owner   string
	balance int
}

func NewBankAccount(owner string, initialBalance int) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

func (ba *BankAccount) Deposit(amount int) {
	ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount int) error {
	if amount > ba.balance {
		return fmt.Errorf("残高不足: 残高 %d円, 出金額 %d円", ba.balance, amount)
	}
	ba.balance -= amount
	return nil
}

func (ba *BankAccount) Balance() int {
	return ba.balance
}

// ========== Stringerインターフェース ==========

type Person struct {
	Name string
	Age  int
}

// fmt.Stringer インターフェースの実装
func (p Person) String() string {
	return fmt.Sprintf("%s (%d歳)", p.Name, p.Age)
}

// ========== ゲッター・セッター ==========

type User struct {
	name  string // 非公開フィールド（小文字）
	email string
}

func NewUser(name, email string) *User {
	return &User{
		name:  name,
		email: email,
	}
}

// ゲッター（Goでは Get プレフィックスを付けない）
func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

// セッター
func (u *User) SetEmail(email string) {
	u.email = email
}

/*
【実行方法】
go run 03_methods.go

【重要な概念】
1. メソッドは特定の型に紐付いた関数
2. レシーバは値型とポインタ型の2種類
3. 任意の型にメソッドを追加できる（基本型は除く）

【メソッドの定義】
func (レシーバ 型) メソッド名(引数) 戻り値 {
    // 処理
}

【値レシーバ vs ポインタレシーバ】
値レシーバ:
- コピーが渡される
- 元の値を変更しない
- 小さな構造体に適している

ポインタレシーバ:
- 参照が渡される
- 元の値を変更できる
- 大きな構造体に適している
- 変更が必要な場合に使う

【ベストプラクティス】
1. 通常はポインタレシーバを使う
2. 一つの型のメソッドは統一する（混在させない）
3. ゲッターには Get プレフィックスを付けない
4. メソッドチェーンは *型 を返す

【次のステップ】
04_interfaces.go でインターフェースを学びましょう
*/
