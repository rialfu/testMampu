package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type WalletLedger struct {
	ID            uint64 `gorm:"primaryKey"`
	WalletID      uint64
	Wallet        Wallet `gorm:"foreignKey:WalletID;references:ID"`
	TransactionID uint64
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	Direction     string      // DEBIT / CREDIT
	Amount        decimal.Decimal
	BalanceBefore decimal.Decimal
	BalanceAfter  decimal.Decimal
	CreatedAt     time.Time
}
