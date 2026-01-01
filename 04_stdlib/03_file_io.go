package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
【学習ポイント】
1. ファイルの読み書き
2. バッファ付きIO
3. ディレクトリ操作
4. ファイル情報の取得
5. エラーハンドリング
*/

func main() {
	// ========== ファイルへの書き込み ==========
	fmt.Println("=== ファイルへの書き込み ===")

	// 方法1: os.WriteFile（簡単）
	content := []byte("Hello, Go!\nこれはテストファイルです。\n")
	err := os.WriteFile("test1.txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test1.txt に書き込み完了")

	// 方法2: os.Create + Write
	file, err := os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Line 1\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString("Line 2\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test2.txt に書き込み完了")

	// 方法3: バッファ付き書き込み
	file3, err := os.Create("test3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	writer := bufio.NewWriter(file3)
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(writer, "行 %d\n", i)
	}
	writer.Flush() // バッファをフラッシュ
	fmt.Println("test3.txt に書き込み完了")

	// ========== ファイルからの読み込み ==========
	fmt.Println("\n=== ファイルからの読み込み ===")

	// 方法1: os.ReadFile（簡単）
	data, err := os.ReadFile("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("test1.txt の内容:\n%s\n", string(data))

	// 方法2: os.Open + Read
	file4, err := os.Open("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file4.Close()

	buf := make([]byte, 1024)
	n, err := file4.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("test2.txt の内容:\n%s\n", string(buf[:n]))

	// 方法3: バッファ付き読み込み（行単位）
	file5, err := os.Open("test3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file5.Close()

	fmt.Println("test3.txt の内容（行単位）:")
	scanner := bufio.NewScanner(file5)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// ========== ファイルへの追記 ==========
	fmt.Println("\n=== ファイルへの追記 ===")

	file6, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file6.Close()

	_, err = file6.WriteString("追加された行\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test1.txt に追記完了")

	// ========== ファイル情報の取得 ==========
	fmt.Println("\n=== ファイル情報の取得 ===")

	fileInfo, err := os.Stat("test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ファイル名: %s\n", fileInfo.Name())
	fmt.Printf("サイズ: %d バイト\n", fileInfo.Size())
	fmt.Printf("パーミッション: %s\n", fileInfo.Mode())
	fmt.Printf("更新日時: %s\n", fileInfo.ModTime())
	fmt.Printf("ディレクトリ: %t\n", fileInfo.IsDir())

	// ========== ファイルの存在確認 ==========
	fmt.Println("\n=== ファイルの存在確認 ===")

	if _, err := os.Stat("test1.txt"); err == nil {
		fmt.Println("test1.txt は存在します")
	} else if os.IsNotExist(err) {
		fmt.Println("test1.txt は存在しません")
	} else {
		log.Fatal(err)
	}

	// ========== ファイルのコピー ==========
	fmt.Println("\n=== ファイルのコピー ===")

	err = copyFile("test1.txt", "test1_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test1.txt を test1_copy.txt にコピー完了")

	// ========== ファイルの移動/リネーム ==========
	fmt.Println("\n=== ファイルの移動/リネーム ===")

	err = os.Rename("test1_copy.txt", "test1_renamed.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test1_copy.txt を test1_renamed.txt にリネーム完了")

	// ========== ファイルの削除 ==========
	fmt.Println("\n=== ファイルの削除 ===")

	err = os.Remove("test1_renamed.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test1_renamed.txt を削除完了")

	// ========== ディレクトリの作成 ==========
	fmt.Println("\n=== ディレクトリの作成 ===")

	err = os.Mkdir("testdir", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	fmt.Println("testdir ディレクトリを作成")

	// 階層的にディレクトリを作成
	err = os.MkdirAll("testdir/subdir1/subdir2", 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("testdir/subdir1/subdir2 を作成")

	// ========== ディレクトリ内のファイル一覧 ==========
	fmt.Println("\n=== ディレクトリ内のファイル一覧 ===")

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("カレントディレクトリの内容:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("[DIR]  %s\n", entry.Name())
		} else {
			fmt.Printf("[FILE] %s\n", entry.Name())
		}
	}

	// ========== ディレクトリの再帰的な走査 ==========
	fmt.Println("\n=== ディレクトリの再帰的な走査 ===")

	err = filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("[DIR]  %s\n", path)
		} else {
			fmt.Printf("[FILE] %s\n", path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// ========== 一時ファイル ==========
	fmt.Println("\n=== 一時ファイル ===")

	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // クリーンアップ
	defer tmpFile.Close()

	fmt.Printf("一時ファイル: %s\n", tmpFile.Name())

	_, err = tmpFile.WriteString("一時データ\n")
	if err != nil {
		log.Fatal(err)
	}

	// ========== パス操作 ==========
	fmt.Println("\n=== パス操作 ===")

	path := "testdir/subdir1/file.txt"
	fmt.Printf("元のパス: %s\n", path)
	fmt.Printf("ディレクトリ: %s\n", filepath.Dir(path))
	fmt.Printf("ファイル名: %s\n", filepath.Base(path))
	fmt.Printf("拡張子: %s\n", filepath.Ext(path))

	// パスの結合
	joined := filepath.Join("testdir", "subdir1", "file.txt")
	fmt.Printf("結合されたパス: %s\n", joined)

	// 絶対パス
	absPath, err := filepath.Abs("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("絶対パス: %s\n", absPath)

	// ========== クリーンアップ ==========
	fmt.Println("\n=== クリーンアップ ===")

	// テストファイルを削除
	os.Remove("test1.txt")
	os.Remove("test2.txt")
	os.Remove("test3.txt")

	// ディレクトリを削除
	os.RemoveAll("testdir")

	fmt.Println("クリーンアップ完了")
}

// ファイルコピー関数
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

/*
【実行方法】
go run 03_file_io.go

【重要な概念】
1. defer で確実にファイルをクローズ
2. エラーハンドリングは必須
3. バッファ付きIOで効率化
4. パーミッションに注意

【ファイル操作】
os.Create(name)           // ファイル作成
os.Open(name)             // ファイルを開く（読み込み専用）
os.OpenFile(name, flag, perm) // フラグ指定で開く
os.WriteFile(name, data, perm) // 簡単な書き込み
os.ReadFile(name)         // 簡単な読み込み

【フラグ】
os.O_RDONLY  // 読み込み専用
os.O_WRONLY  // 書き込み専用
os.O_RDWR    // 読み書き
os.O_APPEND  // 追記
os.O_CREATE  // 存在しなければ作成
os.O_TRUNC   // 開く時に切り詰める

【パーミッション】
0644  // rw-r--r-- (ファイル)
0755  // rwxr-xr-x (ディレクトリ)

【ディレクトリ操作】
os.Mkdir(name, perm)      // ディレクトリ作成
os.MkdirAll(path, perm)   // 階層的に作成
os.ReadDir(name)          // ディレクトリ内容を読む
os.Remove(name)           // 削除
os.RemoveAll(path)        // 再帰的に削除

【ベストプラクティス】
1. defer file.Close() を必ず書く
2. エラーチェックを省略しない
3. 大きなファイルはバッファ付きIOを使う
4. 一時ファイルは os.CreateTemp を使う

【次のステップ】
04_time.go で時刻処理を学びましょう
*/
