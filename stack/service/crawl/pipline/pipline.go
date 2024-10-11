package pipline

import (
	"context"
	"time"

	"outback/geeke/stack/items"

	"gorm.io/gorm"
)

type CreateDal interface {
	CreateProfile(ctx context.Context, data *items.Profile) error
	CreateBalance(ctx context.Context, data *items.Balance) error
	CreateCashFlow(ctx context.Context, data *items.CashFlow) error
	CreateNameCode(ctx context.Context, data *items.NameCode) error
	// SearchNameCode(ctx context.Context, param *items.SearchNameCodeParam) ([]items.NameCode, error)
}

type CreateDao struct {
	db *gorm.DB
}

func (dal *CreateDao) SearchNameCode(ctx context.Context, param items.SearchNameCodeParam) ([]items.NameCode, error) {
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

func (dal *CreateDao) CreateBalance(ctx context.Context, data *items.Balance) error {
	data.Serialize()

	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	db := dal.db.WithContext(ctx).Table(data.TableName())
	if err := db.Where("unique_id = ?", data.UniqueID).Delete(&items.Balance{}).Error; err != nil {
		return err
	}

	err := db.Create(data).Error
	return err
}

func (dal *CreateDao) CreateCashFlow(ctx context.Context, data *items.CashFlow) error {
	data.Serialize()

	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	db := dal.db.WithContext(ctx).Table(data.TableName())

	if err := db.Where("unique_id = ?", data.UniqueID).Delete(&items.CashFlow{}).Error; err != nil {
		return err
	}

	err := db.Create(data).Error
	return err
}

func (dal *CreateDao) CreateProfile(ctx context.Context, data *items.Profile) error {

	data.Serialize()
	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	db := dal.db.Table(data.TableName()).WithContext(ctx)

	if err := db.Where("unique_id = ?", data.UniqueID).
		Delete(&items.Profile{}).Error; err != nil {
		return err
	}

	return db.Create(data).Error
}

func (dal *CreateDao) CreateNameCode(ctx context.Context, data *items.NameCode) error {
	data.Serialize()
	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	db := dal.db.Table(data.TableName()).WithContext(ctx)
	if err := db.Where("code = ?", data.Code).
		Delete(&items.NameCode{}).Error; err != nil {
		return err
	}
	err := db.Create(data).Error
	return err
}

func NewCreate(db *gorm.DB) *CreateDao {
	return &CreateDao{db: db}
}
