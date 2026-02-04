package entities

import "time"

type MasterBank struct {
	ID        uint      `gorm:"primaryKey"`
	BankCode  string    `gorm:"type:varchar(10);uniqueIndex;not null"`
	BankName  string    `gorm:"type:varchar(100);not null"`
	IsActive  *bool     `gorm:"default:true"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
