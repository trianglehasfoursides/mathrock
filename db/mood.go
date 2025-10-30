// db/mood.go
package db

import (
	"time"

	"gorm.io/gorm"
)

type Mood struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null;index"`
	Mood        string    `gorm:"size:20;not null" validate:"required,oneof=happy sad angry anxious excited tired relaxed productive stressed calm"`
	Intensity   int       `gorm:"default:5" validate:"min=1,max=10"` // 1-10 scale
	Description string    `gorm:"type:text" validate:"max=500"`
	Date        time.Time `gorm:"not null"` // Tanggal mood (bisa berbeda dari created_at)
	Tags        string    `gorm:"size:200"` // comma separated tags: work,family,health
	gorm.Model
}
