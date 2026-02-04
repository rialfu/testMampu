package repository

import (
	"context"
	"errors"
	"rialfu/wallet/database/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	TransactionRepository interface {
		CreateTransaction(ctx context.Context, tx *gorm.DB, data entities.Transaction) (entities.Transaction, error)
		CheckReferenceNo(ctx context.Context, tx *gorm.DB, ref string, isLock bool) (entities.Transaction, bool, error)
		Update(ctx context.Context, tx *gorm.DB, data entities.Transaction) error
	}

	transactionRepository struct {
		db *gorm.DB
	}
)

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
func (r *transactionRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}
func (r *transactionRepository) CreateTransaction(ctx context.Context, tx *gorm.DB, data entities.Transaction) (entities.Transaction, error) {
	db := r.getDB(ctx)
	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Transaction{}, err
	}

	return data, nil
}
func (r *transactionRepository) CheckReferenceNo(ctx context.Context, tx *gorm.DB, ref string, isLock bool) (entities.Transaction, bool, error) {
	db := r.getDB(ctx)
	var data entities.Transaction
	if isLock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	if err := db.WithContext(ctx).Where("reference_no = ?", ref).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data, false, nil
		}
		return data, false, err
	}

	return data, true, nil
}
func (r *transactionRepository) Update(ctx context.Context, tx *gorm.DB, data entities.Transaction) error {
	db := r.getDB(ctx)
	if err := db.WithContext(ctx).Model(&data).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
