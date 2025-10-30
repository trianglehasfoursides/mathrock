package db

import "gorm.io/gorm"

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	Name        string `gorm:"size:100;not null;uniqueIndex"` // nama URL (sebagai short code)
	OriginalURL string `gorm:"type:text;not null"`
	Clicks      uint   `gorm:"default:0"`
	gorm.Model
}
