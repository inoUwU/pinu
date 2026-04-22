# PINU 現行システム詳細設計・仕様書（リバースエンジニアリング）

## 1. 文書の目的

本書は、現行の実装、既存ドキュメント、データベース定義をもとに、PINU の現在仕様を再構成した詳細設計書です。  
新規実装時の前提共有、既存機能の把握、今後の仕様整理の土台として利用します。

本書では、**実コードで確認できる事実を優先**し、未実装部分や整合していない箇所はその旨を明記します。

## 2. システム概要

PINU は、小規模飲食店向けの Web ベース注文支援システムです。  
主な利用者は以下の 3 役割です。

- **客席利用者**: QR からメニューを見て注文し、会計依頼まで行う
- **店主 / 管理者**: メニュー、ユーザー、設定を管理する
- **厨房 / レジ担当**: 注文受信、会計対象テーブル確認、注文内容表示を行う

### 2.1 システムの特徴

- テーブル単位で注文状態を管理する
- 注文は `table_sessions` と `order_groups` を軸に蓄積する
- 会計時は金額を**表示**するが、POS のような**売上計上・入金記録は行わない**
- 分析用に KPI / 売れ筋 / 日次推移ビューを持つ
- バックエンドは Go + Fiber + Bun + PostgreSQL
- フロントエンドは Next.js 単一アプリで、`/` 系が客席、`/admin` 系が管理画面

## 3. 現行アーキテクチャ

### 3.1 バックエンド構成

- ルーティング: `backend/app/api/routes.go`
- ハンドラー: `backend/app/handlers/*`
- ユースケース: `backend/app/usecases/*`
- リポジトリ: `backend/app/infrastructure/repositories/*`
- DI: `backend/app/bootstrap/injector.go`

責務の流れは以下です。

`routes -> handlers -> usecases -> repositories -> PostgreSQL`

### 3.2 フロントエンド構成

- 客席画面: `frontend/app/(client)/*`
- 管理画面: `frontend/app/admin/*`
- API 呼び出し: `frontend/lib/api/*`
- SWR フック: `frontend/hooks/*`

### 3.3 データベース構成

主定義:

- `database/init/init.sql`
- `docs/database/schema.dbml`

注文・会計・分析の中心は以下のテーブルです。

- `tables`
- `table_sessions`
- `order_groups`
- `order_items`
- `menus`
- `categories`
- `menu_options`

## 4. ドメインモデル

### 4.1 マスタ系

#### settings

アプリ全体設定を保持します。現状は Key-Value 方式です。

想定用途:

- ロゴ URL
- テーブル数
- 今後の店舗設定

#### users

管理画面利用者のアカウントです。

- `login_id` でログイン
- `password_hash` / `password_salt` を保持
- `is_admin` で管理者権限を識別

#### sessions

管理画面ログイン用セッションです。

- リフレッシュトークン単位で管理
- `is_revoked` による失効管理あり

### 4.2 メニュー系

#### categories

メニュー分類です。

- 表示順 `display_order`
- カテゴリ画像 `image_url`

#### menus

提供メニュー本体です。

- 名称
- 説明
- 価格
- 画像 URL
- 売り切れ状態 `is_sold_out`
- 所属カテゴリ `category_id`

#### menu_options / menu_option_assignments

トッピングやサイズ変更などの追加オプションです。  
`menu_option_assignments` により、メニューとオプションは多対多で紐づきます。

### 4.3 注文・テーブル系

#### tables

店舗の物理テーブルです。

状態:

- `available`: 空席
- `occupied`: 利用中
- `billing`: 会計待ち / 会計対応中

`current_table_session_id` により、現在有効なテーブルセッションを参照します。

#### table_sessions

テーブル利用単位のセッションです。  
QR で着席してから会計で無効化される単位として設計されています。

- `is_revoked`: 強制終了済みか
- `last_used`: 最終注文時刻
- `expires_at`: セッション期限

#### order_groups

同一セッション配下の注文まとまりです。  
追加注文に対応するため、テーブルセッションと 1 対多で保持されます。

状態:

- `open`: 追加注文受付中
- `closed`: 会計または締め処理済み
- `cancelled`: 無効

**同一セッションで open は 1 件まで**です。

#### order_items / order_item_options

注文された個々の明細です。

- `price_at_order` に注文時価格を保持
- メニュー価格変更後も履歴側の金額を固定できる
- オプション選択は `order_item_options` で保持

状態:

- `pending`
- `preparing`
- `served`
- `cancelled`

## 5. 業務フロー

### 5.1 店舗初期設定

管理者が以下を準備する想定です。

- ユーザー登録
- テーブル登録
- カテゴリ登録
- メニュー登録
- オプション登録
- 必要に応じた設定値登録

### 5.2 客席開始

既存 docs と DB 設計上は、**QR 読み取りでテーブルセッションを発行し、テーブルを occupied にする**運用が想定されています。

ただし現状コードでは:

