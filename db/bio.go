// db/linkinbio.go
package db

import "gorm.io/gorm"

type LinkInBio struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;uniqueIndex"` // Satu user satu page
	Username    string `gorm:"size:50;not null;uniqueIndex" validate:"required,alphanum,min=3,max=50"`
	DisplayName string `gorm:"size:100" validate:"max=100"`
	Bio         string `gorm:"type:text" validate:"max=500"`
	Avatar      string `gorm:"type:text"` // URL avatar
	Theme       string `gorm:"size:20;default:'default'" validate:"oneof=default dark light colorful minimal"`
	Active      bool   `gorm:"default:true"`
	gorm.Model
}

type LinkInBioItem struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	Title       string `gorm:"size:100;not null" validate:"required,min=1,max=100"`
	URL         string `gorm:"type:text;not null" validate:"required,url"`
	Description string `gorm:"size:200" validate:"max=200"`
	Icon        string `gorm:"size:50"` // icon name atau emoji
	Color       string `gorm:"size:20" validate:"omitempty,oneof=blue green red yellow purple pink gray black"`
	Position    int    `gorm:"default:0"`
	Active      bool   `gorm:"default:true"`
	gorm.Model
}
