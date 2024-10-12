package models

import (
	"time"
)

type Bookmark struct {
	ID      uint      `gorm:"primaryKey"`
	UserID  uint      `gorm:"not null"`
	VideoID uint      `gorm:"not null"`
	AddedAt time.Time `gorm:"autoCreateTime"`
}
