package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
【学習ポイント】
1. JSON のエンコード/デコード
2. 構造体タグ
3. カスタムマーシャリング
4. JSON と interface{}
*/

func main() {
	// ========== 構造体 → JSON（エンコード） ==========
	fmt.Println("=== 構造体 → JSON ===")

	person := Person{
		Name:  "太郎",
		Age:   25,
		Email: "taro@example.com",
	}

	// json.Marshal でエンコード
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON: %s\n", string(jsonData))

	// json.MarshalIndent で整形
	jsonIndent, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("整形JSON:\n%s\n", string(jsonIndent))

	// ========== JSON → 構造体（デコード） ==========
	fmt.Println("\n=== JSON → 構造体 ===")

	jsonStr := `{"name":"花子","age":23,"email":"hanako@example.com"}`

	var person2 Person
	err = json.Unmarshal([]byte(jsonStr), &person2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("デコード結果: %+v\n", person2)

	// ========== 構造体タグ ==========
	fmt.Println("\n=== 構造体タグ ===")

	user := User{
		ID:       1,
		Username: "taro",
		Email:    "taro@example.com",
		Password: "secret123",
		IsActive: true,
	}

	jsonUser, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("タグ付き構造体:\n%s\n", string(jsonUser))

	// ========== スライスと JSON ==========
	fmt.Println("\n=== スライスと JSON ===")

	people := []Person{
		{Name: "太郎", Age: 25, Email: "taro@example.com"},
		{Name: "花子", Age: 23, Email: "hanako@example.com"},
		{Name: "次郎", Age: 30, Email: "jiro@example.com"},
	}

	jsonPeople, _ := json.MarshalIndent(people, "", "  ")
	fmt.Printf("配列:\n%s\n", string(jsonPeople))

	// ========== マップと JSON ==========
	fmt.Println("\n=== マップと JSON ===")

	data := map[string]interface{}{
		"name":    "太郎",
		"age":     25,
		"active":  true,
		"scores":  []int{85, 90, 78},
		"address": map[string]string{"city": "東京", "country": "日本"},
	}

	jsonMap, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("マップ:\n%s\n", string(jsonMap))

	// ========== interface{} でデコード ==========
	fmt.Println("\n=== interface{} でデコード ===")

	jsonStr2 := `{"name":"太郎","age":25,"active":true}`

	var result map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr2), &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("結果: %+v\n", result)
	fmt.Printf("name: %v (型: %T)\n", result["name"], result["name"])
	fmt.Printf("age: %v (型: %T)\n", result["age"], result["age"])

	// 型アサーションで使用
	if name, ok := result["name"].(string); ok {
		fmt.Printf("名前: %s\n", name)
	}

	// ========== ネストした構造体 ==========
	fmt.Println("\n=== ネストした構造体 ===")

	company := Company{
		Name: "ABC株式会社",
		Address: Address{
			City:    "東京",
			Country: "日本",
		},
		Employees: []Person{
			{Name: "太郎", Age: 25, Email: "taro@example.com"},
			{Name: "花子", Age: 23, Email: "hanako@example.com"},
		},
	}

	jsonCompany, _ := json.MarshalIndent(company, "", "  ")
	fmt.Printf("会社情報:\n%s\n", string(jsonCompany))

	// ========== カスタムマーシャリング ==========
	fmt.Println("\n=== カスタムマーシャリング ===")

	product := Product{
		Name:  "ノートPC",
		Price: 150000,
	}

	jsonProduct, _ := json.MarshalIndent(product, "", "  ")
	fmt.Printf("商品:\n%s\n", string(jsonProduct))

	// ========== エラーハンドリング ==========
	fmt.Println("\n=== エラーハンドリング ===")

	invalidJSON := `{"name": "太郎", "age": "無効な年齢"}`
	var person3 Person
	err = json.Unmarshal([]byte(invalidJSON), &person3)
	if err != nil {
		fmt.Printf("デコードエラー: %v\n", err)
	}

	// ========== 実用例：API レスポンス ==========
	fmt.Println("\n=== 実用例：API レスポンス ===")

	apiResp := APIResponse{
		Success: true,
		Message: "データ取得成功",
		Data: map[string]interface{}{
			"user_id": 123,
			"name":    "太郎",
		},
	}

	jsonResp, _ := json.MarshalIndent(apiResp, "", "  ")
	fmt.Printf("APIレスポンス:\n%s\n", string(jsonResp))
}

// ========== 構造体の定義 ==========

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`                   // JSON に含めない
	IsActive bool   `json:"is_active,omitempty"` // 空の場合は省略
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Company struct {
	Name      string   `json:"name"`
	Address   Address  `json:"address"`
	Employees []Person `json:"employees"`
}

// カスタムマーシャリング
type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	type Alias Product
	return json.Marshal(&struct {
		*Alias
		PriceFormatted string `json:"price_formatted"`
	}{
		Alias:          (*Alias)(&p),
		PriceFormatted: fmt.Sprintf("¥%d", p.Price),
	})
}

type APIResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

/*
【実行方法】
go run 01_json.go

【重要な概念】
1. json.Marshal: 構造体 → JSON
2. json.Unmarshal: JSON → 構造体
3. 構造体タグでフィールド名を制御
4. interface{} で柔軟なデータ処理

【構造体タグ】
`json:"フィールド名,オプション"`

オプション:
- omitempty: 空の場合は省略
- -: JSON に含めない

【よく使う関数】
json.Marshal(v)              // エンコード
json.MarshalIndent(v, "", "  ") // 整形エンコード
json.Unmarshal(data, &v)     // デコード

【注意点】
1. 公開フィールド（大文字始まり）のみエンコード
2. Unmarshal は &（ポインタ）を渡す
3. 数値は float64 としてデコードされる

【ベストプラクティス】
1. 構造体タグを使う
2. omitempty で不要なフィールドを省略
3. Password等は "-" で除外
4. エラーハンドリングを忘れない

【次のステップ】
02_http_client.go で HTTP クライアントを学びましょう
*/
