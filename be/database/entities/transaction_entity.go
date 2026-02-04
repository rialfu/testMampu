package entities

import (
	"time"
)

type Transaction struct {
	ID              uint64    `gorm:"primaryKey"`
	ReferenceNo     string    `gorm:"uniqueIndex;type:varchar(25);not null;"`
	WalletID        uint64    `gorm:"column:wallet_id"`
	Wallet          Wallet    `gorm:"foreignKey:wallet_id;references:ID"`
	UserID          uint64    `gorm:"column:user_id"`
	User            User      `gorm:"foreignKey:user_id;references:ID"`
	TransactionType string    `gorm:"type:varchar(20)"`    // WITHDRAWAL, TOPUP
	Status          int8      `gorm:"default:3;not null;"` //1 success, 2 failed, 3pending
	DateTrans       time.Time `gorm:"column:date_trans"`
	CreatedAt       time.Time
}
