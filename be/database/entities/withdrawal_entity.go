package entities

import "github.com/shopspring/decimal"

type Withdrawal struct {
	ID            uint64          `gorm:"primaryKey"`
	TransactionID uint64          // PK + FK
	Transaction   Transaction     `gorm:"foreignKey:TransactionID;references:ID"`
	Amount        decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	Fee           decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	TargetBank    uint            `gorm:"column:target_bank"`
	Bank          MasterBank      `gorm:"foreignKey:target_bank;references:ID"`
	TargetAccount string          `gorm:"type:varchar(50)"`
}
