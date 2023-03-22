package config

import (
	"github.com/bravedu/brave-go-factory/modules/aliyun_mod"
)

type aLiYunCnf struct {
	ALiYunSms       aliyun_mod.ALiYunSms     `yaml:"aliyun_sms"`
	ALiYunOss       *aliyun_mod.OssCnf       `yaml:"aliyun_oss"`
	AliYunOssUpload *aliyun_mod.OssUploadCnf `yaml:"aliyun_oss_upload"`
}

type AliYunCli struct {
	SMSCli *aliyun_mod.SmsPool
	OssCli *aliyun_mod.OssPool
}

func (c *Config) initALiYun() {
	//读取配置信息
	c.AliYunCli.SMSCli = aliyun_mod.SmsPoolInstance(c.YamlDao.ALiYunCnf.ALiYunSms)
	c.AliYunCli.OssCli = aliyun_mod.OssClient(c.YamlDao.ALiYunCnf.ALiYunOss, c.YamlDao.ALiYunCnf.AliYunOssUpload)
}
