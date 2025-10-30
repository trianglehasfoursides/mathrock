// db/todo.go
package db

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	Title       string `gorm:"size:200;not null" validate:"required,min=1,max=200"`
	Description string `gorm:"type:text" validate:"max=1000"`
	Completed   bool   `gorm:"default:false"`
	DueDate     *time.Time
	Priority    string `gorm:"size:20;default:'medium'" validate:"oneof=low medium high"`
	gorm.Model
}
