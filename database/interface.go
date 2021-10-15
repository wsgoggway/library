package database

import (
	"context"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Transaction interface {
	Executor

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	Get(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	GetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Insert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Update(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Delete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
}

type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type IDatabase interface {
	GetResult(tx pgx.Tx, result interface{}, q string, args ...interface{}) error
	GetTransaction() (*Tx, error)
	Get(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	GetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Count(ctx context.Context, returnValue *int64, sql string, args ...interface{}) error
	Insert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Update(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	Delete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error
	GetContext() context.Context
	GetContextWithCache(ctx context.Context, key string, ttl ...time.Duration) context.Context
}