- `table_sessions` の永続化機能は存在
- `qr_handler.go` は未実装
- テーブルセッション発行 API は未公開

そのため、**セッション開始フローは DB 設計上は存在するが、公開 API / 画面としては未完成**です。

### 5.3 注文

注文作成 API:

- `POST /api/orders`

入力:

- `table_session_id`
- `items[]`
  - `menu_id`
  - `quantity`
  - `menu_option_ids[]`

処理内容:

1. `table_session_id` を検証
2. セッション存在 / revoke / expiry を確認
3. セッション配下の既存 `open` 注文グループを取得
4. なければ新しい `order_group` を作成
5. 各注文についてメニュー存在と売り切れ状態を確認
6. `order_items` を作成
7. 必要に応じて `order_item_options` を作成
8. `table_sessions.last_used` を更新
9. テーブル状態を `occupied` に更新

注文は**都度 1 件ずつ計上するのではなく、open な注文グループへ追加**されます。

### 5.4 注文参照

注文一覧取得 API:

- `GET /api/orders?table_session_id=<uuid>`

返却内容:

- 注文グループ一覧
- 各グループ配下の注文明細
- 明細ごとのオプション ID

### 5.5 厨房通知

SSE エンドポイント:

- `GET /api/sse`
- `PUT /api/publish`

現状はシンプルなメッセージキュー型で、以下を配信します。

- 通常メッセージ
- ハートビート

ただし厨房画面はまだダミーデータ中心で、**注文データとの自動連携は未完成**です。

### 5.6 会計

会計画面は管理画面 `/admin/(authenticated)/bill` で実装が進んでいます。

現行仕様:

1. テーブル一覧取得
2. `occupied` / `billing` テーブルを会計対象として表示
3. `current_table_session_id` を使って注文一覧を取得
4. `price_at_order * quantity` の合計を表示
5. 会計実行で `POST /api/tables/:id/checkout`

会計時のサーバー処理:

1. テーブルが存在することを確認
2. テーブル状態が `occupied` であることを確認
3. 有効な `current_table_session_id` があることを確認
4. テーブル状態を `billing` に更新
5. セッション配下の `open` 注文グループを `closed` に更新
6. 該当テーブルの未 revoke な `table_sessions` を revoke
7. `tables.current_table_session_id` を NULL に更新

### 5.7 会計の扱い

現行コードと DB から判断すると、PINU における会計は**業務上の締め処理**です。  
システムが行うのは以下までです。

- 注文履歴の集約
- 合計金額の表示
- テーブル状態の会計済み相当への移行
- セッションの終了

以下は行いません。

- 入金手段管理
- 現金収受記録
- 売上伝票発行
- 会計仕訳
- POS 会計計上

つまり、**金額表示はするが、システム内で会計計上はしない**設計です。

### 5.8 分析

分析 API:

- `GET /api/analytics`
- `GET /api/analytics/kpi`
- `GET /api/analytics/top-menus`
- `GET /api/analytics/category-sales`
- `GET /api/analytics/menu-performance`
- `GET /api/analytics/daily-sales`
- `GET /api/analytics/menu-daily-trends`

集計元は DB ビューです。  
ダッシュボード画面では以下を表示します。

- KPI サマリー
- トップメニュー
- カテゴリ別売上
- 日次売上推移

## 6. 状態遷移

### 6.1 テーブル状態

`backend/app/domain/table/table.go` で遷移ルールを保持しています。

| 現在 | 遷移可能 |
| --- | --- |
| available | occupied |
| occupied | billing, available |
| billing | available |

### 6.2 注文明細状態

`backend/app/domain/order/order.go` で遷移ルールを保持しています。

| 現在 | 遷移可能 |
| --- | --- |
| pending | preparing, cancelled |
| preparing | served, cancelled |
| served | 遷移不可 |
| cancelled | 遷移不可 |

### 6.3 注文グループ状態

注文グループは現状、主に以下の用途です。

- `open`: 追加注文受付中
- `closed`: 会計締め済み
- `cancelled`: 無効化済み

## 7. 画面仕様（現行実装ベース）

### 7.1 客席画面

| 画面 | パス | 状態 |
| --- | --- | --- |
| ルートページ | `/` | モック表示 |
| カテゴリ一覧 | `/category` | ダミーデータ表示 |
| カテゴリ別メニュー | `/[categoryid]` | プレースホルダ |
| メニュー詳細 | `/[categoryid]/[mid]` | プレースホルダ |
| 注文画面 | `/order` | プレースホルダ |
| 履歴・会計 | `/history` | サンプルテーブル表示 |
| チェックアウト完了 | `/checkout` | 実装あり |

### 7.2 管理画面

