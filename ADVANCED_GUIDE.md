# 📚 追加学習コンテンツ

このドキュメントでは、追加で作成された学習コンテンツについて説明します。

## 🆕 新規追加されたファイル

### 並行処理編（03_concurrency/）

#### 04_context.go - Context の使い方
**学習内容:**
- Context の基本概念
- WithCancel, WithTimeout, WithDeadline
- WithValue での値の受け渡し
- 実用的なパターン（HTTPリクエスト、データベースクエリ）

**実行:**
```bash
go run 03_concurrency/04_context.go
```

**重要ポイント:**
- Context は関数の第一引数として渡す
- 必ず cancel() を呼ぶ（defer で）
- Done() チャネルで完了を検知

#### 05_patterns.go - 並行処理パターン
**学習内容:**
- ファンアウト/ファンイン
- 動的ワーカープール
- パイプラインパターン
- セマフォパターン
- レート制限

**実行:**
```bash
go run 03_concurrency/05_patterns.go
```

**実用例:**
- 並行ダウンロード
- 並行マップ処理
- タスクの並列実行

---

### 標準ライブラリ編（04_stdlib/）

#### 03_file_io.go - ファイル操作
**学習内容:**
- ファイルの読み書き
- バッファ付きIO
- ディレクトリ操作
- パス操作
- 一時ファイル

**実行:**
```bash
go run 04_stdlib/03_file_io.go
```

**重要ポイント:**
- defer file.Close() を必ず書く
- エラーハンドリングは必須
- 大きなファイルはバッファ付きIOを使う

#### 04_time.go - 時刻処理
**学習内容:**
- 時刻の取得と生成
- フォーマットとパース
- 時刻の計算と比較
- Timer と Ticker
- タイムゾーン

**実行:**
```bash
go run 04_stdlib/04_time.go
```

**フォーマット参照日時:**
```
Mon Jan 2 15:04:05 MST 2006
```

---

### データベース編（06_database/）- 新規追加

#### 01_sqlite.go - SQLite 基礎
**学習内容:**
- データベース接続
- CRUD操作（Create, Read, Update, Delete）
- トランザクション
- プリペアドステートメント
- エラーハンドリング

**事前準備:**
```bash
go get github.com/mattn/go-sqlite3
```

**実行:**
```bash
go run 06_database/01_sqlite.go
```

**重要ポイント:**
- defer で rows.Close()
- プレースホルダーで SQL インジェクション対策
- トランザクションで整合性を保つ

---

### REST API編（07_rest_api/）

#### 02_advanced_api.go - 高度なREST API
**学習内容:**
- 認証システム（簡易JWT）
- バリデーション
- ページネーション
- フィルタリング
- エラーレスポンスの統一

**実行:**
```bash
go run 07_rest_api/02_advanced_api.go
```

**エンドポイント:**
```
POST   /api/auth/login      - ログイン
POST   /api/auth/register   - ユーザー登録
GET    /api/users           - ユーザー一覧（ページネーション）
GET    /api/posts           - 投稿一覧（フィルタ）
POST   /api/posts           - 投稿作成
PUT    /api/posts/{id}      - 投稿更新
DELETE /api/posts/{id}      - 投稿削除
```

**テスト例:**
```bash
# ログイン
curl -X POST http://localhost:8080/api/auth/login \
  -d '{"email":"taro@example.com","password":"password123"}' \
  -H "Content-Type: application/json"

# ページネーション
curl "http://localhost:8080/api/users?page=1&per_page=10"

# フィルタリング
curl "http://localhost:8080/api/posts?user_id=1&published=true"
```

---

## 📊 学習コンテンツ一覧

### 合計ファイル数: **27ファイル**

| カテゴリ | ファイル数 | 学習時間目安 |
|---------|-----------|------------|
| 基礎編 | 8 | 1-2週間 |
| 関数とメソッド編 | 5 | 1週間 |
| 並行処理編 | 5 | 1-2週間 |
| 標準ライブラリ編 | 4 | 1週間 |
| Webバックエンド基礎編 | 2 | 1週間 |
| データベース編 | 1 | 数日 |
| REST API編 | 2 | 1週間 |

