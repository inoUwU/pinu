# データベーススキーマドキュメント

このディレクトリには、Pinuプロジェクトのデータベーススキーマに関するドキュメントが含まれています。

## ファイル一覧

### schema.dbml
**DBML (Database Markup Language)** 形式で記述されたデータベーススキーマ定義ファイルです。

#### DBMLの特徴
- **可読性**: 人間が読みやすく、編集しやすい形式
- **ツールサポート**: [dbdiagram.io](https://dbdiagram.io/) などのツールで視覚化可能
- **自動生成**: SQL DDLやドキュメントの自動生成が可能
- **バージョン管理**: Gitでの差分管理が容易

#### schema.dbmlの使用方法

##### 1. オンラインでER図を表示
1. [dbdiagram.io](https://dbdiagram.io/) にアクセス
2. `schema.dbml` の内容をコピー&ペースト
3. 自動的にER図が生成されます

##### 2. SQL DDLの生成
[DBML CLI](https://www.dbml.org/cli/) を使用してPostgreSQL用のDDLを生成できます：

```bash
# DBML CLIのインストール
npm install -g @dbml/cli

# PostgreSQL DDLの生成
dbml2sql schema.dbml --postgres -o schema.sql
```

##### 3. ドキュメント生成
DBMLファイルから自動的にドキュメントを生成するツールもあります：
- [dbdocs.io](https://dbdocs.io/) - オンラインデータベースドキュメント

### er_diagram.md
Mermaid形式で記述された従来のER図ドキュメントです。
GitHubやMarkdownビューアーで表示可能ですが、DBMLの方がより詳細で扱いやすいため、今後は `schema.dbml` を参照することを推奨します。

## データベース構造

### 主要テーブル

#### ユーザー管理
- **users**: 従業員・管理者アカウント
- **sessions**: ユーザーのログインセッション

#### メニュー管理
- **categories**: メニューカテゴリ
- **menus**: メニュー項目
- **menu_options**: メニューオプション（トッピングなど）
- **menu_option_assignments**: メニューとオプションの関連

#### テーブル・注文管理
- **tables**: 店舗の物理テーブル
- **table_sessions**: テーブルの利用セッション（QR入店〜会計まで）
- **order_groups**: 注文グループ
- **order_items**: 注文明細
- **order_item_options**: 注文明細のオプション

#### 設定
- **settings**: アプリケーション全体の設定

### ENUM型

#### table_status
テーブルの状態を表します：
- `available`: 空席
- `occupied`: 使用中
- `billing`: 会計待ち

#### order_status
注文の状態を表します：
- `pending`: 受付待ち
- `preparing`: 調理中
- `served`: 提供済み
- `cancelled`: キャンセル

#### order_group_status
注文グループの状態を表します：
- `open`: オープン（追加注文可能）
- `closed`: クローズ（確定済み）
- `cancelled`: キャンセル

## 実装ファイル

実際のデータベース初期化スクリプトは以下にあります：
- `/database/init/init.sql`: PostgreSQL初期化SQL

## 更新手順

データベーススキーマを変更する場合：

1. **init.sqlを更新**: 実際のSQL DDLを変更
2. **schema.dbmlを更新**: DBMLファイルに変更を反映
3. **必要に応じてer_diagram.mdを更新**: Mermaid形式も維持する場合

## 参考リンク

- [DBML公式ドキュメント](https://www.dbml.org/docs/)
- [dbdiagram.io](https://dbdiagram.io/) - オンラインER図ツール
- [DBML CLI](https://www.dbml.org/cli/) - コマンドラインツール
- [dbdocs.io](https://dbdocs.io/) - オンラインドキュメント生成
