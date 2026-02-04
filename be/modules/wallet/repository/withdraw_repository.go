package repository

import (
	"context"
	"rialfu/wallet/database/entities"

	"gorm.io/gorm"
)

type (
	WithdrawRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data entities.Withdrawal) (entities.Withdrawal, error)
	}

	withdrawRepository struct {
		db *gorm.DB
	}
)

func NewWithdrawRepository(db *gorm.DB) WithdrawRepository {
	return &withdrawRepository{
		db: db,
	}
}
func (r *withdrawRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}
func (r *withdrawRepository) Create(ctx context.Context, tx *gorm.DB, data entities.Withdrawal) (entities.Withdrawal, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Withdrawal{}, err
	}

	return data, nil
}
