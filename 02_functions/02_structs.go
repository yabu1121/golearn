package main

import "fmt"

/*
【学習ポイント】
1. 構造体の定義と初期化
2. フィールドへのアクセス
3. 構造体のポインタ
4. 埋め込み（継承のような機能）
5. タグ
*/

func main() {
	// ========== 構造体の定義と初期化 ==========
	fmt.Println("=== 構造体の基本 ===")

	// 方法1: フィールド名を指定
	person1 := Person{
		Name: "太郎",
		Age:  25,
		City: "東京",
	}
	fmt.Printf("person1: %+v\n", person1)

	// 方法2: 順序で指定（非推奨）
	person2 := Person{"花子", 23, "大阪"}
	fmt.Printf("person2: %+v\n", person2)

	// 方法3: 一部のフィールドのみ初期化
	person3 := Person{Name: "次郎"}
	fmt.Printf("person3: %+v\n", person3) // 未指定はゼロ値

	// 方法4: var で宣言（ゼロ値で初期化）
	var person4 Person
	fmt.Printf("person4: %+v\n", person4)

	// ========== フィールドへのアクセス ==========
	fmt.Println("\n=== フィールドへのアクセス ===")

	person1.Age = 26
	fmt.Printf("更新後: %+v\n", person1)
	fmt.Printf("名前: %s, 年齢: %d\n", person1.Name, person1.Age)

	// ========== 構造体のポインタ ==========
	fmt.Println("\n=== 構造体のポインタ ===")

	// &演算子でポインタを取得
	p := &Person{Name: "四郎", Age: 30}
	fmt.Printf("ポインタ: %p\n", p)
	fmt.Printf("値: %+v\n", *p)

	// ポインタ経由でのアクセス（自動的に間接参照される）
	p.Age = 31 // (*p).Age と同じ
	fmt.Printf("更新後: %+v\n", *p)

	// new関数で作成
	p2 := new(Person)
	p2.Name = "五郎"
	p2.Age = 28
	fmt.Printf("new: %+v\n", *p2)

	// ========== 構造体のコピー ==========
	fmt.Println("\n=== 構造体のコピー ===")

	original := Person{Name: "太郎", Age: 25}
	copy := original // 値のコピー
	copy.Age = 30

	fmt.Printf("original: %+v\n", original) // 変わらない
	fmt.Printf("copy: %+v\n", copy)

	// ========== 匿名構造体 ==========
	fmt.Println("\n=== 匿名構造体 ===")

	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("config: %+v\n", config)

	// ========== ネストした構造体 ==========
	fmt.Println("\n=== ネストした構造体 ===")

	employee := Employee{
		Person: Person{
			Name: "山田太郎",
			Age:  30,
			City: "東京",
		},
		Company:  "ABC株式会社",
		Position: "エンジニア",
	}
	fmt.Printf("employee: %+v\n", employee)
	fmt.Printf("名前: %s, 会社: %s\n", employee.Person.Name, employee.Company)

	// ========== 埋め込み（Embedding） ==========
	fmt.Println("\n=== 埋め込み ===")

	manager := Manager{
		Person: Person{
			Name: "佐藤花子",
			Age:  35,
			City: "大阪",
		},
		Department: "開発部",
		TeamSize:   10,
	}

	// 埋め込まれた構造体のフィールドに直接アクセス可能
	fmt.Printf("名前: %s\n", manager.Name) // manager.Person.Name と同じ
	fmt.Printf("部署: %s\n", manager.Department)

	// ========== 構造体のスライス ==========
	fmt.Println("\n=== 構造体のスライス ===")

	people := []Person{
		{Name: "太郎", Age: 25, City: "東京"},
		{Name: "花子", Age: 23, City: "大阪"},
		{Name: "次郎", Age: 30, City: "名古屋"},
	}

	for i, person := range people {
		fmt.Printf("%d: %s (%d歳) - %s\n", i+1, person.Name, person.Age, person.City)
	}

	// ========== 構造体のマップ ==========
	fmt.Println("\n=== 構造体のマップ ===")

	users := map[string]Person{
		"user1": {Name: "太郎", Age: 25},
		"user2": {Name: "花子", Age: 23},
	}

	for id, user := range users {
		fmt.Printf("%s: %s (%d歳)\n", id, user.Name, user.Age)
	}

	// ========== 構造体の比較 ==========
	fmt.Println("\n=== 構造体の比較 ===")

	p1 := Person{Name: "太郎", Age: 25, City: "東京"}
	p2 := Person{Name: "太郎", Age: 25, City: "東京"}
	p3 := Person{Name: "花子", Age: 23, City: "大阪"}

	fmt.Printf("p1 == p2: %t\n", p1 == p2) // true
	fmt.Printf("p1 == p3: %t\n", p1 == p3) // false

	// ========== タグ（メタデータ） ==========
	fmt.Println("\n=== タグ ===")

	user := User{
		ID:       1,
		Username: "taro",
		Email:    "taro@example.com",
		Password: "secret",
	}

	// タグはリフレクションで取得可能（JSONエンコード等で使用）
	fmt.Printf("user: %+v\n", user)

	// ========== 構造体を返す関数 ==========
	fmt.Println("\n=== 構造体を返す関数 ===")

	newPerson := NewPerson("太郎", 25)
	fmt.Printf("newPerson: %+v\n", newPerson)

	// ポインタを返すコンストラクタ
	newPerson2 := NewPersonPointer("花子", 23)
	fmt.Printf("newPerson2: %+v\n", *newPerson2)
}

// ========== 構造体の定義 ==========

type Person struct {
	Name string
	Age  int
	City string
}

type Employee struct {
	Person   Person // ネスト
	Company  string
	Position string
}

type Manager struct {
	Person     // 埋め込み（フィールド名を省略）
	Department string
	TeamSize   int
}

// タグ付き構造体
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // JSONに含めない
}

// ========== コンストラクタ関数 ==========

// 値を返すコンストラクタ
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// ポインタを返すコンストラクタ（推奨）
func NewPersonPointer(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

/*
【実行方法】
go run 02_structs.go

【重要な概念】
1. 構造体は値型（コピーされる）
2. 埋め込みで継承のような機能を実現
3. タグはメタデータとして使用（JSON, DB等）
4. 構造体は比較可能（全フィールドが比較可能な場合）

【構造体の定義】
type 構造体名 struct {
    フィールド名 型
    ...
}

【初期化のベストプラクティス】
1. フィールド名を明示する
2. コンストラクタ関数を使う
3. 大きな構造体はポインタで返す

【埋め込み vs ネスト】
埋め込み:
- フィールド名を省略
- 埋め込まれた型のフィールドに直接アクセス可能

ネスト:
- フィールド名を明示
- フィールド名経由でアクセス

【次のステップ】
03_methods.go でメソッドを学びましょう
*/