| 画面 | パス | 状態 |
| --- | --- | --- |
| ログイン | `/admin/login` | フォームあり |
| 新規登録 | `/admin/register` | UI のみ |
| ダッシュボード | `/admin` | 分析 API 連携あり |
| メニュー管理 | `/admin/menu` | ローカル状態中心、API 未接続中心 |
| オペレーション | `/admin/operation` | SSE 接続あり、注文一覧はダミーデータ |
| 会計 | `/admin/bill` | テーブル / 注文 API 連携あり |
| ユーザー管理 | `/admin/user` | ローカル状態中心 |
| 設定 | `/admin/settings` | 最小表示のみ |
| QR 表示 | `/admin/qrcode` | 端末接続用 QR 表示 |

## 8. API 一覧

### 8.1 認証

- `POST /api/auth/login`
- `POST /api/auth/logout`
- `POST /api/auth/renew`
- `POST /api/auth/revoke`

### 8.2 ユーザー

- `GET /api/user/`
- `GET /api/user/:id`
- `POST /api/user/register`
- `POST /api/user/update`
- `DELETE /api/user/delete`

### 8.3 カテゴリ

- `POST /api/categories/`
- `GET /api/categories/`
- `GET /api/categories/:id`
- `PUT /api/categories/:id`
- `DELETE /api/categories/:id`

### 8.4 メニュー

- `POST /api/menus/`
- `GET /api/menus/`
- `GET /api/menus/:id`
- `GET /api/menus/category/:categoryId`
- `PUT /api/menus/:id`
- `DELETE /api/menus/:id`

### 8.5 メニューオプション

- `POST /api/menu-options/`
- `GET /api/menu-options/`
- `GET /api/menu-options/:id`
- `PUT /api/menu-options/:id`
- `DELETE /api/menu-options/:id`

### 8.6 テーブル

- `POST /api/tables/`
- `GET /api/tables/`
- `GET /api/tables/:id`
- `GET /api/tables/status/:status`
- `PUT /api/tables/:id/status`
- `DELETE /api/tables/:id`
- `POST /api/tables/:id/checkout`

### 8.7 設定

- `GET /api/settings/`
- `GET /api/settings/:key`
- `POST /api/settings/`
- `DELETE /api/settings/:key`

### 8.8 注文

- `GET /api/orders?table_session_id=...`
- `POST /api/orders`

### 8.9 分析

- `GET /api/analytics/`
- `GET /api/analytics/kpi`
- `GET /api/analytics/top-menus`
- `GET /api/analytics/category-sales`
- `GET /api/analytics/menu-performance`
- `GET /api/analytics/daily-sales`
- `GET /api/analytics/menu-daily-trends`

### 8.10 SSE

- `GET /api/sse`
- `PUT /api/publish`

## 9. 非機能・運用上の前提

- 小規模店舗での単純運用を前提とする
- 管理画面は認証前提
- 注文金額は履歴参照と表示用に保持する
- 金額は分析にも使うが、会計システムとしての確定処理は持たない
- セッション失効や会計時の整合性はトランザクションで保つ設計になっている

## 10. 現行実装から見えるギャップ

### 10.1 実装済み度が高い領域

- バックエンド CRUD API の骨格
- 注文作成 / 注文参照
- テーブル会計処理
- 分析ビューとダッシュボード表示
- ログイン処理

### 10.2 未完成または暫定実装の領域

- QR 読み取りからのテーブルセッション開始
- 客席画面の API 接続
- 厨房画面の実注文連携
- 管理画面のユーザー管理 API 連携
- 設定画面
- メニュー管理の永続化連携
- 注文ステータス更新 UI

### 10.3 注意すべき不整合

- 既存 docs では「会計済み」と表現される箇所があるが、現行コード上のテーブル状態は `billing`
- フロントの一部認証 API 呼び出しは、バックエンドの期待フィールド名と差がある
- 分析では revenue を扱うが、会計計上システムではなく表示・分析用途に留まる

## 11. 現時点の整理された仕様要約

PINU は、**テーブルセッションを軸に注文を蓄積し、厨房通知・会計表示・分析へつなぐ小規模飲食店向け注文支援システム**です。  
会計は「注文を締める業務操作」であり、**金額は表示するが、システム内で売上計上や入金処理までは行わない**ことが現行設計の前提です。

現状は、**バックエンドの注文 / 会計 / 分析基盤が先行し、客席 UI と一部管理 UI が追従中**の状態です。

## 12. 参照元

主に以下を参照して本書を作成しました。

- `/home/runner/work/pinu/pinu/database/init/init.sql`
- `/home/runner/work/pinu/pinu/docs/database/schema.dbml`
- `/home/runner/work/pinu/pinu/backend/app/api/routes.go`
- `/home/runner/work/pinu/pinu/backend/app/usecases/order/order_usecase.go`
- `/home/runner/work/pinu/pinu/backend/app/usecases/table/table_usecase.go`
- `/home/runner/work/pinu/pinu/backend/app/usecases/analytics/analytics_usecase.go`
- `/home/runner/work/pinu/pinu/frontend/app/admin/(authenticated)/bill/page.tsx`
- `/home/runner/work/pinu/pinu/frontend/app/admin/(authenticated)/page.tsx`
- `/home/runner/work/pinu/pinu/frontend/app/(client)/*`
