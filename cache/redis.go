package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	utils "github.com/wsgoggway/library/utils"
)

type ICache interface {
	Get(key string, in interface{}, hstore bool, field ...string) (err error)
	HGetAll(key string) (data map[string]string, err error)
	Store(key string, data interface{}, duration time.Duration, hstore bool, field ...string) (err error)
	Reset(key string) (err error)
	HRemove(key string, field ...string) error
	ResetByParent(parentKey string) (err error)
	ExistKey(key string) (ok bool)
	HExistKey(key, field string) (ok bool)
	Push(key string, data []byte) error
	Pop(key string, in interface{}) error
	HLen(key string) (uint64, error)
	LLen(key string) (uint64, error)
	RateLimit(key string, count int64, expire ...time.Duration) (ok bool)
}

type Cache struct {
	redis redis.UniversalClient
	log   utils.Logger
}

func (c *Cache) SetCacheImplementation(cacheClient redis.UniversalClient) {
	c.redis = cacheClient
}

func (c *Cache) SetLogger(l utils.Logger) {
	c.log = l
}

func (c *Cache) HGetAll(key string) (data map[string]string, err error) {
	return c.redis.HGetAll(context.Background(), key).Result()
}

func (c *Cache) Get(key string, in interface{}, hstore bool, field ...string) (err error) {
	var data string
	if hstore {
		if data, err = c.hGet(key, field[0]); err != nil {
			return
		}
	} else {
		if data, err = c.get(key); err != nil {
			return
		}

	}

	err = json.Unmarshal([]byte(data), &in)
	if err != nil {
		return fmt.Errorf("seriliazarion error %s", err)
	}

	return nil
}

func (c *Cache) Store(key string, data interface{}, duration time.Duration, hstore bool, field ...string) (err error) {
	serializedData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("seriliazarion error %s", err)
	}

	if hstore {
		c.setHash(key, serializedData, field[0], duration)
	} else {
		c.setSimple(key, serializedData, duration)
	}

	return nil
}

func (c *Cache) Reset(key string) (err error) {
	_, err = c.redis.Del(context.Background(), key).Result()
	return err
}

func (c *Cache) ResetByParent(parentKey string) (err error) {
	keys := c.redis.Keys(context.Background(), parentKey).Val()
	_, err = c.redis.Del(context.Background(), keys...).Result()
	return
}

func (c *Cache) ExistKey(key string) (ok bool) {
	if exists := c.redis.Exists(context.Background(), key).Val(); exists > 0 {
		ok = true
		return
	}

	return
}

func (c *Cache) HExistKey(key, field string) (ok bool) {
	return c.redis.HExists(context.Background(), key, field).Val()
}

func (c *Cache) Push(key string, data []byte) error {
	cmd := c.redis.LPush(context.Background(), key, data)
	return cmd.Err()
}

func (c *Cache) Pop(key string, in interface{}) error {
	bytes, err := c.redis.LPop(context.Background(), key).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(bytes), &in)
	if err != nil {
		return fmt.Errorf("seriliazarion error %s", err)
	}

	return nil
}

func (c *Cache) HLen(key string) (uint64, error) {
	return c.redis.HLen(context.Background(), key).Uint64()
}

func (c *Cache) LLen(key string) (uint64, error) {
	return c.redis.LLen(context.Background(), key).Uint64()
}

func (c *Cache) HRemove(key string, field ...string) error {
	return c.redis.HDel(context.Background(), key, field...).Err()
}

func (c *Cache) RateLimit(key string, count int64, expire ...time.Duration) (ok bool) {
	ctx := context.Background()
	res := c.redis.Incr(ctx, key)
	if res.Err() != nil {
		c.log.Errorf("RateLimit.Incr: %s", res.Err().Error())
		return false
	}

	if len(expire) > 0 {
		c.redis.Expire(ctx, key, expire[0])
	}

	if res.Val() >= count {
		return false
	}

	return true
}
