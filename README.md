![Pinu Logo](.github/assets/readme_logo.webp)

# 🚧Pinu

ポートフォリオプロジェクト用のフルスタック開発環境

## 概要

このプロジェクトは、バックエンド（Go）とフロントエンドを含むポートフォリオアプリケーションです。Task ランナーと Docker を利用して簡単に開発環境を構築できます。

## 前提条件

- [Docker](https://www.docker.com/get-started)
- [Task](https://taskfile.dev) (インストールコマンド: `sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d`)
- Go 言語環境
- Node.js と npm

## セットアップ

初めてプロジェクトをセットアップする場合：

```bash
# Task ツールをインストール（初回のみ）
task install:task

# プロジェクトをセットアップ（環境変数、依存関係のインストール）
task setup
```

## 開発環境の起動

```bash
# 開発環境を起動（データベース、バックエンド、フロントエンド）
task start

# 開発環境一式（adminer含む）を起動
task dev
```

開発環境起動後、以下のURLでアクセスできます：

- バックエンド: <http://localhost:8000>
- フロントエンド: <http://localhost:3000>
- Adminer (DB管理): <http://localhost:8080>

## 主要なタスク

| タスク | 説明 |
|--------|------|
| `task setup` | 初期セットアップ（環境変数、依存関係のインストール） |
| `task start` | 開発環境を起動（DB + バックエンド + フロントエンド） |
| `task dev` | 開発環境一式起動（adminer含む） |
| `task stop` | 開発環境を停止 |
| `task clean` | データベースを含めて完全に削除 |
| `task test` | 全体のテストを実行 |
| `task build` | 全体をビルド |
| `task deploy:prep` | デプロイ準備（ビルド + テスト） |
| `task fresh` | 完全に新しい状態で環境を起動 |

すべてのタスクを表示するには:

```bash
task --list
```

## データベース

PostgreSQL データベースは Docker コンテナとして実行され、以下のデフォルト設定が適用されます：

- データベース名: `portfolio`
- ユーザー名: `portfolio_user`
- パスワード: `portfolio_pass`
- ポート: `5432`

データベース関連のタスク：

```bash
task db:start    # PostgreSQLを起動
task db:stop     # PostgreSQLを停止
task db:reset    # データベースをリセット（全データ削除）
task db:logs     # PostgreSQLのログを表示
```

## プロジェクト構造

```bash
.
├── backend/         # Goバックエンドコード
├── frontend/        # フロントエンドコード
├── database/        # データベース関連ファイル
│   ├── init/        # 初期化SQL
│   └── postgresql.conf  # PostgreSQL設定
├── .docs/           # プロジェクトドキュメント
│   ├── ACCESSIBILITY_INDEX.md  # アクセシビリティドキュメント索引
│   └── accessibility-*.md      # アクセシビリティ関連ドキュメント
├── bin/             # ビルド生成物
├── compose.yml      # Docker Compose設定
├── Taskfile.yml     # Taskランナー設定
└── .env             # 環境変数（セットアップ時に作成）
```

## ドキュメント

### アクセシビリティ

フロントエンドのアクセシビリティに関する包括的なドキュメントは [.docs/ACCESSIBILITY_INDEX.md](.docs/ACCESSIBILITY_INDEX.md) を参照してください。

- 📊 [詳細監査レポート](.docs/accessibility-audit-report.md) - 42件の問題と推奨対応
- ✅ [実装チェックリスト](.docs/accessibility-improvements-checklist.md) - すぐに使えるコードスニペット
- 📘 [クイックリファレンス](.docs/accessibility-quick-reference.md) - 開発時のベストプラクティス
- 📝 [エグゼクティブサマリー](.docs/accessibility-audit-summary.md) - 概要と対応計画

## ライセンス

[MIT License](LICENSE)
