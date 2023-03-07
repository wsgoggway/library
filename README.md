# The Library collection for a database and cache layers

## Cache
Layer for redis cache


### Environment variable
```shell
export REDIS_ADDRS=localhost:6379
export REDIS_PASSWORD=12345qwerty
```

### Code
```go
redisClient := redis.NewClient(
    &redis.Options{
        Addr:         conf.RedisAddr,
        Password:     conf.RedisPassword,
        DialTimeout:  3,
        ReadTimeout:  3,
        WriteTimeout: 3,
    },
)
if err := redisClient.Conn(ctx); err != nil {
    panic(err)
}
defer redisClient.Close()

cacheIplm := new(cache.Cache) // package github.com/justfigurs/library/cache
cacheIplm.SetLogger(log)      // some logger like zap or logrus or zerolog
cacheIplm.SetCacheImplementation(redisClient)
```

---

## Database
Layer for database

### Environment variable
```shell
export POSTGRES_MASTER_ADDR=postgres://user:password@host:master_port/db?param1=on&param2=off
export POSTGRES_SLAVE_ADDR=postgres://user:password@host:slave_port/db?param1=on&param2=off
```

### This connection required the cache layer ([from this the library](#cache))
```go
db := database.New(cacheIplm, log)

// ... some code ...
```