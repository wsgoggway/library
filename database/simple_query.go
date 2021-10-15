package database

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	cache "github.com/wsgoggway/library/cache"
	utils "github.com/wsgoggway/library/utils"
)

// Database is IDatabase implementation
type Database struct {
	readPool  *pgxpool.Pool
	writePool *pgxpool.Pool
	cache     cache.ICache
	log       utils.Logger
}

/*

DATABASE

*/

func (dbl *Database) GetResult(tx pgx.Tx, result interface{}, q string, args ...interface{}) error {
	return pgxscan.Get(dbl.GetContext(), tx, result, q, args...)
}

func (dbl *Database) GetTransaction() (tx *Tx, err error) {
	tx = new(Tx)
	tx.transaction, err = dbl.writePool.Begin(dbl.GetContext())
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (dbl *Database) GetContext() context.Context { return context.Background() }
func (dbl *Database) GetContextWithCache(ctx context.Context, key string, ttl ...time.Duration) context.Context {
	if ctx == nil {
		ctx = dbl.GetContext()
	}
	ctx = context.WithValue(ctx, CONTEXT_CACHE_MODE_KEY, true)
	ctx = context.WithValue(ctx, CONTEXT_CACHE_KEY, key)

	localTTL := time.Duration(-1)
	if len(ttl) > 0 {
		localTTL = ttl[0]
	}

	ctx = context.WithValue(ctx, CONTEXT_TTL_KEY, localTTL)
	return ctx
}

func (dbl *Database) Get(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if _, ok := ctx.Value(CONTEXT_CACHE_MODE_KEY).(bool); ok {
		key, ok := ctx.Value(CONTEXT_CACHE_KEY).(string)
		if !ok {
			return fmt.Errorf(ERROR_KEY_EMPTY)
		}

		if dbl.cache.ExistKey(key) {
			return dbl.cache.Get(key, returnValue, false)
		} else {
			if err := pgxscan.Select(ctx, dbl.readPool, returnValue, sql, args...); err != nil { // get data from db
				return err
			}

			ttl := ctx.Value(CONTEXT_TTL_KEY).(time.Duration)
			if err := dbl.cache.Store(key, returnValue, ttl, false); err != nil { // store cache
				return err
			}

			return nil
		}
	}

	// without cache
	return pgxscan.Select(ctx, dbl.readPool, returnValue, sql, args...)
}

func (dbl *Database) GetOne(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	if _, ok := ctx.Value(CONTEXT_CACHE_MODE_KEY).(bool); ok {
		key, ok := ctx.Value(CONTEXT_CACHE_KEY).(string)
		if !ok {
			return fmt.Errorf(ERROR_KEY_EMPTY)
		}

		if dbl.cache.ExistKey(key) {
			return dbl.cache.Get(key, returnValue, false)
		} else {
			if err := pgxscan.Get(ctx, dbl.readPool, returnValue, sql, args...); err != nil { // get data from db
				if pgxscan.NotFound(err) {
					return fmt.Errorf("no rows")
				}
				return err
			}

			ttl := ctx.Value(CONTEXT_TTL_KEY).(time.Duration)
			if err := dbl.cache.Store(key, returnValue, ttl, false); err != nil { // store cache
				return err
			}

			return nil
		}
	}

	// without cache
	return pgxscan.Get(ctx, dbl.readPool, returnValue, sql, args...)
}

func (dbl *Database) Count(ctx context.Context, returnValue *int64, sql string, args ...interface{}) error {
	if _, ok := ctx.Value(CONTEXT_CACHE_MODE_KEY).(bool); ok {
		key, ok := ctx.Value(CONTEXT_CACHE_KEY).(string)
		if !ok {
			return fmt.Errorf(ERROR_KEY_EMPTY)
		}

		if dbl.cache.ExistKey(key) {
			return dbl.cache.Get(key, returnValue, false)
		} else {
			err := dbl.readPool.QueryRow(ctx, sql, args...).Scan(&returnValue)
			if err != nil {
				return err
			}

			ttl := ctx.Value(CONTEXT_TTL_KEY).(time.Duration)
			if err := dbl.cache.Store(key, returnValue, ttl, false); err != nil { // store cache
				return err
			}

			return nil
		}
	}

	// without cache
	err := dbl.readPool.QueryRow(ctx, sql, args...).Scan(&returnValue)
	if err != nil {
		return err
	}

	return nil
}

func (dbl *Database) Insert(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	conn, err := dbl.writePool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}

	if returnValue != nil {
		if err := pgxscan.Get(ctx, conn, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := conn.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Insert() {
			dbl.log.Error(err)
		}
	}

	return nil
}

func (dbl *Database) Update(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	conn, err := dbl.writePool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}

	if returnValue != nil {
		if err := pgxscan.Get(ctx, conn, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := conn.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Update() {
			dbl.log.Error(err)
		}
	}

	return nil
}

func (dbl *Database) Delete(ctx context.Context, returnValue interface{}, sql string, args ...interface{}) error {
	conn, err := dbl.writePool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}

	if returnValue != nil {
		if err := pgxscan.Get(ctx, conn, returnValue, sql, args...); err != nil {
			return err
		}
	} else {
		cmd, err := conn.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		if !cmd.Delete() {
			dbl.log.Error(err)
		}
	}

	return nil
}
