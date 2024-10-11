package dao

import (
	"context"
	"time"

	"outback/geeke/stack/items"

	"gorm.io/gorm"
)

type SearchDal interface {
	SearchNameCode(ctx context.Context, param *items.SearchNameCodeParam) ([]items.NameCode, error)
}

type SearchDao struct {
	db *gorm.DB
}

func (dal *SearchDao) SearchNameCode(ctx context.Context, param items.SearchNameCodeParam) ([]items.NameCode, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	res := make([]items.NameCode, 0)
	db := dal.db.WithContext(ctx).Table(new(items.NameCode).TableName())
	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}
	if param.Code != "" {
		db = db.Where("code = ?", param.Code)
	}
	err := db.Find(&res).Error
	return res, err
}

func NewSearchDao(db *gorm.DB) *SearchDao {
	return &SearchDao{db: db}
}