**総学習時間: 約6-8週間**

---

## 🎯 学習の進め方（推奨）

### 週1-2: 基礎固め
```bash
# 基礎編を順番に実行
go run 01_basics/01_hello.go
go run 01_basics/02_variables.go
# ... 08_pointers.go まで
```

### 週3: 関数とメソッド
```bash
# 関数とメソッド編
go run 02_functions/01_functions.go
# ... 05_errors.go まで
```

### 週4-5: 並行処理をマスター
```bash
# 並行処理編（重要！）
go run 03_concurrency/01_goroutines.go
go run 03_concurrency/02_channels.go
go run 03_concurrency/03_sync.go
go run 03_concurrency/04_context.go
go run 03_concurrency/05_patterns.go
```

### 週6: 標準ライブラリ
```bash
# 標準ライブラリ編
go run 04_stdlib/01_json.go
go run 04_stdlib/02_http_client.go
go run 04_stdlib/03_file_io.go
go run 04_stdlib/04_time.go
```

### 週7: Webバックエンド
```bash
# Webバックエンド基礎編
go run 05_web_basics/01_http_server.go
go run 05_web_basics/02_middleware.go
```

### 週8: データベースとREST API
```bash
# データベース編
go run 06_database/01_sqlite.go

# REST API編
go run 07_rest_api/01_rest_api.go
go run 07_rest_api/02_advanced_api.go
```

---

## 💡 実践的な学習のヒント

### 1. コードを変更してみる
各ファイルを実行したら、以下を試してみましょう:
- 変数の値を変える
- 新しい関数を追加する
- エラーを意図的に起こしてみる
- 自分なりの機能を追加する

### 2. 小さなプロジェクトを作る
学んだことを組み合わせて:
- TODOアプリのバックエンド
- 簡単なブログシステム
- ユーザー管理API
- ファイルアップロードサーバー

### 3. エラーから学ぶ
- エラーメッセージをよく読む
- Google で検索する
- 公式ドキュメントを確認する

### 4. コミュニティを活用
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://www.reddit.com/r/golang/)
- Stack Overflow

---

## 🔧 よく使うコマンド

```bash
# プログラムを実行
go run main.go

# ビルド
go build main.go

# フォーマット
go fmt ./...

# 依存関係の管理
go mod init myproject
go mod tidy

# レースコンディションの検出
go run -race main.go

# テスト
go test ./...

# ベンチマーク
go test -bench=.
```

---

## 📖 次のステップ

### 学習を深める
1. **テストを書く** - testing パッケージ
2. **フレームワークを使う** - Gin, Echo, Fiber
3. **ORM を使う** - GORM
4. **Docker化する**
5. **デプロイする** - AWS, GCP, Heroku

### 推奨ライブラリ
- **Webフレームワーク**: Gin, Echo, Fiber
- **ORM**: GORM, sqlx
- **認証**: golang-jwt/jwt
- **バリデーション**: go-playground/validator
- **環境変数**: godotenv
- **ロギング**: zap, logrus
- **テスト**: testify

---

## ✅ 学習チェックリスト

### 並行処理編
- [ ] Goroutine を使える
- [ ] Channel でデータをやり取りできる
- [ ] Context を理解した
- [ ] ワーカープールを実装できる
- [ ] パイプラインパターンを理解した

### 標準ライブラリ編
- [ ] ファイルの読み書きができる
- [ ] 時刻のフォーマットができる
- [ ] Timer/Ticker を使える

### データベース編
- [ ] データベースに接続できる
- [ ] CRUD操作ができる
- [ ] トランザクションを使える

### REST API編
- [ ] バリデーションを実装できる
- [ ] ページネーションを実装できる
- [ ] フィルタリングを実装できる
- [ ] エラーハンドリングができる

---

頑張ってください! 🚀
