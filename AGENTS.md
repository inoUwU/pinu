# AGENTS.md

このファイルは、GitHub CopilotやOpenAI Codexなどの AI コーディングエージェント向けのプロジェクトガイドです。

## プロジェクト概要

**Pinu** は、町中華などの小規模個人経営店舗向けのQRオーダーシステムです。店主一人でも運用可能な、シンプルで直感的なフルスタックアプリケーションです。

- **バックエンド**: Go + Fiber (ポート&アダプターアーキテクチャ)
- **フロントエンド**: Next.js + TypeScript + Shadcn UI + Tailwind CSS
- **データベース**: PostgreSQL
- **開発環境**: Docker + Task runner

### プロジェクトスローガン

**日本語**: 「PINU ― piっとメニュー、らくらくオーダー！」  
**英語**: "PINU — Tap the menu, order with ease!"

## 開発環境のセットアップ

### 前提条件

- [Docker](https://www.docker.com/get-started)
- [Task](https://taskfile.dev) (インストール: `sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d`)
- Go 1.22以降
- Node.js 18以降 & pnpm

### 初期セットアップ

```bash
# Taskツールをインストール（初回のみ）
task install:task

# プロジェクトのセットアップ（環境変数、依存関係）
task setup
```

### 開発環境の起動

```bash
# 開発環境を起動（データベース、バックエンド、フロントエンド）
task start

# 開発環境一式を起動（adminer含む）
task dev
```

アクセスURL:
- **バックエンド**: http://localhost:8000
- **フロントエンド**: http://localhost:3000
- **Adminer (DB管理)**: http://localhost:8080

### その他の便利なコマンド

```bash
# 開発環境を停止
task stop

# データベースをリセット
task db:reset

# 完全クリーンアップ
task clean

# 完全に新しい状態で起動
task fresh
```

## ビルド・テスト・Lintコマンド

### 全体

```bash
# 全体をビルド
task build

# 全体のテストを実行
task test

# 全体のLintを実行
task lint

# デプロイ準備（ビルド + テスト）
task deploy:prep
```

### バックエンド（Go）

```bash
# 依存関係のインストール
cd backend && go mod tidy && go mod download

# 開発モードで起動
cd backend && go run cmd/server/main.go

# ビルド
cd backend && go build -o ../bin/backend cmd/server/main.go

# テスト実行
cd backend && go test ./... -v

# Lint実行（golangci-lint必要）
cd backend && golangci-lint run
```

### フロントエンド（Next.js + TypeScript）

```bash
# 依存関係のインストール
cd frontend && pnpm install

# 開発モードで起動
cd frontend && pnpm dev

# ビルド
cd frontend && pnpm build

# テスト実行
cd frontend && pnpm test

# Lint実行
cd frontend && pnpm lint

# フォーマット（Biome）
cd frontend && pnpm format
```

## アーキテクチャとコードスタイル

### バックエンドアーキテクチャ

**ポート&アダプター（ヘキサゴナルアーキテクチャ）** を採用しています。

```
Controllers (Primary Adapter)
    ↓
UseCases (Application Layer)
    ↓
Repository Interface (Domain Port)
    ↓
Repository Implementation (Secondary Adapter)
    ↓
Database
```

#### レイヤー別の責任

1. **Controllers** (`app/handlers/`)
   - HTTPリクエスト/レスポンスの処理
   - ユースケースの呼び出し
   - プレゼンテーション層

2. **UseCases** (`app/usecases/`)
   - ビジネスロジックの実装
   - バリデーション処理
   - アプリケーション固有のロジック
   - Input/Output DTOの使用

3. **Domain** (`app/domain/`)
   - エンティティ定義 (`entities/`)
   - リポジトリインターフェース (`repositories/`)

4. **Infrastructure** (`app/infrastructure/`)
   - リポジトリ実装 (`repositories/`)
   - データベース接続・操作

#### アーキテクチャの重要原則

- **依存性逆転の原則**: 内側の層は外側の層に依存しない
- **サービス層の削除**: 不要な抽象化を排除し、ユースケースに直接実装
- **ポインタ型の使用**: サービス・リポジトリはポインタ型で統一
- **インターフェース**: `*Interface` ではなく `Interface` 型を使用

### フロントエンドアーキテクチャ

- **Next.js App Router** を使用
- **Turbo Repo** によるモノレポ構成
- **apps/**: 個別アプリケーション（client, admin）
- **packages/**: 共有パッケージ・コンポーネント

### コーディング規約（全般）

詳細は `.github/instructions/general.instructions.md` を参照してください。

主な原則:
- **命名**: 意味のある変数名・関数名を使用
- **コメント**: 簡潔で具体的に記述
- **マジックナンバー禁止**: 定数として定義
- **可読性優先**: コードの読みやすさを重視
- **DRY原則**: コードの重複を避ける
- **SOLID・KISS原則** に従う
- **型安全性**: 型定義・アノテーションを積極的に使用
- **エラーハンドリング**: 例外や失敗ケースを慎重に扱う
- **単一責任の原則**: 関数・クラスは単一の責任を持つ
- **セキュリティ意識**: インジェクション、XSSなどのリスクを考慮

### Go言語の特記事項

- **構造体名**: 公開する構造体は大文字で始める（例: `UserUsecaseImpl`）
- **インターフェース型**: ポインタを使わない（例: `repositories.IUserRepository`）
- **依存性注入**: DIコンテナを使用 (`app/middleware/injection.go`)

### TypeScript/Next.jsの特記事項

- **Strict モード**: TypeScript strict mode を使用
- **Validation**: Zodを使用
- **Form**: Conformを使用
- **API通信**: SWRを使用
- **UI**: Shadcn UI + Tailwind CSS

## データベース

### 接続情報（デフォルト）

- **データベース名**: `pinu`
- **ユーザー名**: `pinu_user`
- **パスワード**: `pinu_pass`
- **ポート**: `5432`

環境変数は `.env` で設定可能（`.env.example` を参照）。

### データベース操作

```bash
# PostgreSQL起動
task db:start

# PostgreSQL停止
task db:stop

# データベースリセット（全データ削除）
task db:reset

# ログ表示
task db:logs
```

### ER図とスキーマ

詳細は `.docs/database/er_diagram.md` を参照してください。

主なテーブル:
- `tables`: 店舗内の物理テーブル管理
- `table_sessions`: テーブルのセッション（QR入店〜会計）
- `order_groups`: セッション配下の注文グループ
- `order_items`: 各注文項目
- `menus`: メニュー情報
- `categories`: メニューカテゴリ
- `users`: 従業員・管理者

## コミット・PRガイドライン

### Conventional Commits を使用

詳細は `.docs/commit-guidelines.md` を参照してください。

#### フォーマット

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

#### 主な type とgitmoji

| Emoji | Type      | 説明                       |
|-------|-----------|----------------------------|
| ✨    | feat      | 新機能追加                 |
| 🐛    | fix       | バグ修正                   |
| 📝    | docs      | ドキュメント追加・更新     |
| 🎨    | style     | コードスタイル改善         |
| ♻️    | refactor  | リファクタリング           |
| ⚡️    | perf      | パフォーマンス改善         |
| ✅    | test      | テスト追加・更新           |
| 🔧    | chore     | ビルド・ツール関連         |

#### Breaking Change

Breaking Changeがある場合は `!` を追加するか、フッターに `BREAKING CHANGE:` を記載。

```
feat(api)!: drop support for legacy v1 endpoints

BREAKING CHANGE: API clients must now use the v2 endpoints.
```

### PRのルール

- コミットメッセージはConventional Commitsに従う
- テストが通ることを確認してからPR作成
- Lintエラーがないことを確認
- PRには関連するissue番号を記載（Closes: #XX）

## プロジェクト構造

```
.
├── backend/                 # Goバックエンドコード
│   ├── app/
│   │   ├── handlers/        # Controllers (HTTPハンドラー)
│   │   ├── usecases/        # Application Layer (ビジネスロジック)
│   │   ├── domain/          # Domain Layer
│   │   │   ├── entities/    # エンティティ
│   │   │   └── repositories/# リポジトリインターフェース
│   │   ├── infrastructure/  # Infrastructure Layer
│   │   │   └── repositories/# リポジトリ実装
│   │   └── middleware/      # DI、認証など
│   ├── cmd/server/          # エントリポイント
│   ├── go.mod
│   └── go.sum
├── frontend/                # Next.js フロントエンド（Turbo Repo）
│   ├── apps/
│   │   ├── client/          # 顧客向けアプリ
│   │   └── admin/           # 管理者向けアプリ
│   ├── packages/            # 共有パッケージ
│   ├── package.json
│   └── turbo.json
├── database/                # データベース関連ファイル
│   ├── init/                # 初期化SQL
│   └── postgresql.conf      # PostgreSQL設定
├── .docs/                   # プロジェクトドキュメント
│   ├── current-architecture.md
│   ├── architecture-decisions.md
│   ├── technology_selection.md
│   ├── commit-guidelines.md
│   ├── spec/                # 要件定義・仕様書
│   └── database/            # ER図
├── docker/                  # Dockerfiles
├── compose.yml              # Docker Compose設定
├── Taskfile.yml             # Task runner設定
├── .env                     # 環境変数（git管理外）
└── .env.example             # 環境変数テンプレート
```

## ドキュメント参照

プロジェクトの詳細な情報は `.docs/` ディレクトリを参照してください:

- **アーキテクチャ**: `.docs/current-architecture.md`
- **アーキテクチャ決定記録**: `.docs/architecture-decisions.md`
- **技術選定**: `.docs/technology_selection.md`
- **コミットガイドライン**: `.docs/commit-guidelines.md`
- **要件定義**: `.docs/spec/requirements_definition.md`
- **機能一覧**: `.docs/spec/feature_draft.md`
- **ER図**: `.docs/database/er_diagram.md`

## セキュリティ

- シークレット情報をコミットしない
- `.env` ファイルは `.gitignore` で除外済み
- 環境変数は `.env.example` を参考に `.env` を作成
- パスワードはハッシュ化して保存（bcrypt使用）
- SQLインジェクション対策としてORMまたはプリペアドステートメントを使用

## トラブルシューティング

### データベース接続エラー

```bash
# PostgreSQLの状態確認
docker ps
task db:logs

# データベースの再起動
task db:stop
task db:start
```

### ポートが既に使用されている

デフォルトポート:
- `5432`: PostgreSQL
- `8000`: バックエンド
- `3000`: フロントエンド

`.env` で変更可能:
```
DB_PORT=5433
BACKEND_PORT=8001
FRONTEND_PORT=3001
```

### 依存関係のエラー

```bash
# バックエンド
cd backend && go mod tidy && go mod download

# フロントエンド
cd frontend && rm -rf node_modules pnpm-lock.yaml && pnpm install
```

## 追加情報

- **ライセンス**: MIT License
- **言語**: コードは英語、ドキュメント・コメントは日本語を推奨
- **対象ユーザー**: 町中華などの小規模個人経営店舗の店主
- **UI/UX**: 高齢の店主でも直感的に操作できるシンプル設計を重視
