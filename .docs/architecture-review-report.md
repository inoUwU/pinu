# ヘキサゴナルアーキテクチャ調査レポート

## 実施日
2025-12-05

## 概要
バックエンドコードのヘキサゴナルアーキテクチャ（Port & Adapter）パターンの実装を調査し、アーキテクチャ違反やアンチパターンを特定、修正しました。

## 調査範囲
- `backend/app/domain/` - ドメイン層（エンティティとリポジトリインターフェース）
- `backend/app/usecases/` - アプリケーション層（ユースケース実装）
- `backend/app/infrastructure/` - インフラストラクチャ層（リポジトリ実装、モデル）
- `backend/app/handlers/` - プライマリアダプター（HTTPハンドラー）

## 発見した問題と修正

### 🔴 重大な違反 (Critical Violations)

#### 問題1: UsecaseがInfrastructure層の具体型に直接依存 ✅ 修正完了

**違反内容:**
- `UserUsecaseImpl`と`AuthUsecaseImpl`が`*repositories.TxRepository`という具体的な実装に依存
- ヘキサゴナルアーキテクチャの原則「内側の層は外側の層に依存しない」に違反

**影響:**
- レイヤー間の結合度が高い
- テスタビリティの低下
- Infrastructure層の変更がUsecase層に影響

**修正内容:**
- `app/domain/port/transaction.go`に`TransactionManager`インターフェースを定義
- `TxRepository`を`port.TransactionManager`の実装として再定義
- すべてのUsecaseでインターフェースに依存するように変更

**修正ファイル:**
- `backend/app/domain/port/transaction.go` (新規作成)
- `backend/app/infrastructure/repositories/tx.go`
- `backend/app/usecases/user/user_usecase.go`
- `backend/app/usecases/auth/auth_usercase.go`
- `backend/app/usecases/analytics/analytics_usecase.go`
- `backend/app/middleware/injection.go`

#### 問題4: Order Repository実装の不完全性とビルドエラー ✅ 修正完了

**違反内容:**
- 未定義の型`OrderGroup`、`OrderItem`を参照
- ビルドが失敗する状態

**修正内容:**
- 正しい`models.OrderGroupModel`を使用
- 必要なimportを追加

**修正ファイル:**
- `backend/app/infrastructure/repositories/order/order_repository.go`

### 🟡 中程度の問題 (Moderate Issues)

#### 問題5: Repository インターフェースの命名の不一致 ✅ 修正完了

**違反内容:**
- 一部のRepositoryインターフェースに`I`プレフィックスあり（`IUserRepository`, `ISessionRepository`）
- 他のRepositoryにはプレフィックスなし（`MenuRepository`, `CategoryRepository`）
- Goの慣習では、インターフェース名に`I`プレフィックスを付けない

**修正内容:**
- すべてのRepositoryインターフェースから`I`プレフィックスを削除
  - `IUserRepository` → `UserRepository`
  - `ISessionRepository` → `SessionRepository`
  - `IOrderRepository` → `OrderRepository`
  - `IAnalyticsRepository` → `AnalyticsRepository`

**修正ファイル:**
- `backend/app/domain/user/user_repository.go`
- `backend/app/domain/session/session_repository.go`
- `backend/app/domain/order/order_repository.go`
- `backend/app/domain/analytics/analytics_repository.go`
- 上記インターフェースを参照するすべてのファイル

#### 問題7: GetAllUsersでの不要なトランザクション使用 ✅ 修正完了

**違反内容:**
- 読み取り専用操作にトランザクションを使用
- パフォーマンスの低下とリソースの無駄遣い

**修正内容:**
- `GetAllUsers`メソッドからトランザクション処理を削除
- 直接Repositoryを呼び出すように変更

**修正ファイル:**
- `backend/app/usecases/user/user_usecase.go`

### 🟢 軽微な問題 (Minor Issues)

#### 問題8: Domain EntityでのDB tagの使用 ✅ 修正完了

**違反内容:**
- Domainエンティティに`bun`（ORM）のタグが付与されている
- Domain層がInfrastructure層の技術詳細に依存

**修正内容:**
- すべてのDomainエンティティから`bun`タグを削除
- JSONタグのみを残す（ドメインの表現として適切）

**修正ファイル:**
- `backend/app/domain/order/order.go`
- `backend/app/domain/category/category.go`
- `backend/app/domain/settings/settings.go`
- `backend/app/domain/table/table.go`
- `backend/app/domain/menu_option/menu_option.go`

## 未修正の問題

### 問題2: Infrastructure層のモデルとDomain層のエンティティが混在

**現状:**
- `app/infrastructure/models/user.go`と`app/domain/user/user.go`が共存
- Repository実装で変換処理が必要

**推奨される対応:**
1. 最小限の変更: 現状維持（両方の型を使い分ける）
2. 理想的: Infrastructure modelsを削除し、Domainエンティティに直接bunタグを付与（ただし、問題8の修正と矛盾）
3. 折衷案: Infrastructure modelsを明確にORMマッピング専用として位置づけ、変換ロジックを整理

**理由:**
- 大規模な変更が必要
- 既存コードへの影響が大きい
- プロジェクトの方針決定が必要

### 問題3: Repository実装でのトランザクション処理のアンチパターン

