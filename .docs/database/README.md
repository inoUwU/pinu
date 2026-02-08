# データベーススキーマドキュメント

このディレクトリには、Pinuプロジェクトのデータベーススキーマに関するドキュメントが含まれています。

## 概要

PinuはQR Order Systemのためのデータベースで、PostgreSQL 17.5を使用しています。スキーマはDBML (Database Markup Language) 形式で `schema.dbml` に定義されており、正確な型定義、制約、インデックス、リレーションシップを含んでいます。

## スキーマファイル

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

## データベース構造

### テーブル一覧と詳細説明

#### 設定テーブル

##### settings

アプリケーション全体の設定を保存します（例：ロゴのURL、店舗のテーブル総数）。

**主要カラム:**

- `key` (varchar(255), PK): 設定キー
- `value` (text): 設定値
- `created_at` (timestamptz): 作成時間
- `updated_at` (timestamptz): 更新時間

#### ユーザー管理テーブル

##### users

従業員・管理者などのユーザーを表します。

**主要カラム:**

- `user_id` (uuid, PK): ユーザーID
- `login_id` (varchar(255), UNIQUE): ログインID
- `password_hash` (varchar(255)): パスワードハッシュ
- `password_salt` (varchar(255)): パスワードソルト
- `name` (varchar(255)): 表示名
- `is_admin` (boolean): 管理者フラグ
- `created_at` (timestamptz): 作成時間
- `updated_at` (timestamptz): 更新時間

##### sessions

ユーザーのログインセッション情報を管理します。

**主要カラム:**

- `session_id` (varchar(255), PK): セッションID
- `user_id` (uuid, FK → users): ユーザーID
- `refresh_token` (varchar(512)): リフレッシュトークン
- `is_revoked` (boolean): 無効化フラグ
- `created_at` (timestamptz): 作成時間
- `expires_at` (timestamptz): 有効期限

#### メニュー管理テーブル

##### categories

メニューの分類を管理します。

**主要カラム:**

- `category_id` (varchar(255), PK): カテゴリID
- `name` (varchar(255)): カテゴリ名
- `display_order` (integer): 表示順
- `image_url` (varchar(255)): 画像URL
- `created_at` (timestamptz): 作成時間
- `updated_at` (timestamptz): 更新時間

##### menus

提供される各メニューの詳細情報を保存します。

**主要カラム:**

- `menu_id` (varchar(255), PK): メニューID
- `name` (varchar(255)): メニュー名
- `description` (text): 説明
- `price` (decimal(10,2)): 価格
- `image_url` (varchar(255)): 画像URL
- `is_sold_out` (boolean): 売り切れフラグ
- `category_id` (varchar(255), FK → categories): カテゴリID
- `created_at` (timestamptz): 作成時間
- `updated_at` (timestamptz): 更新時間

##### menu_options

メニューに追加できるオプション（トッピング、サイズ変更など）を管理します。

**主要カラム:**

- `menu_option_id` (varchar(255), PK): オプションID
- `name` (varchar(255)): オプション名
- `price` (decimal(10,2)): 追加価格
- `created_at` (timestamptz): 作成時間
- `updated_at` (timestamptz): 更新時間

##### menu_option_assignments

どのメニューにどのオプションが利用可能かを示す中間テーブルです。

**主要カラム:**

- `menu_id` (varchar(255), PK, FK → menus): メニューID
- `menu_option_id` (varchar(255), PK, FK → menu_options): オプションID

#### テーブル・注文管理テーブル

##### tables

店舗内の物理テーブルを表します。テーブルの状態（空席/使用中/会計待ち）と現在のテーブルセッションIDを管理します。

**主要カラム:**

- `table_id` (varchar(255), PK): テーブルID
- `status` (table_status): 状態（available/occupied/billing）
- `current_table_session_id` (uuid, FK → table_sessions): 現在のテーブルセッションID
- `last_updated` (timestamptz): 最終更新

##### table_sessions

各テーブルのセッション（QR入店〜会計・有効期限まで）を表します。

**主要カラム:**

- `table_session_id` (uuid, PK): テーブルセッションID
- `table_id` (varchar(255), FK → tables): テーブルID
- `is_revoked` (boolean): 無効化フラグ
- `created_at` (timestamptz): 作成時間
- `last_used` (timestamptz): 最終利用時刻
- `expires_at` (timestamptz): 有効期限

##### order_groups

セッション配下の注文グループです。statusでopen/closed/cancelledを管理します。

**主要カラム:**

- `orders_id` (uuid, PK): 注文グループID
- `table_session_id` (uuid, FK → table_sessions): テーブルセッションID
- `status` (order_group_status): グループ状態（open/closed/cancelled）
- `created_at` (timestamptz): 作成時間

