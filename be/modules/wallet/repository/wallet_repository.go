package repository

import (
	"context"
	"errors"
	"rialfu/wallet/database/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	WalletRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data entities.Wallet) (entities.Wallet, error)
		GetByUserId(ctx context.Context, tx *gorm.DB, id string, isLock bool) (entities.Wallet, bool, error)
		GetById(ctx context.Context, tx *gorm.DB, id string, isLock bool) (entities.Wallet, bool, error)
		UpdateBalance(ctx context.Context, tx *gorm.DB, data entities.Wallet) error
	}

	walletRepository struct {
		db *gorm.DB
	}
)

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{
		db: db,
	}
}
func (r *walletRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}
func (r *walletRepository) Create(ctx context.Context, tx *gorm.DB, data entities.Wallet) (entities.Wallet, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Wallet{}, err
	}

	return data, nil
}

func (r *walletRepository) GetByUserId(ctx context.Context, tx *gorm.DB, id string, isLock bool) (entities.Wallet, bool, error) {
	db := r.getDB(ctx)
	var data entities.Wallet
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	if err := db.WithContext(ctx).Where("user_id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data, false, nil
		}
		return data, false, err
	}

	return data, true, nil
}
func (r *walletRepository) GetById(ctx context.Context, tx *gorm.DB, id string, isLock bool) (entities.Wallet, bool, error) {
	db := r.getDB(ctx)
	var data entities.Wallet
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	if err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data, false, nil
		}
		return data, false, err
	}

	return data, true, nil
}
func (r *walletRepository) UpdateBalance(ctx context.Context, tx *gorm.DB, data entities.Wallet) error {
	db := r.getDB(ctx)
	if err := db.WithContext(ctx).Model(&data).Updates(entities.Wallet{Balance: data.Balance}).Error; err != nil {
		return err
	}
	return nil
}
