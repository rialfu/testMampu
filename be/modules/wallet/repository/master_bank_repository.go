package repository

import (
	"context"
	"errors"
	"rialfu/wallet/database/entities"

	"gorm.io/gorm"
)

type (
	MasterBankRepositry interface {
		GetAll(ctx context.Context, tx *gorm.DB) ([]entities.MasterBank, error)
		GetById(ctx context.Context, tx *gorm.DB, id string) (entities.MasterBank, bool, error)
	}

	masterBankRepositry struct {
		db *gorm.DB
	}
)

func NewMasterBankRepositry(db *gorm.DB) MasterBankRepositry {
	return &masterBankRepositry{
		db: db,
	}
}
func (r *masterBankRepositry) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}
func (r *masterBankRepositry) GetAll(ctx context.Context, tx *gorm.DB) ([]entities.MasterBank, error) {
	db := r.getDB(ctx)
	var datas []entities.MasterBank
	if err := db.WithContext(ctx).Take(&datas).Error; err != nil {
		return datas, err
	}
	return datas, nil
}
func (r *masterBankRepositry) GetById(ctx context.Context, tx *gorm.DB, id string) (entities.MasterBank, bool, error) {
	db := r.getDB(ctx)
	var data entities.MasterBank
	if err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}
