package config

import (
	"fmt"
	"github.com/bravedu/brave-go-factory/pkg/util"
	"github.com/bravedu/brave-go-factory/pkg/xlog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"strings"
)

type naCosCnf struct {
	IpAddR      string `yaml:"ip_addr"`
	ContextPath string `yaml:"context_path"`
	Port        uint64 `yaml:"port"`
	TimeOut     int    `yaml:"time_out"`
	Scheme      string `yaml:"scheme"`
	AppName     string `yaml:"app_name"`
	NameSpaceId string `yaml:"name_space_id"`
	DataId      string `yaml:"data_id"`
	Group       string `yaml:"group"`
	UserName    string `yaml:"user_name"`
	UserPwd     string `yaml:"user_pwd"`
}

func (c *Config) initNaCosCnf() {
	projectEnv := strings.ToLower(util.GetEnvDefault("APP_ENV", c.YamlDao.ProCnf.Env))
	projectName := strings.ToLower(util.GetEnvDefault("PROJECT_NAME", c.YamlDao.ProCnf.ProjectName))
	if projectEnv == "" {
		projectEnv = "online"
	}
	naCosLogPath := fmt.Sprintf("logs/%s/nacos_logs", projectName)
	//日志目录不存在创建
	util.CheckDirAndCreate(naCosLogPath)
	clientConfig := constant.ClientConfig{
		AppName:             "muse-api",
		TimeoutMs:           5000,
		NamespaceId:         fmt.Sprintf("%s_id", projectEnv),
		NotLoadCacheAtStart: true,
		LogDir:              util.GetEnvDefault("logs_path", "logs/na_cos_cnf"),
		CacheDir:            util.GetEnvDefault("logs_path", "logs/"),
		Username:            "system_read_user",
		Password:            "system_read_user",
		//LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      c.YamlDao.NaCos.IpAddR,
			ContextPath: c.YamlDao.NaCos.ContextPath,
			Port:        c.YamlDao.NaCos.Port,
			Scheme:      c.YamlDao.NaCos.Scheme,
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: c.YamlDao.NaCos.DataId,
		Group:  c.YamlDao.NaCos.Group})

	if err != nil {
		xlog.Error("获取配置信息报错", err)
		panic("获取配置信息出错-")
	}
	err = yaml.Unmarshal([]byte(content), &c.YamlDao)
	if err != nil {
		panic("yaml解析配置文件出错")
	}
	// 动态监听
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: c.YamlDao.NaCos.DataId,
		Group:  c.YamlDao.NaCos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			err = yaml.Unmarshal([]byte(data), &c.YamlDao)
			if err != nil {
				xlog.Error("监听配置出错", err)
			}
		},
	})
	if err != nil {
		xlog.Error("监听配置出错", err)
	}
}
