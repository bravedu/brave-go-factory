package main

import (
	"flag"
	"fmt"
	"github.com/bravedu/brave-go-factory/pkg/cmd/build_mysql"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var (
	config *YamlCnf
)

type Db struct {
	Ip       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Params   string `yaml:"params"`
}

type YamlCnf struct {
	Db          Db          `yaml:"database"`
	BuildStruct BuildStruct `yaml:"build_struct"`
}

type BuildStruct struct {
	BasePath        string `yaml:"base_path"`
	MysqlStructPath string `yaml:"mysql_struct_path"`
}

func initYaml(file string) {
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
	config = ym
}

func main() {
	pathFile := "../../dev_conf.yaml"
	if _, err := os.Stat(pathFile); err != nil {
		panic("数据库生成只允许在开发环境执行,并需要 dev_conf.yaml 文件")
		return
	}
	initYaml(pathFile)
	//只有开发环境才能执行
	parser()
}

//正常使用go build cli.go -table "表明" -pgname "包名"
func parser() {
	help := flag.Bool("help", false, "帮助, 常用: go build cli.go -table \"表明\" -packageName \"包名\"")
	h := flag.Bool("h", false, "帮助, 常用: go build cli.go -table \"表明\" -packageName \"包名\"")
	dsn := flag.String("dsn", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.Db.User, config.Db.Password, config.Db.Ip, config.Db.Port, config.Db.Database), "数据库dsn配置")
	file := flag.String("file", config.BuildStruct.BasePath+config.BuildStruct.MysqlStructPath, "mysql struct 保存路径")
	table := flag.String("table", "", "数据库表明")
	realNameMethod := flag.String("realNameMethod", "", "结构体对应的表名")
	packageName := flag.String("pgname", "db_struct", "生成的struct包名")
	tagKey := flag.String("tagKey", "gorm", "字段tag的key")
	prefix := flag.String("prefix", "", "表前缀")
	enableJsonTag := flag.Bool("enableJsonTag", true, "是否添加json的tag,默认false")
	// 开始
	flag.Parse()

	if *h || *help {
		flag.Usage()
		return
	}

	// 初始化
	t2t := build_mysql.NewTable2Struct()
	// 个性化配置
	t2t.Config(&build_mysql.T2tConfig{
		StructNameToHump: true,
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(*table).
		// 表前缀
		Prefix(*prefix).
		// 是否添加json tag
		EnableJsonTag(*enableJsonTag).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName(*packageName).
		// tag字段的key值,默认是orm
		TagKey(*tagKey).
		// 是否添加结构体方法获取表名
		RealNameMethod(*realNameMethod).
		// 生成的结构体保存路径
		SavePath(*file + *table + ".go").
		// 数据库dsn
		Dsn(*dsn).
		// 执行
		Run()

	if err != nil {
		log.Println(err.Error())
	}
}
