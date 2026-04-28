# セッション管理フロー図

```mermaid
flowchart TD
    A[来店客がテーブルの QR を読み取る] --> B[フロントが qr_token を送信]
    B --> C[サーバーが qr_token から table_id を解決]
    C --> D{テーブル状態と active session を判定}
    D -- available --> E[新規 table_session_id を発行]
    E --> F[current_table_session_id を設定]
    F --> G[テーブル状態を occupied に更新]
    G --> H[table_session_id を返却]
    D -- occupied かつ active session あり --> I[既存 table_session_id を返却]
    I --> H
    D -- billing --> J[利用不可を返却]
    H --> K[フロントが table_session_id を保持して客用画面へ遷移]
    K --> L[注文・履歴取得で table_session_id を送信]
    L --> M[サーバーが session を検証]
    M -- 有効 --> N[注文処理 / 履歴返却]
    M -- 無効 --> J
    N --> O[来店客が会計依頼を送信]
    O --> P[スタッフに会計依頼を通知]
    P --> Q[スタッフが会計確定]
    Q --> R[order_groups を closed に更新]
    R --> S[table session を revoke]
    S --> T[current_table_session_id をクリア]
    T --> U[テーブル状態を available に更新]
```

- QR コードは固定のテーブル識別子であり、読み取り成功時にサーバーが `table_session_id` を新規発行または再利用します。
- テーブル状態は `available` → `occupied` → `billing` → `available` を基本サイクルとします。
- 専用の heartbeat API は設けず、`last_used` はテーブルセッション参照、注文、会計依頼などの成功時に更新します。
- 会計依頼はスタッフへの通知イベントであり、会計確定時に `table_session_id` を revoke します。
