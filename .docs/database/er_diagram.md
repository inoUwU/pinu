# テーブル設計 (ER図)

```mermaid
erDiagram
    settings {
        string key PK "設定キー (例: logo_url, table_count)"
        string value "設定値"
    }
    
    users {
        string user_id PK "ユーザーID (UUIDv7)"
        string username "ログインID"
        string password_hash "パスワードハッシュ"
        string name "表示名"
        timestamp created_at "アカウント作成時間"
        boolean is_admin "管理者フラグ"
    }

    tables ||--o{ table_sessions : "has"
    tables {
        string table_id PK "テーブルID（例：T003）"
        table_status status "テーブルの状態 (ENUM)"
        string current_orders_id FK "現在の注文ID (UUIDv7)"
        timestamp last_updated "最終更新時間"
    }

    table_sessions {
        uuid table_session_id PK "テーブルセッションID(uuidv7)"
        string table_id FK "関連するテーブルID"
        string orders_id FK "注文ID (UUIDv7)"
        timestamp created_at "作成時間"
        timestamp last_used "最終使用時間"
        timestamp expires_at "有効期限（フェイルセーフ用/2時間）"
    }
    
    orders {
        string order_item_id PK "注文明細ID"
        string orders_id FK "注文ID (UUIDv7)"
        string menu_id FK "メニューID"
        integer quantity "数量"
        decimal price_at_order "注文時の価格"
        order_status status "注文状態 (ENUM)"
        timestamp created_at "注文時間"
    }

    order_item_options {
        string order_item_id FK "注文明細ID"
        string menu_option_id FK "オプションID"
    }

    categories {
        string category_id PK "カテゴリID"
        string name "カテゴリ名 (例: ドリンク, 前菜)"
        integer display_order "表示順"
    }

    menus {
        string menu_id PK "メニューID"
        string name "メニュー名"
        string description "説明"
        decimal price "価格 (DECIMAL)"
        string image_url "画像URL"
        boolean is_sold_out "品切れフラグ"
        string category_id FK "カテゴリID"
    }

    menu_options {
        string menu_option_id PK "オプションID"
        string name "オプション名 (例: 大盛り, 辛さ増し)"
        decimal price "追加料金 (DECIMAL)"
    }

    menu_option_assignments {
        string menu_id FK "メニューID"
        string menu_option_id FK "オプションID"
    }

    categories  ||--o{ menus : "contains"
    menus       }|--|| menu_option_assignments : "can have"
    menu_options }|--|| menu_option_assignments : "can be"
    table_sessions }|--o{ orders : "has"
    menus       ||--o{ orders : "is ordered in"
    orders      }o--|| order_item_options : "has"
    menu_options }|--o{ order_item_options : "is selected for"
```

## 各テーブルの説明

- **settings**: アプリケーション全体の設定を保存します（例：ロゴのURL、店舗のテーブル総数）。
- **tables**: 店舗内の物理テーブルを表します。テーブルの状態（空席/使用中/会計済み）と現在の注文グループIDを管理します。
- **order_tokens**: 顧客がQRコードを読み取った際に生成されるアクセストークンを管理します。group_idはUUIDv7で発行され、2時間未使用または会計時に無効化されます。
- **categories**: メニューの分類を管理します。
- **menus**: 提供される各メニューの詳細情報を保存します。
- **menu_options**: メニューに追加できるオプション（トッピング、サイズ変更など）を管理します。
- **menu_option_assignments**: どのメニューにどのオプションが利用可能かを示す中間テーブルです。
- **orders**: 各注文に含まれる個別のメニュー項目を管理します。どの注文グループに属し、どのメニューがいくつ注文されたかを記録します。
- **order_item_options**: 注文された各メニュー項目にどのオプションが選択されたかを記録する中間テーブルです。

## データ型 (ENUM)

PostgreSQLのENUM型を利用して、特定カラムの値を制限し、データ整合性を高めます。

- **table_status**: テーブルの状態を管理します。
  - `('available', 'occupied', 'billing')`  `('空席', '使用中', '会計済み')`
- **order_status**: 注文の状態を管理します。
  - `('pending', 'preparing', 'served', 'cancelled')` `('受付待ち', '調理中', '提供済み', 'キャンセル')`

---

- tables.current_orders_id, order_tokens.group_id, orders.orders_idはすべてUUIDv7で統一し、参照整合性を担保します。
- order_tokens.is_activeで論理削除・無効化を明示できます。
- 有効期限切れや会計時の自動クリーンアップ運用を推奨します。
