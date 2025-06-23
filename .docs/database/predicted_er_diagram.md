# テーブル設計 (ER図)

```mermaid
erDiagram
    settings {
        string key PK "設定キー (例: logo_url, table_count)"
        string value "設定値"
    }

    tables ||--o{ order_tokens : "has"
    tables {
        string table_id PK "テーブルの識別子（例：T003）"
        string status "テーブルの状態（空席/使用中/会計済み）"
        string current_group_id FK "現在の注文グループID (UUIDv7)"
        timestamp last_updated "最終更新時間"
    }

    order_tokens ||--o{ orders : "belongs to"
    order_tokens {
        uuid token PK "UUIDトークン"
        string table_id FK "関連するテーブルID"
        string group_id FK "注文グループID (UUIDv7)"
        timestamp created_at "作成時間"
        timestamp last_used "最終使用時間"
        timestamp expires_at "有効期限（フェイルセーフ用/2時間）"
        boolean is_active "有効フラグ（会計時や期限切れでfalse）"
    }

    categories {
        string id PK "カテゴリID"
        string name "カテゴリ名 (例: ドリンク, 前菜)"
        integer display_order "表示順"
    }

    menus {
        string id PK "メニューID"
        string name "メニュー名"
        string description "説明"
        integer price "価格"
        string image_url "画像URL"
        boolean is_sold_out "品切れフラグ"
        string category_id FK "カテゴリID"
    }

    menu_options {
        string id PK "オプションID"
        string name "オプション名 (例: 大盛り, 辛さ増し)"
        integer price "追加料金"
    }

    menu_option_assignments {
        string menu_id FK "メニューID"
        string option_id FK "オプションID"
    }

    orders {
        string order_id PK "注文ID"
        string group_id FK "注文グループID (UUIDv7)"
        string menu_item_id "商品ID"
        integer quantity "数量"
        decimal price "価格"
        string status "注文状態"
        timestamp created_at "注文時間"
    }

    order_items {
        string id PK "注文明細ID"
        string order_id FK "注文ID"
        string menu_id FK "メニューID"
        integer quantity "数量"
    }

    order_item_options {
        string order_item_id FK "注文明細ID"
        string option_id FK "オプションID"
    }

    categories  ||--o{ menus : "contains"
    menus       }|--|| menu_option_assignments : "can have"
    menu_options }|--|| menu_option_assignments : "can be"
    orders      ||--o{ order_items : "contains"
    menus       ||--o{ order_items : "is ordered in"
    order_items }o--|| order_item_options : "has"
    menu_options }|--o{ order_item_options : "is selected for"
```

## 各テーブルの説明

- **settings**: アプリケーション全体の設定を保存します（例：ロゴのURL、店舗のテーブル総数）。
- **tables**: 店舗内の物理テーブルを表します。テーブルの状態（空席/使用中/会計済み）と現在の注文グループIDを管理します。
- **order_tokens**: 顧客がQRコードを読み取った際に生成されるアクセストークンを管理します。group_idはUUIDv7で発行され、2時間未使用または会計時に無効化されます。
- **orders**: 顧客からの注文情報を保存します。各注文は特定の注文グループ（group_id）に属します。
- **categories**: メニューの分類を管理します。
- **menus**: 提供される各メニューの詳細情報を保存します。
- **menu_options**: メニューに追加できるオプション（トッピング、サイズ変更など）を管理します。
- **menu_option_assignments**: どのメニューにどのオプションが利用可能かを示す中間テーブルです。
- **order_items**: 各注文に含まれる個別のメニュー項目を管理します。
- **order_item_options**: 注文された各メニュー項目にどのオプションが選択されたかを記録する中間テーブルです。

---

- tables.current_group_id, order_tokens.group_id, orders.group_idはすべてUUIDv7で統一し、参照整合性を担保します。
- order_tokens.is_activeで論理削除・無効化を明示できます。
- 有効期限切れや会計時の自動クリーンアップ運用を推奨します。
