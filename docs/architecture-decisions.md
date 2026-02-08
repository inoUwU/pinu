# アーキテクチャ決定議事録 (Architecture Decision Record)

**日付**: 2025年7月29日  
**参加者**: 開発チーム  
**議題**: ポート&アダプターアーキテクチャの実装見直し

## 背景

Go言語でのクリーンアーキテクチャ（ポート&アダプター）実装において、サービス層の必要性とアーキテクチャの正しい実装について検討を行った。

## 検討事項

### 1. 依存性の型（ポインタ vs 非ポインタ）

**質問**: サービスやリポジトリの依存性はポインタ型と非ポインタ型のどちらが良いか？

**決定**: **ポインタ型を採用**

**理由**:
- メソッドのレシーバーがポインタ型であることが多い
- 依存性注入（DI）との相性が良い
- パフォーマンス面でのメリット（値のコピーが発生しない）
- Go言語の慣習に従っている

**実装例**:
```go
type UserController struct {
    userService *services.UserService
    authService *services.AuthService
    userRepo    *repositories.UserRepository
}
```

### 2. インターフェースの型定義

**問題**: `*repositories.IUserRepository`エラーの発生

**決定**: **インターフェース型にはポインタを使用しない**

**理由**:
- Goではインターフェース型自体を使用する
- インターフェースへのポインタは不要で、エラーの原因となる

**修正前**:
```go
type UserUsecaseImpl struct {
    userRepo *repositories.IUserRepository // ❌ エラー
}
```

**修正後**:
```go
type UserUsecaseImpl struct {
    userRepo repositories.IUserRepository // ✅ 正しい
}
```

### 3. サービス層の必要性

**問題**: `domain/services`層がユースケース層の単純なラッパーになっている

**決定**: **サービス層を削除**

**理由**:
- 不要な抽象化層となっている
- 依存性の方向が逆転している（domain → usecases）
- 実質的な処理はユースケース層で行われている
- バリデーションなどのビジネスロジックはユースケース層に実装可能

**削除されたファイル**:
- `e:\Github\pinu\backend\app\domain\services\` フォルダ全体

## 最終的なアーキテクチャ

### アーキテクチャ構成

```
┌─────────────────┐
│   Controllers   │ ← HTTPハンドラー（プライマリアダプター）
│   (Fiber)       │
└─────────────────┘
         ↓
┌─────────────────┐
│    Usecases     │ ← ビジネスロジック・バリデーション
│  (Application)  │   アプリケーション層
└─────────────────┘
         ↓
┌─────────────────┐
│ Repository      │ ← ポート（インターフェース）
│ Interface       │   ドメイン層
│   (Domain)      │
└─────────────────┘
         ↓
┌─────────────────┐
│ Repository      │ ← セカンダリアダプター
│ Implementation  │   インフラストラクチャ層
│(Infrastructure) │
└─────────────────┘
         ↓
┌─────────────────┐
│    Database     │
└─────────────────┘
```

### 各層の責任

1. **Controllers**: 
   - HTTPリクエスト/レスポンスの処理
   - ユースケースの呼び出し
   - プレゼンテーション層

2. **Usecases**: 
   - ビジネスロジックの実装
   - バリデーション処理
   - アプリケーション固有のロジック
   - リポジトリインターフェースの使用

3. **Repository Interface**: 
   - データアクセスの抽象化
   - ドメイン層のポート定義

4. **Repository Implementation**: 
   - 具体的なデータアクセス実装
   - データベース操作

### 依存性の方向

```
Controllers → Usecases → Repository Interface ← Repository Implementation
```

- 内側の層（ドメイン）は外側の層（インフラ）に依存しない
- 依存性逆転の原則に従っている

## 実装の変更点

### 1. UserController の修正

**変更前**:
```go
type UserController struct {
    userService services.IUserService
}
```

**変更後**:
```go
type UserController struct {
    userUsecase usecases.IUserUsecase
}
```

### 2. UserUsecase の構造体名修正

**変更前**:
```go
type userUsecase struct { // 小文字（非公開）
    userRepo repositories.IUserRepository
}
```

**変更後**:
```go
type UserUsecaseImpl struct { // 大文字（公開）
    userRepo repositories.IUserRepository
}
```

## 利点

1. **シンプルな構造**: 不要な抽象化を排除
2. **理解しやすさ**: レイヤー数の削減
3. **保守性の向上**: 責任の明確化
4. **テスタビリティ**: ユースケース単位でのテスト容易性
5. **Go慣習への準拠**: Goコミュニティのベストプラクティス

## 今後の拡張方針

- 複雑なドメインロジックが必要になった場合は、その時点でドメインサービスの追加を検討
- バリデーションとビジネスロジックはユースケース層で実装
- 新しい機能追加時も同じアーキテクチャパターンを踏襲

## 承認

**決定者**: 開発チーム  
**承認日**: 2025年7月29日  
**ステータス**: 承認済み・実装完了
