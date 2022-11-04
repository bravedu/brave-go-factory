package db_dao

import (
	"errors"
	"github.com/bravedu/brave-go-factory/modules/mysql_mod/db_struct"
	"gorm.io/gorm"
)

func (s *Store) GetGoodsInfo(id int) (*db_struct.SpGoods, error) {
	var goodsInfo *db_struct.SpGoods
	err := s.db.Model(&db_struct.SpGoods{}).Where("id = ?", id).First(&goodsInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return goodsInfo, err
}

func (s *Store) InsertGoodsInfo(g *db_struct.SpGoods) (int, error) {
	err := s.db.Create(g).Error
	return g.Id, err
}