**特記事項:**

- 同一セッション内で開いている注文グループは高々1つ（ユニーク部分インデックス: `ux_order_groups_session_open`）

##### order_items

各注文グループに含まれるメニュー項目を管理します。どの注文グループに属し、どのメニューがいくつ注文されたかを記録します。

**主要カラム:**

- `order_item_id` (varchar(255), PK): 注文明細ID
- `orders_id` (uuid, FK → order_groups): 注文グループID
- `menu_id` (varchar(255), FK → menus): メニューID
- `quantity` (integer): 数量
- `price_at_order` (decimal(10,2)): 注文時の価格（後の価格変更に影響されない）
- `status` (order_status): 注文ステータス（pending/preparing/served/cancelled）
- `created_at` (timestamptz): 作成時間

##### order_item_options

注文された各メニュー項目にどのオプションが選択されたかを記録する中間テーブルです。

**主要カラム:**

- `order_item_id` (varchar(255), PK, FK → order_items): 注文明細ID
- `menu_option_id` (varchar(255), PK, FK → menu_options): オプションID

### ENUM型

PostgreSQL の ENUM 型で特定カラムの値を制限し、データの整合性を高めます。

#### table_status

テーブルの状態を管理します。

**値:**

- `available`: 空席
- `occupied`: 使用中
- `billing`: 会計待ち

#### order_status

注文の状態を管理します。

**値:**

- `pending`: 受付待ち
- `preparing`: 調理中
- `served`: 提供済み
- `cancelled`: キャンセル

#### order_group_status

注文グループの状態を管理します。

**値:**

- `open`: オープン（追加注文可能）
- `closed`: クローズ（確定済み）
- `cancelled`: キャンセル

### リレーションシップ

#### 主要な関係性

1. **categories → menus (1:N)**
   - 1つのカテゴリは複数のメニューを持つ

2. **menus → order_items (1:N)**
   - 1つのメニューは複数の注文明細に含まれる

3. **menus ↔ menu_options (N:N)**
   - メニューとオプションは多対多の関係（menu_option_assignmentsで管理）

4. **tables → table_sessions (1:N)**
   - 1つの物理テーブルは複数のセッションを持つ
   - ただし、現在アクティブなセッションは1つのみ（current_table_session_id）

5. **table_sessions → order_groups (1:N)**
   - 1つのセッションは複数の注文グループを持つ

6. **order_groups → order_items (1:N)**
   - 1つの注文グループは複数の注文明細を持つ

7. **order_items ↔ menu_options (N:N)**
   - 注文明細とオプションは多対多の関係（order_item_optionsで管理）

8. **users → sessions (1:N)**
   - 1人のユーザーは複数のログインセッションを持つ

### インデックス

パフォーマンス最適化のために以下のインデックスが定義されています：

- **主キーインデックス**: すべてのテーブル
- **外部キーインデックス**:
  - `idx_table_sessions_table_id`
  - `idx_tables_current_table_session_id`
  - `idx_order_items_orders_id`
  - `idx_order_items_menu_id`
  - `idx_menus_category_id`
- **パフォーマンスインデックス**:
  - `idx_menus_is_sold_out` (売り切れメニューの検索)
  - `idx_order_items_created_at` (日時での集計)
  - `idx_order_groups_created_at` (日時での集計)
- **ユニーク部分インデックス**:
  - `ux_order_groups_session_open` (WHERE status = 'open')

### ビュー（KPI & Analytics）

データ分析とKPI計測のために以下のビューが定義されています：

- **v_kpi_orders_menu**: 注文・収益指標（注文数・総収益・平均注文額・販売アイテム数）
- **v_top_menus_30d**: トップメニュー（過去30日、数量順）
- **v_sales_by_category_30d**: カテゴリ別売上（過去30日）
- **v_menu_performance**: メニュー別パフォーマンス（累計）
- **v_menu_daily_trend_30d**: メニュー日次トレンド（過去30日）
- **v_orders_daily_30d**: 日次トレンド（過去30日：日ごとの注文数・収益）

## 設計上の重要事項

### 循環参照の回避

`tables.current_table_session_id` と `table_sessions.table_id` の関係は循環参照を形成しますが、これは `ALTER TABLE` で後から外部キー制約を追加することで実現しています。

```sql
-- table_sessionsテーブル作成後
ALTER TABLE table_sessions
    ADD CONSTRAINT fk_table_sessions_table FOREIGN KEY (table_id) REFERENCES tables(table_id);

-- tablesテーブルに外部キー制約を追加
ALTER TABLE tables
    ADD CONSTRAINT fk_tables_current_table_session FOREIGN KEY (current_table_session_id) 
    REFERENCES table_sessions(table_session_id);
```

