package db_struct

type SpGoods struct {
	TableBaseFields
	GName string `gorm:"column:g_name" json:"g_name"`
	CId   int    `gorm:"column:c_id" json:"c_id"`
	GNums int    `gorm:"column:g_nums" json:"g_nums"`
	IsHot int    `gorm:"column:is_hot" json:"is_hot"`
}
