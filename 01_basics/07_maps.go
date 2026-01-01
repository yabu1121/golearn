package main

import "fmt"

/*
【学習ポイント】
1. マップの宣言と初期化
2. マップの操作（追加、取得、削除）
3. キーの存在確認
4. マップのループ
*/

func main() {
	// ========== マップの宣言と初期化 ==========
	fmt.Println("=== マップの宣言 ===")

	// 宣言のみ（nilマップ）
	var m1 map[string]int
	fmt.Printf("m1: %v (nil: %t)\n", m1, m1 == nil)

	// make関数で初期化
	m2 := make(map[string]int)
	fmt.Printf("m2: %v (nil: %t)\n", m2, m2 == nil)

	// リテラルで初期化
	m3 := map[string]int{
		"太郎": 25,
		"花子": 23,
		"次郎": 30,
	}
	fmt.Printf("m3: %v\n", m3)

	// ========== マップへの追加・更新 ==========
	fmt.Println("\n=== 追加・更新 ===")

	ages := make(map[string]int)

	// 要素の追加
	ages["太郎"] = 25
	ages["花子"] = 23
	ages["次郎"] = 30
	fmt.Printf("ages: %v\n", ages)

	// 要素の更新
	ages["太郎"] = 26
	fmt.Printf("更新後: %v\n", ages)

	// ========== マップからの取得 ==========
	fmt.Println("\n=== 取得 ===")

	// 値の取得
	age := ages["太郎"]
	fmt.Printf("太郎の年齢: %d\n", age)

	// 存在しないキー（ゼロ値が返る）
	unknownAge := ages["四郎"]
	fmt.Printf("四郎の年齢: %d\n", unknownAge)

	// キーの存在確認（2つ目の戻り値で判定）
	age, exists := ages["太郎"]
	if exists {
		fmt.Printf("太郎は存在します。年齢: %d\n", age)
	}

	age, exists = ages["四郎"]
	if !exists {
		fmt.Println("四郎は存在しません")
	}

	// ========== マップからの削除 ==========
	fmt.Println("\n=== 削除 ===")

	fmt.Printf("削除前: %v\n", ages)
	delete(ages, "次郎")
	fmt.Printf("削除後: %v\n", ages)

	// 存在しないキーを削除してもエラーにならない
	delete(ages, "存在しないキー")

	// ========== マップのループ ==========
	fmt.Println("\n=== ループ ===")

	scores := map[string]int{
		"数学": 85,
		"英語": 92,
		"国語": 78,
		"理科": 88,
	}

	// キーと値の両方
	for subject, score := range scores {
		fmt.Printf("%s: %d点\n", subject, score)
	}

	// キーのみ
	fmt.Println("\nキーのみ:")
	for subject := range scores {
		fmt.Printf("- %s\n", subject)
	}

	// 値のみ（キーを_で無視）
	fmt.Println("\n値のみ:")
	total := 0
	for _, score := range scores {
		total += score
	}
	fmt.Printf("合計点: %d\n", total)

	// ========== マップのサイズ ==========
	fmt.Println("\n=== サイズ ===")
	fmt.Printf("要素数: %d\n", len(scores))

	// ========== ネストしたマップ ==========
	fmt.Println("\n=== ネストしたマップ ===")

	students := map[string]map[string]int{
		"太郎": {
			"数学": 85,
			"英語": 90,
		},
		"花子": {
			"数学": 95,
			"英語": 88,
		},
	}

	for name, subjects := range students {
		fmt.Printf("%sの成績:\n", name)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d点\n", subject, score)
		}
	}

	// ========== マップと構造体の組み合わせ ==========
	fmt.Println("\n=== マップと構造体 ===")

	type Person struct {
		Name string
		Age  int
	}

	people := map[string]Person{
		"user1": {Name: "太郎", Age: 25},
		"user2": {Name: "花子", Age: 23},
	}

	for id, person := range people {
		fmt.Printf("%s: %s (%d歳)\n", id, person.Name, person.Age)
	}

	// ========== 実用例：カウンター ==========
	fmt.Println("\n=== 実用例：文字カウンター ===")

	text := "hello world"
	charCount := make(map[rune]int)

	for _, char := range text {
		charCount[char]++
	}

	for char, count := range charCount {
		fmt.Printf("'%c': %d回\n", char, count)
	}

	// ========== 実用例：グループ化 ==========
	fmt.Println("\n=== 実用例：年齢別グループ化 ===")

	users := []struct {
		Name string
		Age  int
	}{
		{"太郎", 25},
		{"花子", 25},
		{"次郎", 30},
		{"四郎", 25},
	}

	ageGroups := make(map[int][]string)
	for _, user := range users {
		ageGroups[user.Age] = append(ageGroups[user.Age], user.Name)
	}

	for age, names := range ageGroups {
		fmt.Printf("%d歳: %v\n", age, names)
	}

	// ========== マップの注意点 ==========
	fmt.Println("\n=== マップの注意点 ===")

	// nilマップには追加できない（panicになる）
	var nilMap map[string]int
	// nilMap["key"] = 1 // panic!

	// make()で初期化すれば追加できる
	nilMap = make(map[string]int)
	nilMap["key"] = 1
	fmt.Printf("初期化後: %v\n", nilMap)

	// マップは参照型
	original := map[string]int{"a": 1}
	reference := original
	reference["a"] = 2
	fmt.Printf("original: %v\n", original)   // {a: 2}
	fmt.Printf("reference: %v\n", reference) // {a: 2}
}

/*
【実行方法】
go run 07_maps.go

【重要な概念】
1. マップはキーと値のペアを格納する
2. キーは比較可能な型（==が使える型）
3. マップは参照型
4. マップの順序は保証されない

【マップの操作】
- 追加/更新: map[key] = value
- 取得: value := map[key]
- 削除: delete(map, key)
- 存在確認: value, exists := map[key]
- サイズ: len(map)

【注意点】
1. nilマップには要素を追加できない（make()で初期化が必要）
2. マップは並行アクセスに安全ではない（sync.Mapを使う）
3. マップの反復順序は保証されない

【ベストプラクティス】
- 容量が分かっている場合: make(map[K]V, capacity)
- 存在確認は2値受け取りで: value, ok := map[key]
- nilチェック: if map == nil

【次のステップ】
08_pointers.go でポインタを学びましょう
*/
