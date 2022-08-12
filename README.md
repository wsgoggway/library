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


## Sender
Layer for work with sender

```go
client := New("http://http-sender-reg.http-sender.svc.k8s.datapro", // or http://http-sender-reg.http-sender.svc.k8s.dataline
    map[string]string{},
    func(ms int64) {})

d := new(EmailSendRequest)
d.To = "vladimirm@example.com"
d.Key, _ = uuid.GenerateUUID()
d.Title = "Blizzard. World of warcraft. карта времени (30 дней)"
d.Username = "Василий Васичкин"

err := client.SendEmail(d)
```

## Finauth
Authorization api client

```go
c := New("https://payments.wildberries.ru/authtest/connect/token")
// or "https://payments.wildberries.ru/auth/connect/token"

v, err := c.GetToken(&GetTokenInput{
    ClientID:     "wbdigital-client",
    ClientSecret: "jh8747w839ieyt848wiedkjf43weuew",
    Scopes:       "fiscalization-creator",
})
```