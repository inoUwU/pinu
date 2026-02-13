package repositories

import (
	"context"
	"database/sql"

	"inoUwU/pinu/app/infrastructure/ctx"

	"github.com/uptrace/bun"
)

type TxRepository struct {
	db *bun.DB
}

func NewTxRepository(db *bun.DB) *TxRepository {
	return &TxRepository{db: db}
}

func (r *TxRepository) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.DoInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, _ bun.Tx) error {
		return fn(ctx)
	})
}

func (r *TxRepository) DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx bun.Tx) error) error {
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

	if err := fn(c, tx); err != nil {
		return err
	}

	done = true
	return tx.Commit()
}
