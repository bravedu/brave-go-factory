package config

import (
	"github.com/bravedu/brave-go-factory/modules/baidu_mod"
)

type baiDuCnf struct {
	AuditCnf baidu_mod.AuditCnf `yaml:"audit_cnf"`
}

type BaiDuCli struct {
	Audit *baidu_mod.AuditClientPool
}

func (c *Config) initBaidu() {
	//读取配置信息
	c.BaiDuCli.Audit = baidu_mod.AuditInstance(&c.YamlDao.BaiDuCnf.AuditCnf)
}
