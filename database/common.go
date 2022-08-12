package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kelseyhightower/envconfig"
	cache "github.com/wsgoggway/library/cache"
	utils "github.com/wsgoggway/library/utils"
)

type ContextKey string

var (
	c = new(config)
)

const (
	ERROR_KEY_EMPTY                   = "cant empty cache key"
	CONTEXT_CACHE_MODE_KEY ContextKey = "cache"
	CONTEXT_CACHE_KEY      ContextKey = "key"
	CONTEXT_TTL_KEY        ContextKey = "ttl"
)

// Config read only.
type config struct {
	PostgresMasterAddr string `envconfig:"POSTGRES_MASTER_ADDR" required:"true"`
	PostgresSlaveAddr  string `envconfig:"POSTGRES_SLAVE_ADDR" required:"true"`
}

// Init initialization from environment variables
func init() {
	if err := envconfig.Process("", c); err != nil {
		log.Fatalf("db layer failed to load configuration: %s", err)
	}
}

var _ IDatabase = (*Database)(nil)
var _ Transaction = (*Tx)(nil)

func New(cache cache.ICache, log utils.Logger) (*Database, error) {
	dbl := &Database{cache: cache, log: log}
	var err error

	log.Info("Connect to master postgresql")
	masterConfig, err := pgxpool.ParseConfig(c.PostgresMasterAddr)
	if err != nil {
		return nil, err
	}
	masterConfig.ConnConfig.PreferSimpleProtocol = true
	dbl.writePool, err = pgxpool.ConnectConfig(context.Background(), masterConfig)
	if err != nil {
		return nil, err
	}
	log.Info("Connected...")

	log.Info("Connect to slaves postgresql")
	slaveConfig, err := pgxpool.ParseConfig(c.PostgresSlaveAddr)
	if err != nil {
		return nil, err
	}
	slaveConfig.ConnConfig.PreferSimpleProtocol = true
	dbl.readPool, err = pgxpool.ConnectConfig(context.Background(), slaveConfig)
	if err != nil {
		return nil, err
	}
	log.Info("Connected...")

	return dbl, nil
}
