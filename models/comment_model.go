package models

import(
	"time"
)

type Comment struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"`
	VideoID   uint           `gorm:"not null"`
	Content   string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
}