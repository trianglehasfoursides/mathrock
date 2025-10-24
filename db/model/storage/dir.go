package sg

import "gorm.io/gorm"

type Directory struct {
	gorm.Model
	Name   string
	UserID uint
	Files  []File
}

func (Directory) TableName() string {
	return "storage.directories"
}
