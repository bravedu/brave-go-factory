package redis_mod

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	redisOnce sync.Once
	rdb       *redis.Client
)

const Nil = redis.Nil
const OK = "OK"

type RedisClient struct {
	Rds *redis.Client
}

type RedisCnf struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}

func RedisInstance(cnf *RedisCnf) *RedisClient {
	redisOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     cnf.Addr,
			Password: cnf.Password,
			DB:       cnf.DB,
		})
	})
	return &RedisClient{
		Rds: rdb,
	}
}
