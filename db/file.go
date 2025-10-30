package db

import "gorm.io/gorm"

type File struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;index"`
	Name      string `gorm:"size:255;not null"`
	Hash      string `gorm:"uniqueIndex"`
	Size      int64
	Encrypt   bool   `gorm:"default:false"`
	Locked    bool   `gorm:"default:false"`
	Hidden    bool   `gorm:"default:false"`
	Pinned    bool   `gorm:"default:false"`
	ShareLink string `gorm:"size:500"`
	SourceURL string
	gorm.Model
}

type FileTag struct {
	ID     uint   `gorm:"primaryKey"`
	FileID uint   `gorm:"not null;index"`
	Tag    string `gorm:"size:100;not null"`
}

type FileMeta struct {
	ID     uint   `gorm:"primaryKey"`
	FileID uint   `gorm:"not null;index"`
	Key    string `gorm:"size:100;not null"`
	Value  string `gorm:"type:text"`
}
