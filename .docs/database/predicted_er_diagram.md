# テーブル設計 (ER図)

機能一覧から予測されるデータベースのテーブル設計を以下に示します。

```mermaid
erDiagram
    settings {
        string key PK "設定キー (例: logo_url, table_count)"
        string value "設定値"
    }

    tables {
        string id PK "テーブルID"
        string name "テーブル名 (例: テーブル1)"
        string status "状態 (空席, 使用中, 会計済み)"
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
        string id PK "注文ID"
        string table_id FK "テーブルID"
        string status "注文状態 (調理中, 提供済み, キャンセル)"
        timestamp created_at "注文日時"
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

    tables      ||--o{ orders : "has"
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
- **tables**: 店舗内の物理的なテーブルとその状態を管理します。
- **categories**: メニューの分類を管理します。
- **menus**: 提供される各メニューの詳細情報を保存します。
- **menu_options**: メニューに追加できるオプション（トッピング、サイズ変更など）を管理します。
- **menu_option_assignments**: どのメニューにどのオプションが利用可能かを示す中間テーブルです。
- **orders**: 顧客からの注文単位を管理します。どのテーブルからの注文かが記録されます。
- **order_items**: 各注文に含まれる個別のメニュー項目を管理します。
- **order_item_options**: 注文された各メニュー項目にどのオプションが選択されたかを記録する中間テーブルです。
