# QRオーダーシステム ER図

以下のER図は、QRオーダーシステムのデータベース設計を表しています。テーブル間の関係と各テーブルのフィールドを示しています。

```mermaid
erDiagram
    tables ||--o{ order_tokens : "has"
    tables {
        string table_id PK "テーブルの識別子（例：T003）"
        string status "テーブルの状態（空席/使用中/会計済み）"
        string current_group_id "現在の注文グループID"
        timestamp last_updated "最終更新時間"
    }
    
    order_tokens ||--o{ orders : "belongs to"
    order_tokens {
        uuid token PK "UUIDトークン"
        string table_id FK "関連するテーブルID"
        string group_id "注文グループID"
        timestamp created_at "作成時間"
        timestamp last_used "最終使用時間"
        timestamp expires_at "有効期限（フェイルセーフ用）"
    }
    
    orders {
        string order_id PK "注文ID"
        string group_id FK "注文グループID"
        string menu_item_id "商品ID"
        integer quantity "数量"
        decimal price "価格"
        string status "注文状態"
        timestamp created_at "注文時間"
    }
```

## テーブル説明

### tables

店舗内の物理テーブルを表します。テーブルの状態（空席/使用中/会計済み）を管理します。

### order_tokens

顧客がQRコードを読み取った際に生成されるアクセストークンを管理します。同じテーブルの複数の顧客は同じgroup_idに関連付けられます。

### orders

顧客からの注文情報を保存します。各注文は特定の注文グループに属します。
