package db_struct

import "time"

type TableBaseFields struct {
	Id         int       `gorm:"column:id" json:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}
