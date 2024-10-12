package models

import (
	"time"
)

type Like struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	VideoID   uint      `gorm:"not null"`
	IsLike    bool      `gorm:"not null"` // true for like, false for dislike
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
