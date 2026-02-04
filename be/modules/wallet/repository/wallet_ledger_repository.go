package repository

import (
	"context"
	"rialfu/wallet/database/entities"

	"gorm.io/gorm"
)

type (
	WalletLedgerRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data entities.WalletLedger) (entities.WalletLedger, error)
	}

	walletLedgerRepository struct {
		db *gorm.DB
	}
)

func NewWalletLedgerRepository(db *gorm.DB) WalletLedgerRepository {
	return &walletLedgerRepository{
		db: db,
	}
}
func (r *walletLedgerRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}
func (r *walletLedgerRepository) Create(ctx context.Context, tx *gorm.DB, data entities.WalletLedger) (entities.WalletLedger, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.WalletLedger{}, err
	}

	return data, nil
}
