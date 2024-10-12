package models

import (
	"time"
)

type VideoView struct {
	ID       uint      `gorm:"primaryKey"`
	UserID   *uint     // May be null for anonymous users
	VideoID  uint      `gorm:"not null"`
	ViewedAt time.Time `gorm:"autoCreateTime"`
}
