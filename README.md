# Pinu

Pinu is toy ordering system for restaurants.

## 🚀 クイックスタート

### 前提条件

- [Task](https://taskfile.dev/) のインストール

# Task未インストールの場合

```bash
task install:task
```

1. 環境のセットアップ
   ```bash
   bashtask setup
   ```
2. 開発環境の起動
   ```bash
   task start # 一括起動
   ```

# または個別起動

```bash
task db:start
task backend:dev
task frontend:dev
```

3. アクセス

- フロントエンド: http://localhost:3000
- バックエンド: http://localhost:8000
- DB管理画面: http://localhost:8080

🛠️ 開発用コマンド
bash# 利用可能なコマンド一覧
task

# 開発環境一式起動（adminer含む）

task dev

# テスト実行

task test

# リント実行

task lint

# ビルド

task build

# 完全リセット

task fresh

# ヘルスチェック

task health
