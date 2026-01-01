package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
【学習ポイント】
1. データベース接続
2. CRUD操作
3. トランザクション
4. プリペアドステートメント
5. エラーハンドリング

【事前準備】
go get github.com/mattn/go-sqlite3
*/

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func main() {
	// ========== データベース接続 ==========
	fmt.Println("=== データベース接続 ===")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("データベース接続成功")

	// ========== テーブル作成 ==========
	fmt.Println("\n=== テーブル作成 ===")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("usersテーブル作成完了")

	// ========== INSERT（作成） ==========
	fmt.Println("\n=== INSERT ===")

	insertSQL := `INSERT INTO users (name, email) VALUES (?, ?)`

	result, err := db.Exec(insertSQL, "太郎", "taro@example.com")
	if err != nil {
		log.Printf("挿入エラー: %v\n", err)
	} else {
		id, _ := result.LastInsertId()
		fmt.Printf("ユーザー作成: ID=%d\n", id)
	}

	// 複数挿入
	users := []struct {
		name  string
		email string
	}{
		{"花子", "hanako@example.com"},
		{"次郎", "jiro@example.com"},
		{"四郎", "shiro@example.com"},
	}

	for _, u := range users {
		result, err := db.Exec(insertSQL, u.name, u.email)
		if err != nil {
			log.Printf("挿入エラー [%s]: %v\n", u.name, err)
		} else {
			id, _ := result.LastInsertId()
			fmt.Printf("ユーザー作成: %s (ID=%d)\n", u.name, id)
		}
	}

	// ========== SELECT（読み取り） ==========
	fmt.Println("\n=== SELECT ===")

	// 全件取得
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("全ユーザー:")
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  ID=%d, Name=%s, Email=%s, Created=%s\n",
			user.ID, user.Name, user.Email, user.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// エラーチェック
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// ========== SELECT（単一行） ==========
	fmt.Println("\n=== SELECT（単一行） ===")

	var user User
	err = db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", 1).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err == sql.ErrNoRows {
		fmt.Println("ユーザーが見つかりません")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ユーザー: %+v\n", user)
	}

	// ========== UPDATE（更新） ==========
	fmt.Println("\n=== UPDATE ===")

	updateSQL := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	result, err = db.Exec(updateSQL, "太郎2", "taro2@example.com", 1)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("更新された行数: %d\n", rowsAffected)

	// ========== DELETE（削除） ==========
	fmt.Println("\n=== DELETE ===")

	deleteSQL := `DELETE FROM users WHERE id = ?`
	result, err = db.Exec(deleteSQL, 4)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ = result.RowsAffected()
	fmt.Printf("削除された行数: %d\n", rowsAffected)

	// ========== プリペアドステートメント ==========
	fmt.Println("\n=== プリペアドステートメント ===")

	stmt, err := db.Prepare("SELECT id, name, email FROM users WHERE name LIKE ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows2, err := stmt.Query("%郎%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()

	fmt.Println("名前に「郎」を含むユーザー:")
	for rows2.Next() {
		var id int
		var name, email string
		rows2.Scan(&id, &name, &email)
		fmt.Printf("  %s (%s)\n", name, email)
	}

	// ========== トランザクション ==========
	fmt.Println("\n=== トランザクション ===")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// トランザクション内で操作
	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "五郎", "goro@example.com")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "六郎", "rokuro@example.com")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// コミット
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("トランザクション完了")

	// ========== 集計関数 ==========
	fmt.Println("\n=== 集計関数 ===")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ユーザー総数: %d\n", count)

	// ========== エラーハンドリング ==========
	fmt.Println("\n=== エラーハンドリング ===")

	// 重複エラー
	_, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "重複", "taro@example.com")
	if err != nil {
		fmt.Printf("重複エラー: %v\n", err)
	}

	// 存在しないレコード
	err = db.QueryRow("SELECT id, name FROM users WHERE id = ?", 9999).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		fmt.Println("レコードが見つかりません")
	}

	// ========== 最終確認 ==========
	fmt.Println("\n=== 最終確認 ===")

	rows3, _ := db.Query("SELECT id, name, email FROM users ORDER BY id")
	defer rows3.Close()

	fmt.Println("全ユーザー（最終）:")
	for rows3.Next() {
		var id int
		var name, email string
		rows3.Scan(&id, &name, &email)
		fmt.Printf("  ID=%d, Name=%s, Email=%s\n", id, name, email)
	}

	// ========== クリーンアップ ==========
	fmt.Println("\n=== クリーンアップ ===")

	_, err = db.Exec("DROP TABLE users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("テーブル削除完了")
}

/*
【実行方法】
# SQLiteドライバをインストール
go get github.com/mattn/go-sqlite3

# 実行
go run 01_sqlite.go

【重要な概念】
1. sql.Open でデータベース接続
2. defer で確実にクローズ
3. Exec: INSERT, UPDATE, DELETE
4. Query: 複数行取得
5. QueryRow: 単一行取得

【CRUD操作】
Create: INSERT INTO ... VALUES (?, ?)
Read:   SELECT ... FROM ... WHERE ...
Update: UPDATE ... SET ... WHERE ...
Delete: DELETE FROM ... WHERE ...

【プレースホルダー】
SQLite: ?
PostgreSQL: $1, $2, ...
MySQL: ?

【エラーハンドリング】
sql.ErrNoRows: レコードが見つからない
その他: 接続エラー、構文エラー等

【トランザクション】
tx, _ := db.Begin()
tx.Exec(...)
tx.Commit() または tx.Rollback()

【ベストプラクティス】
1. defer で rows.Close()
2. プリペアドステートメントで効率化
3. トランザクションで整合性を保つ
4. エラーハンドリングを適切に
5. SQL インジェクション対策（プレースホルダー使用）

【次のステップ】
02_postgres.go で PostgreSQL を学びましょう
*/
