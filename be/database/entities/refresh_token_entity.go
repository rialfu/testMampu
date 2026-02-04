package entities

import (
	"time"
)

type RefreshToken struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"token"`
	ExpiresAt time.Time `gorm:"type:timestamp with time zone;not null" json:"expires_at"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	// Timestamp
}