**現状:**
- Repository内でContext経由でトランザクションを取得
- `ctx.Value(ctxkey.TxCtxKey).(bun.Tx)`のパターンが複数箇所

**推奨される対応:**
- 現在のトランザクション管理は機能しているため、即座の変更は不要
- 将来的にはRepository層の責務を再検討

### 問題6: Usecase層でのEntityの不適切な使用

**現状:**
- Output DTOがDomainエンティティをそのまま含む
- レイヤー間の境界が曖昧

**推奨される対応:**
- Output DTOに独自の型を定義
- エンティティからDTOへの変換を明示的に実施

### 問題9: Usecase内での時間生成

**現状:**
- `time.Now()`を直接呼び出し
- テスト時に時間を制御できない

**推奨される対応:**
- Clock interfaceを定義
- 本番環境とテスト環境で実装を切り替え可能にする

### 問題10: エラーハンドリングの不一致

**現状:**
- カスタムエラーと標準エラーが混在
- エラーハンドリングのパターンが統一されていない

**推奨される対応:**
- Domain層でカスタムエラー型を定義
- すべてのレイヤーで一貫したエラーハンドリングを実施

## アーキテクチャの推奨構造

### 依存関係の方向
```
┌─────────────────────────────────────┐
│  Handlers (Primary Adapter)         │ ← HTTP入力
│  app/handlers/                      │
└────────────┬────────────────────────┘
             │ depends on
             ▼
┌─────────────────────────────────────┐
│  UseCases (Application Layer)       │
│  app/usecases/                      │
└────────────┬────────────────────────┘
             │ depends on (interface)
             ▼
┌─────────────────────────────────────┐
│  Domain (Entities + Ports)          │ ← ビジネスロジックの中心
│  app/domain/                        │
└────────────▲────────────────────────┘
             │ implements
             │
┌────────────┴────────────────────────┐
│  Infrastructure (Adapters)          │
│  app/infrastructure/repositories/   │
└────────────┬────────────────────────┘
             │ depends on
             ▼
┌─────────────────────────────────────┐
│  Database (PostgreSQL)              │
└─────────────────────────────────────┘
```

### レイヤーの責務

**Domain Layer** (`app/domain/`)
- エンティティの定義
- リポジトリインターフェース（Port）の定義
- ビジネスルールの定義
- 他のレイヤーに依存しない

**Application Layer** (`app/usecases/`)
- ユースケースの実装
- ビジネスロジックのオーケストレーション
- Domain層のPortに依存
- Infrastructure層の具体実装に依存しない

**Infrastructure Layer** (`app/infrastructure/`)
- リポジトリの具体実装（Adapter）
- データベースアクセス
- 外部サービスとの連携
- Domain層のPortを実装

**Handler Layer** (`app/handlers/`)
- HTTPリクエスト/レスポンスの処理
- Usecaseの呼び出し
- プレゼンテーション層

## ベストプラクティス

### DO（推奨）
1. ✅ Interface（Port）をDomain層に定義
2. ✅ UsecaseはInterfaceに依存
3. ✅ Infrastructure実装はInterfaceを実装
4. ✅ Domainエンティティは技術詳細に依存しない
5. ✅ Goの慣習に従った命名（Interfaceに`I`プレフィックスなし）
6. ✅ 読み取り専用操作はトランザクション不要

### DON'T（非推奨）
1. ❌ UsecaseからInfrastructure実装に直接依存
2. ❌ Domainエンティティにインフラ固有のタグ
3. ❌ Interface命名の不統一
4. ❌ 不要なトランザクションの使用
5. ❌ レイヤー境界の曖昧さ

## 検証結果

### ビルド状態
- ✅ すべての変更後にビルド成功
- ✅ 型の整合性を確認
- ✅ 依存関係の循環なし

### テスト
- ⚠️ 既存のテストコードが存在しないため、テスト実行は未実施
- 推奨: ユニットテストとインテグレーションテストの追加

## 今後の推奨事項

### 短期（1-2週間）
1. 問題6: Output DTOの適切な定義と実装
2. テストコードの追加（特にUsecase層）
3. 既存の変換ロジックの見直し

### 中期（1-2ヶ月）
1. 問題9: Clock interfaceの導入
2. 問題10: エラーハンドリングの統一
3. ドキュメントの充実化

### 長期（3ヶ月以上）
1. 問題2: モデル層の整理と統一
2. 問題3: トランザクション管理の再設計
3. アーキテクチャテストの導入（archunit等）

## まとめ

今回の調査と修正により、以下を達成しました：

✅ **完了した修正（7項目）**
1. トランザクション管理のPort定義
2. Repository命名規則の統一
3. 不要なトランザクション削除
4. Order Repositoryのビルドエラー修正
5. Domain EntityからDB tagを削除
6. Analytics UsecaseのTransactionManager対応
7. 依存関係の整理

🔄 **継続的な改善が必要な項目（4項目）**
1. Infrastructure/Domainモデルの整理
2. DTO層の適切な分離
3. 時間生成の抽象化
4. エラーハンドリングの統一

プロジェクトのヘキサゴナルアーキテクチャは大幅に改善され、レイヤー間の依存関係が適切になりました。残りの項目については、優先度と影響範囲を考慮しながら段階的に対応することを推奨します。
