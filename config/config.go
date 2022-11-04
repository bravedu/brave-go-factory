package config

import (
	"fmt"
	"github.com/bravedu/brave-go-factory/modules/mysql_mod/db_dao"
	"sync"
)

var (
	Conf     *Config
	confOnce sync.Once
)

type Config struct {
	BaseDao *baseCnf
	YamlDao *YamlCnf
	DbDao   db_dao.IStore
}

func ConfInstance(env string) *Config {
	confOnce.Do(func() {
		Conf = new(Config)
		//配置文件读取
		Conf.initYaml(fmt.Sprintf("%s_conf.yaml", env))
		//处理数据库资源
		Conf.initDB()
	})
	return Conf
}

func ConfInstanceDev(env string) *Config {
	confOnce.Do(func() {
		Conf = new(Config)
		//配置文件读取
		Conf.initYaml(fmt.Sprintf("%s_conf.yaml", env))
		Conf.initNaCosCnf()
		//处理数据库资源
		Conf.initDB()
	})
	return Conf
}

func (c *Config) Close() {
	c.DbDao.Close()
}
