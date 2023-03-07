package cache

import (
	"context"
	"time"
)

func (c *Cache) hGet(key string, field string) (data string, err error) {
	if data, err = c.redis.HGet(context.Background(), key, field).Result(); err != nil {
		return
	}
	return
}

func (c *Cache) get(key string) (data string, err error) {
	if data, err = c.redis.Get(context.Background(), key).Result(); err != nil {
		return
	}
	return
}

func (c *Cache) setHash(key string, data []byte, field string, duration ...time.Duration) (err error) {
	dur := 1 * time.Minute
	if len(duration) > 0 {
		dur = duration[0]
	}

	if _, err = c.redis.HSet(context.Background(), key, field, data).Result(); err != nil {
		return err
	}

	if err = c.redis.Expire(context.Background(), key, dur).Err(); err != nil {
		return
	}
	return nil
}

func (c *Cache) setSimple(key string, data []byte, duration ...time.Duration) (err error) {
	dur := 1 * time.Minute
	if len(duration) > 0 {
		dur = duration[0]
	}

	if _, err = c.redis.Set(context.Background(), key, data, dur).Result(); err != nil {
		return err
	}
	return nil
}
