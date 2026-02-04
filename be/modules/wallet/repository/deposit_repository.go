package repository

import (
	"context"
	"rialfu/wallet/database/entities"

	"gorm.io/gorm"
)

type (
	DepositRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data entities.Deposit) (entities.Deposit, error)
	}

	depositRepository struct {
		db *gorm.DB
	}
)

func NewDepositRepository(db *gorm.DB) DepositRepository {
	return &depositRepository{
		db: db,
	}
}
func (r *depositRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}
func (r *depositRepository) Create(ctx context.Context, tx *gorm.DB, data entities.Deposit) (entities.Deposit, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Deposit{}, err
	}

	return data, nil
}
