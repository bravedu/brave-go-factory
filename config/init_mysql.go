package config

import (
	"github.com/bravedu/brave-go-factory/modules/mysql_mod/db_dao"
	"gopkg.in/yaml.v3"
)

//数据库连接句柄
func (c *Config) initDB() {
	//读取配置信息
	dbConf, err := yaml.Marshal(c.YamlDao.Db)
	if err != nil {
		panic(err)
	}
	c.DbDao = db_dao.ConnectDbDao(dbConf)
}

type db db_dao.DB
