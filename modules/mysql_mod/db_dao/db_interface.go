package db_dao

import "github.com/bravedu/brave-go-factory/modules/mysql_mod/db_struct"

type IStore interface {
	//数据库基础方法
	BeginTx() (IStore, error)
	Rollback() error
	Commit() error
	Close()
	StoreClone() IStore

	//业务代码自定义方法
	Orders
	Goods
	Users
}

//order 模块数据库操作
type Orders interface {
}

type Goods interface {
	GetGoodsInfo(id int) (*db_struct.SpGoods, error)
	InsertGoodsInfo(g *db_struct.SpGoods) (int, error)
}

type Users interface {
}
