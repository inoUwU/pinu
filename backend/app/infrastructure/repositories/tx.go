package repositories

import (
	"context"
	"database/sql"

	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/infrastructure/ctx"

	"github.com/uptrace/bun"
)

// TxRepository トランザクション管理リポジトリ（port.TransactionManagerの実装）
type TxRepository struct {
	db *bun.DB
}

// NewTxRepository トランザクション管理リポジトリを生成する
func NewTxRepository(db *bun.DB) port.TransactionManager {
	return &TxRepository{db: db}
}

// DoInTx トランザクション内で関数を実行する
func (r *TxRepository) DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context) error) error {
	tx, err := r.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	// bunの `RunInTx`をベースに途中でcontextにトランザクションオブジェクトを入れる処理を追加
	c := context.WithValue(ctx, ctxkey.TxCtxKey, tx)

	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	if err := fn(c); err != nil {
		return err
	}

	done = true
	return tx.Commit()
}
