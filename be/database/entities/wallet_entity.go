package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID          uint64          `gorm:"primaryKey"`
	UserID      uint64          `gorm:"index"`
	User        User            `gorm:"foreignKey:UserID"`
	Balance     decimal.Decimal `gorm:"type:decimal(16,2);default:0"`
	Transaction []Transaction   `gorm:"foreignKey:WalletID"`
	UpdatedAt   time.Time
}
