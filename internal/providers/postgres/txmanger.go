package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type txContextKey struct{}

var key = txContextKey{}

type (
	TxBeginner interface {
		Begin(context.Context) (pgx.Tx, error)
	}
	TxManger struct {
		txBeginner TxBeginner
	}
)

func NewTxManger(txBeginner TxBeginner) *TxManger {
	return &TxManger{txBeginner: txBeginner}
}

func (tm *TxManger) WithinTransaction(ctx context.Context, f func(context.Context) bool) error {
	tx, err := tm.txBeginner.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	txCtx := context.WithValue(ctx, key, tx)
	if f(txCtx) {
		return tx.Commit(ctx)
	}
	return tx.Rollback(ctx)
}

func GetTxFromContext(ctx context.Context) any {
	return ctx.Value(key)
}
