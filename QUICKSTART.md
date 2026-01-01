# 🚀 クイックスタートガイド

## はじめに

このディレクトリには、Go言語の基礎から実践的なWebバックエンド開発まで、体系的に学習できる教材が含まれています。

## 📋 前提条件

- Go 1.21以上がインストールされていること
- テキストエディタまたはIDE（VS Code推奨）

## 🏃 最初の一歩

### 1. Hello World を実行

```bash
go run 01_basics/01_hello.go
```

出力:
```
Hello, Go!
Go is awesome!
私の名前は太郎で、25歳です。
太郎さんは25歳です
```

### 2. 各ファイルを順番に実行

```bash
# 基礎編
go run 01_basics/02_variables.go
go run 01_basics/03_types.go
go run 01_basics/04_operators.go
go run 01_basics/05_control_flow.go
go run 01_basics/06_arrays_slices.go
go run 01_basics/07_maps.go
go run 01_basics/08_pointers.go

# 関数とメソッド編
go run 02_functions/01_functions.go
go run 02_functions/02_structs.go
go run 02_functions/03_methods.go
go run 02_functions/04_interfaces.go
go run 02_functions/05_errors.go

# 並行処理編
go run 03_concurrency/01_goroutines.go
go run 03_concurrency/02_channels.go
go run 03_concurrency/03_sync.go

# 標準ライブラリ編
go run 04_stdlib/01_json.go
go run 04_stdlib/02_http_client.go

# Webバックエンド基礎編
go run 05_web_basics/01_http_server.go
# 別のターミナルで: curl http://localhost:8080/

go run 05_web_basics/02_middleware.go
# 別のターミナルで: curl -H "Authorization: Bearer secret-token" http://localhost:8080/with-middleware

# REST API編
go run 07_rest_api/01_rest_api.go
# 別のターミナルで: curl http://localhost:8080/api/users
```

## 🎯 学習の進め方

### ステップ1: 基礎を固める（1-2週間）
- `01_basics/` の全ファイルを実行
- コードを読んで理解する
- 自分で変更を加えて実験する

### ステップ2: 関数とメソッドを学ぶ（1週間）
- `02_functions/` の全ファイルを実行
- 構造体とメソッドの概念を理解
- インターフェースの使い方を習得

### ステップ3: 並行処理をマスター（1週間）
- `03_concurrency/` の全ファイルを実行
- GoroutineとChannelの使い方を理解
- 実際の並行処理パターンを学ぶ

### ステップ4: 標準ライブラリを使う（数日）
- `04_stdlib/` のファイルを実行
- JSONの扱い方を習得
- HTTP通信の基本を学ぶ

### ステップ5: Webバックエンドを構築（1-2週間）
- `05_web_basics/` のファイルを実行
- HTTPサーバーの作り方を学ぶ
- ミドルウェアパターンを理解

### ステップ6: REST APIを実装（1週間）
- `07_rest_api/` のファイルを実行
- CRUD操作を実装
- RESTful設計を理解

## 💡 効果的な学習方法

### 1. コードを読む
各ファイルには詳細なコメントが含まれています。まずはコードを読んで理解しましょう。

### 2. コードを実行する
```bash
go run <ファイル名>.go
```

### 3. コードを変更する
- 変数の値を変える
- 新しい関数を追加する
- エラーを意図的に起こしてみる

### 4. 自分で書く
- ファイルを見ずに同じコードを書いてみる
- 学んだ概念を使って新しいプログラムを作る

### 5. 問題を解く
各章の最後に「次のステップ」があります。それに従って進みましょう。

## 🛠️ よく使うコマンド

```bash
# プログラムを実行
go run main.go

# ビルド（実行ファイルを作成）
go build main.go

# フォーマット
go fmt ./...

# テスト
go test ./...

# 依存関係の管理
go mod tidy

# レースコンディションの検出
go run -race main.go
```

## 📖 各章の概要

### 01_basics/ - 基礎編
Go言語の基本文法を学びます。変数、型、制御構文、配列、マップ、ポインタなど。

**学習時間**: 1-2週間

### 02_functions/ - 関数とメソッド編
関数、構造体、メソッド、インターフェース、エラーハンドリングを学びます。

**学習時間**: 1週間

### 03_concurrency/ - 並行処理編
Goroutine、Channel、sync パッケージを使った並行処理を学びます。

**学習時間**: 1週間

### 04_stdlib/ - 標準ライブラリ編
JSON処理、HTTP通信など、実用的な標準ライブラリの使い方を学びます。

**学習時間**: 数日

### 05_web_basics/ - Webバックエンド基礎編
HTTPサーバーの構築、ルーティング、ミドルウェアを学びます。

**学習時間**: 1-2週間

### 07_rest_api/ - REST API編
完全なREST APIの実装方法を学びます。

**学習時間**: 1週間

## 🎓 学習後の次のステップ

### データベース統合
- SQLiteの基本
- PostgreSQL接続
- GORM（ORM）の使用

### 認証・認可
- JWT認証の実装
- セッション管理
- パスワードハッシュ化

### テスト
- ユニットテスト
- テーブル駆動テスト
- モックの使用

### デプロイ
- Docker化
- クラウドデプロイ（AWS, GCP, Azure）
- CI/CD

### フレームワーク
- Gin
- Echo
- Fiber

## 🆘 困ったときは

### エラーが出たら
1. エラーメッセージをよく読む
2. Google で検索する
3. 公式ドキュメントを確認する
4. Stack Overflow で質問する

### 理解できないときは
1. コードを小さく分解する
2. print文でデバッグする
3. 公式のGo Tourを見る
4. コミュニティで質問する

## 📚 推奨リソース

- [Go公式サイト](https://go.dev/)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

## 🎯 目標設定

このカリキュラムを完了すると:

✅ Go言語の基本文法を理解し、活用できる
✅ 並行処理を使った効率的なプログラムが書ける
✅ REST APIを設計・実装できる
✅ Webアプリケーションのバックエンドを構築できる

頑張ってください！ 🚀

---

**質問やフィードバックがあれば、Issueを作成してください。**
