package port

import "context"

// UnitOfWork トランザクション境界のポート
// 実装詳細（DB/ORM）はインフラ層に閉じ込める。
type UnitOfWork interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
}
