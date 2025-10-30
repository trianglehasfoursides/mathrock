// db/kanban.go
package db

import "gorm.io/gorm"

type Kanban struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	Title       string `gorm:"size:100;not null" validate:"required,min=1,max=100"`
	Description string `gorm:"type:text" validate:"max=500"`
	Status      string `gorm:"size:20;not null" validate:"required,oneof=todo doing done"`
	Position    int    `gorm:"default:0"` // untuk sorting dalam status yang sama
	Color       string `gorm:"size:20" validate:"omitempty,oneof=blue green red yellow purple gray"`
	gorm.Model
}
