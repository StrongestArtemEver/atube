package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(50);not null"`
	Email string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"not null"`
	Created time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
    Videos    []Video        `gorm:"foreignKey:UserID"`
    Bookmarks []Bookmark     `gorm:"foreignKey:UserID"`
    Comments  []Comment      `gorm:"foreignKey:UserID"`
}