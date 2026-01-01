package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

/*
【学習ポイント】
1. エラーの基本
2. カスタムエラー
3. エラーのラップ
4. errors.Is と errors.As
5. panic と recover
*/

func main() {
	// ========== 基本的なエラーハンドリング ==========
	fmt.Println("=== 基本的なエラーハンドリング ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}
	fmt.Printf("結果: %d\n", result)

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== errors.New でエラーを作成 ==========
	fmt.Println("\n=== errors.New ===")

	err = validateAge(-5)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== fmt.Errorf でフォーマット付きエラー ==========
	fmt.Println("\n=== fmt.Errorf ===")

	err = processUser("", 25)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== カスタムエラー型 ==========
	fmt.Println("\n=== カスタムエラー型 ===")

	err = withdraw(1000, 1500)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)

		// 型アサーションでカスタムエラーを取得
		if insufficientErr, ok := err.(*InsufficientFundsError); ok {
			fmt.Printf("  残高: %d円\n", insufficientErr.Balance)
			fmt.Printf("  必要額: %d円\n", insufficientErr.Amount)
		}
	}

	// ========== エラーのラップ ==========
	fmt.Println("\n=== エラーのラップ ===")

	err = readConfig("config.json")
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== errors.Is ==========
	fmt.Println("\n=== errors.Is ===")

	err = openFile("nonexistent.txt")
	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("ファイルが見つかりません")
		} else {
			fmt.Printf("その他のエラー: %v\n", err)
		}
	}

	// ========== errors.As ==========
	fmt.Println("\n=== errors.As ===")

	err = processData("invalid")
	if err != nil {
		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("バリデーションエラー:\n")
			fmt.Printf("  フィールド: %s\n", validationErr.Field)
			fmt.Printf("  値: %v\n", validationErr.Value)
			fmt.Printf("  メッセージ: %s\n", validationErr.Message)
		}
	}

	// ========== 複数のエラーチェック ==========
	fmt.Println("\n=== 複数のエラーチェック ===")

	err = validateUser("", -5, "")
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}

	// ========== defer, panic, recover ==========
	fmt.Println("\n=== panic と recover ===")

	safeDivide(10, 0)
	fmt.Println("プログラムは続行します")

	// ========== ログ出力 ==========
	fmt.Println("\n=== ログ出力 ===")

	logExample()
}

// ========== 基本的なエラー ==========

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age < 0 {
		return errors.New("年齢は0以上である必要があります")
	}
	return nil
}

func processUser(name string, age int) error {
	if name == "" {
		return fmt.Errorf("名前が空です")
	}
	if age < 0 {
		return fmt.Errorf("無効な年齢: %d", age)
	}
	return nil
}

// ========== カスタムエラー型 ==========

type InsufficientFundsError struct {
	Balance int
	Amount  int
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("残高不足: 残高 %d円, 必要額 %d円", e.Balance, e.Amount)
}

func withdraw(balance, amount int) error {
	if amount > balance {
		return &InsufficientFundsError{
			Balance: balance,
			Amount:  amount,
		}
	}
	return nil
}

// ========== エラーのラップ ==========

func readFile(filename string) error {
	return errors.New("ファイルが見つかりません")
}

func readConfig(filename string) error {
	err := readFile(filename)
	if err != nil {
		// %w でエラーをラップ
		return fmt.Errorf("設定ファイルの読み込みに失敗: %w", err)
	}
	return nil
}

// ========== 定義済みエラー ==========

var (
	ErrFileNotFound     = errors.New("ファイルが見つかりません")
	ErrPermissionDenied = errors.New("権限がありません")
	ErrInvalidInput     = errors.New("無効な入力です")
)

func openFile(filename string) error {
	// 実際のファイル操作の代わりにエラーを返す
	return fmt.Errorf("ファイルを開けません: %w", ErrFileNotFound)
}

// ========== ValidationError ==========

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("バリデーションエラー [%s]: %s", e.Field, e.Message)
}

func processData(data string) error {
	if data == "invalid" {
		return &ValidationError{
			Field:   "data",
			Value:   data,
			Message: "データが無効です",
		}
	}
	return nil
}

// ========== 複数のエラーチェック ==========

func validateUser(name string, age int, email string) error {
	if name == "" {
		return fmt.Errorf("名前が空です")
	}
	if age < 0 {
		return fmt.Errorf("年齢が無効です: %d", age)
	}
	if email == "" {
		return fmt.Errorf("メールアドレスが空です")
	}
	return nil
}

// ========== panic と recover ==========

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("パニックを回復: %v\n", r)
		}
	}()

	if b == 0 {
		panic("0で割ろうとしました")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

// ========== ログ出力 ==========

func logExample() {
	// 標準ログ
	log.Println("これは情報ログです")

	// カスタムログ
	logger := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("カスタムログメッセージ")

	// エラーログ
	errLogger := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
	errLogger.Println("エラーが発生しました")
}

/*
【実行方法】
go run 05_errors.go

【重要な概念】
1. エラーは値として扱う
2. エラーは error インターフェースを実装
3. エラーチェックは明示的に行う
4. panic は回復不可能なエラーのみ

【エラーの作成方法】
1. errors.New("message")
2. fmt.Errorf("format", args...)
3. カスタムエラー型

【エラーハンドリングのパターン】
if err != nil {
    // エラー処理
    return err
}

【エラーのラップ】
fmt.Errorf("context: %w", err)
- %w でラップすると errors.Is/As が使える

【panic vs error】
panic:
- プログラムを停止させる
- 回復不可能なエラー
- ライブラリでは使わない

error:
- 通常のエラー処理
- 呼び出し側で処理可能
- 推奨される方法

【ベストプラクティス】
1. エラーは無視しない
2. エラーメッセージは小文字で始める
3. エラーは早期リターン
4. カスタムエラーで詳細情報を提供
5. panic は最小限に

【次のステップ】
03_concurrency/ で並行処理を学びましょう
*/
