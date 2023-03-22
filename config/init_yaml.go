package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func (c *Config) initYaml(file string) {
	yamFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	ym := new(YamlCnf)
	err = yaml.Unmarshal(yamFile, ym)
	if err != nil {
		panic(err)
	}
	//配置文件属性
	c.YamlDao = ym
}

type YamlCnf struct {
	ProCnf    baseCnf    `yaml:"pro_cnf"`
	NaCos     naCosCnf   `yaml:"na_cos_cnf"`
	Db        db         `yaml:"database"`
	Jwt       jwt        `yaml:"jwt"`
	SuperDt   projectCnf `yaml:"super_dt"`
	Redis     redisCnf   `yaml:"redis""`
	ALiYunCnf aLiYunCnf  `yaml:"aliyun_cnf"`
	BaiDuCnf  baiDuCnf   `yaml:"bai_du_cnf"`
	Wxweb     wxweb      `yaml:"wxweb"`
	Wxh5      wxh5       `yaml:"wxh5"`
	Pay       pay        `yaml:"pay"`
}
