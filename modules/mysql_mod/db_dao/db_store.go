package db_dao

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"sync"
	"time"
)

var storeGormDbDao IStore
var daoOnce sync.Once

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) BeginTx() (IStore, error) {
	// begin自动开启一个新的数据库连接，必须用begin后的*DB连接（而不是原连接s本身）回滚或提交
	// 即使是新连接也不能事务套事务
	d := s.db.Begin()
	return NewStore(d), d.Error
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) Commit() error {
	return s.db.Commit().Error
}

func (s *Store) Close() {
	//现在不需要关闭了
}

func (s *Store) StoreClone() IStore {
	return NewStore(s.db)
}

//是实实在在对外提供服务
func ConnectDbDao(dbByte []byte) IStore {
	//只执行一次-数据库连接
	daoOnce.Do(func() {

		c := new(DB)
		err := yaml.Unmarshal(dbByte, c)
		if err != nil {
			panic(err)
		}
		dnsStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.User, c.Password, c.Ip, c.Port, c.Database, c.Params)
		db1Dsn := mysql.New(mysql.Config{
			DSN:                       dnsStr, // DSN data source name
			DefaultStringSize:         256,    // string 类型字段的默认长度
			DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   false,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,  // 根据当前 MySQL 版本自动配置
		})
		db, err := gorm.Open(db1Dsn, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             500 * time.Millisecond,
					Colorful:                  true,
					IgnoreRecordNotFoundError: true,
					LogLevel:                  logger.LogLevel(c.LogMode),
				}),
		})
		if err != nil {
			panic(err)
		}
		//判断是否存在只读实例
		if len(c.SlaveArray) > 0 {
			replicas := make([]gorm.Dialector, 0, len(c.SlaveArray))
			for _, v := range c.SlaveArray {
				dnsSlaveStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", v.User, v.Password, v.Ip, v.Port, v.Database, v.Params)
				dbSlaveDsn := mysql.New(mysql.Config{
					DSN:                       dnsSlaveStr, // DSN data source name
					DefaultStringSize:         256,         // string 类型字段的默认长度
					DisableDatetimePrecision:  true,        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
					DontSupportRenameIndex:    true,        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
					DontSupportRenameColumn:   false,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
					SkipInitializeWithVersion: false,       // 根据当前 MySQL 版本自动配置
				})
				replicas = append(replicas, dbSlaveDsn)
			}
			err = db.Use(
				dbresolver.Register(dbresolver.Config{
					//Sources: []gorm.Dialector{mysql.Open(dnsStr)},
					Replicas: replicas,
					// sources/replicas 负载均衡策略
					Policy: dbresolver.RandomPolicy{},
				}).
					SetConnMaxIdleTime(time.Hour).
					SetConnMaxLifetime(time.Hour).
					SetMaxOpenConns(c.MaxOpenConns).
					SetMaxIdleConns(c.MaxIdleConns),
			)
		}
		if err != nil {
			panic(err)
		}
		storeGormDbDao = NewStore(db)
	})
	return storeGormDbDao
}
