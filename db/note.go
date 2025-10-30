// db/note.go
package db

import "gorm.io/gorm"

type Note struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;index"`
	Title     string `gorm:"size:200;not null"`
	Content   string `gorm:"type:text;not null"`
	WordCount int    `gorm:"default:0"`
	Pinned    bool   `gorm:"default:false"`
	gorm.Model
}
