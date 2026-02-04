package entities

import "time"

type Deposit struct {
	ID            uint64 `gorm:"primaryKey"`
	TransactionID uint64
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	Source        string
	// Provider
	// ExternalRef
	PaidAt time.Time
}
