package helpers

import (
	"context"

	"gorm.io/gorm"
)

type Transactor interface {
	WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
}

type gormTransactor struct {
	db *gorm.DB
}

func NewGormTransactor(db *gorm.DB) Transactor {
	return &gormTransactor{db: db}
}

func (t *gormTransactor) WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, "DB_TX", tx)
		return fn(txCtx)
	})
}
