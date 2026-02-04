package entities

import (
	"time"
)

type User struct {
	ID          uint64          `gorm:"primaryKey"`
	Name        string          `gorm:"type:varchar(100);not null" json:"name"`
	Email       string          `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	TelpNumber  string          `gorm:"type:varchar(20);index;not null;" json:"telp_number"`
	Password    string          `gorm:"type:varchar(255);not null" json:"password"`
	IsActive    *bool           `gorm:"default:true"`
	Information InformationUser `gorm:"foreignKey:UserID"`
	// Role       string `gorm:"type:varchar(50);not null;default:'user'" json:"role"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// IsVerified bool   `gorm:"default:false" json:"is_verified"`

	// Timestamp
}
