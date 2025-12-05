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
	// 
	// パラメータ:
	//   ctx: コンテキスト
	//   opts: トランザクションオプション（分離レベルなど）
	//   fn: トランザクション内で実行する関数。この関数に渡されるcontextには
	//       トランザクション情報が含まれており、Repository実装はこのcontextを
	//       使用してトランザクション内でのDB操作を実行できる
	DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context) error) error
}
