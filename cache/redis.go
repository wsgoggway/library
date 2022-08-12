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
	Store(key string, data interface{}, duration time.Duration, hstore bool, field ...string) (err error)
	Reset(key string) (err error)
	ResetByParent(parentKey string) (err error)
	ExistKey(key string) (ok bool)
	HExistKey(key, field string) (ok bool)
	Push(key string, data []byte) error
	Pop(key string, in interface{}) error
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

func (c *Cache) getHash(key string, field string) (data string, err error) {
	if data, err = c.redis.HGet(context.Background(), key, field).Result(); err != nil {
		return
	}
	return
}

func (c *Cache) getSimple(key string) (data string, err error) {
	if data, err = c.redis.Get(context.Background(), key).Result(); err != nil {
		return
	}
	return
}

func (c *Cache) Get(key string, in interface{}, hstore bool, field ...string) (err error) {
	var data string
	if hstore {
		if data, err = c.getHash(key, field[0]); err != nil {
			return
		}
	} else {
		if data, err = c.getSimple(key); err != nil {
			return
		}

	}

	err = json.Unmarshal([]byte(data), &in)
	if err != nil {
		return fmt.Errorf("seriliazarion error %s", err)
	}

	return nil
}

func (c *Cache) setHash(key string, data []byte, field string) (err error) {
	if _, err = c.redis.HSet(context.Background(), key, field, data).Result(); err != nil {
		return err
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

func (c *Cache) Store(key string, data interface{}, duration time.Duration, hstore bool, field ...string) (err error) {
	serializedData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("seriliazarion error %s", err)
	}

	if hstore {
		c.setHash(key, serializedData, field[0])
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
