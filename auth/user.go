package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    string
	Name  string
	Email string
}
