package models

import (
	"time"
)

type Video struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"not null"`
	Title       string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	FileURL     string         `gorm:"type:varchar(255);not null"`
	ThumbnailURL string        `gorm:"type:varchar(255)"`
	UploadDate  time.Time      `gorm:"autoCreateTime"`
	Duration    int            `gorm:"not null"`  // Duration in seconds
	ViewsCount  int            `gorm:"default:0"`
	LikesCount  int            `gorm:"default:0"`
	Comments    []Comment      `gorm:"foreignKey:VideoID"`
	Bookmarks   []Bookmark     `gorm:"foreignKey:VideoID"`
}