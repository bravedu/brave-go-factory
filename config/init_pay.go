package config

import (
	"github.com/go-pay/gopay/wechat"
)

type pay struct {
	Appid      string `yaml:"appid"`
	Mchid      string `yaml:"mchid"`
	Mchkey     string `yaml:"mchkey"`
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
	Notifyurl  string `yaml:"notifyurl"`
}

func (c *Config) initWeChatPayV2() {
	apiKey := c.YamlDao.Pay.Mchkey
	client := wechat.NewClient(c.YamlDao.Pay.Appid, c.YamlDao.Pay.Mchid, apiKey, true)
	client.SetCountry(wechat.China)
	c.WeChatPayCli = client
}
