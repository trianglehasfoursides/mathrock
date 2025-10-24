package sg

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name        string
	Size        int64
	Chunks      []Chunk
	Checksum    string
	MimeType    string
	UserID      uint
	DirectoryID uint
}

func (File) TableName() string {
	return "storage.files"
}
