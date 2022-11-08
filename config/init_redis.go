package config

import (
	"github.com/bravedu/brave-go-factory/modules/redis_mod"
	"gopkg.in/yaml.v3"
)

type redisCnf redis_mod.RedisCnf

func (c *Config) initRedis() {
	//读取配置信息
	redisByte, err := yaml.Marshal(c.YamlDao.Redis)
	if err != nil {
		panic(err)
	}
	rdsCnf := new(redis_mod.RedisCnf)
	err = yaml.Unmarshal(redisByte, rdsCnf)
	if err != nil {
		panic("redis unmarshal config error")
	}
	c.RedisCli = redis_mod.RedisInstance(rdsCnf)
}
