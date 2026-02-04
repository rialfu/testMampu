package entities

import "time"

type InformationUser struct {
	ID         uint64 `gorm:"primaryKey"`
	NIK        string `gorm:"varchar(16)"`
	IsVerified *bool  `gorm:"default:false"`
	UserID     uint64 `gorm:"index"`
	// Role       string `gorm:"type:varchar(50);not null;default:'user'" json:"role"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
