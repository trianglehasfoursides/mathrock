package sg

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	Name        string      `json:"name" validate:"required"`
	Files       []File      ``
	Directories []Directory ``
	UserID      uint        ``
}

func (Storage) TableName() string {
	return "storage.storages"
}
