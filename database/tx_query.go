package database

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"

	utils "github.com/wsgoggway/library/utils"
)

// Tx is IDatabase Transaction
type Tx struct {
	Transaction
	transaction pgx.Tx
	log         utils.Logger
}

/*

TRANSACTION

*/

func (tx *Tx) Get(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	return pgxscan.Select(ctx, tx, returnValue, sql, args...)
}

func (tx *Tx) GetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	return pgxscan.Get(ctx, tx.transaction, returnValue, sql, args...)
}

func (tx *Tx) Insert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if returnValue != nil {
		if err := pgxscan.Get(ctx, tx.transaction, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := tx.transaction.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Insert() {
			tx.log.Error(err)
		}
	}

	return nil
}

func (tx *Tx) Update(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if returnValue != nil {
		if err := pgxscan.Get(ctx, tx.transaction, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := tx.transaction.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Update() {
			tx.log.Error(err)
		}
	}

	return nil
}

func (tx *Tx) Delete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if returnValue != nil {
		if err := pgxscan.Get(ctx, tx.transaction, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := tx.transaction.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Delete() {
			tx.log.Error(err)
		}
	}

	return nil
}

func (tx *Tx) Commit(ctx context.Context) error {
	return tx.transaction.Commit(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	return tx.transaction.Rollback(ctx)
}
