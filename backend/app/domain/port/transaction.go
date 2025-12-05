package port

import (
	"context"
	"database/sql"
)

// TransactionManager トランザクション管理のインターフェース（ポート）
// UsecaseレイヤーはこのPortを使用してトランザクションを管理する
type TransactionManager interface {
	// DoInTx トランザクション内で関数を実行する
	// トランザクションが正常に完了した場合はコミット、エラーが発生した場合はロールバックする
	DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context) error) error
}