### データ整合性

- `tables.current_table_session_id` と `order_groups.orders_id` は UUID で参照整合性を担保します
- 外部キー制約により、データの整合性が保証されます
- カスケード削除が適切に設定されています（menu_option_assignments、order_item_optionsなど）

### 運用推奨事項

#### セッション終了時の処理

有効期限切れや会計時には、同一トランザクションで以下を実行することを推奨します：

1. `tables.current_table_session_id` を NULL に設定
2. 当該セッション配下の open な `order_groups` を closed に更新

```sql
BEGIN;

-- テーブルの現在のセッションをクリア
UPDATE tables 
SET current_table_session_id = NULL 
WHERE current_table_session_id = <session_id>;

-- セッション配下のオープンな注文グループをクローズ
UPDATE order_groups 
SET status = 'closed' 
WHERE table_session_id = <session_id> AND status = 'open';

COMMIT;
```

#### 注文の追加

同一セッション内で追加注文を受け付ける場合：

1. 現在 `open` 状態の `order_groups` が存在するか確認
2. 存在する場合はそこに `order_items` を追加
3. 存在しない場合は新しい `order_groups` を作成（status = 'open'）

### パフォーマンス最適化

#### インデックス戦略

- **頻繁にJOINされるカラム**: 外部キーには自動的にインデックスが設定されます
- **検索条件に使用されるカラム**: `is_sold_out`, `created_at` などにインデックスを設定
- **集計クエリ**: 日付カラム（`created_at`）にインデックスを設定してKPI計算を高速化
- **部分インデックス**: `WHERE` 句付きのインデックスでストレージとパフォーマンスを最適化

#### ビューの活用

集計クエリは計算コストが高いため、頻繁に参照される統計情報はビューとして定義されています。さらに高速化が必要な場合はマテリアライズドビューの検討も可能です。

## 実装ファイル

実際のデータベース初期化スクリプトは以下にあります：

- `/database/init/init.sql`: PostgreSQL初期化SQL（DDL、インデックス、ビュー定義）
- `/database/postgresql.conf`: PostgreSQL設定ファイル

## 更新手順

データベーススキーマを変更する場合は、以下の順序で更新してください：

1. **`/database/init/init.sql` を更新**: 実際のSQL DDLを変更
2. **`.docs/database/schema.dbml` を更新**: DBMLファイルに変更を反映
3. **`README.md` を更新**: 必要に応じてドキュメントを更新

### 更新時のチェックリスト

- [ ] テーブル定義の変更を `init.sql` に反映
- [ ] 外部キー制約の整合性を確認
- [ ] インデックスが適切に定義されているか確認
- [ ] ENUM型の変更がある場合は既存データの移行を考慮
- [ ] `schema.dbml` に同じ変更を反映
- [ ] マイグレーションスクリプトの作成（必要に応じて）
- [ ] テストデータの更新（`/database/test_data/`）

## 参考リンク

### DBML関連

- [DBML公式ドキュメント](https://www.dbml.org/docs/)
- [dbdiagram.io](https://dbdiagram.io/) - オンラインER図ツール
- [DBML CLI](https://www.dbml.org/cli/) - コマンドラインツール
- [dbdocs.io](https://dbdocs.io/) - オンラインドキュメント生成

### PostgreSQL関連

- [PostgreSQL 17.5 Documentation](https://www.postgresql.org/docs/17/)
- [PostgreSQL ENUM Types](https://www.postgresql.org/docs/current/datatype-enum.html)
- [PostgreSQL Indexes](https://www.postgresql.org/docs/current/indexes.html)

## トラブルシューティング

### よくある問題

#### 循環参照エラー

**症状**: テーブル作成時に外部キー制約のエラーが発生
**解決**: `init.sql` では、循環参照を避けるため外部キー制約を `ALTER TABLE` で後から追加しています

#### ENUM型の値変更

**症状**: ENUM型の値を変更したいが直接変更できない
**解決**: PostgreSQLのENUM型は値の追加は可能ですが、削除や変更には特別な手順が必要です

```sql
-- 新しい値を追加
ALTER TYPE order_status ADD VALUE 'new_status';

-- 値の削除・変更は型の再作成が必要
-- 1. 新しい型を作成
-- 2. カラムの型を変更
-- 3. 古い型を削除
```

#### インデックスのパフォーマンス

**症状**: クエリが遅い
**解決**:

1. `EXPLAIN ANALYZE` でクエリプランを確認
2. 頻繁に検索されるカラムにインデックスを追加
3. 複合インデックスの検討
4. 不要なインデックスの削除（INSERT/UPDATEが遅くなるため）
