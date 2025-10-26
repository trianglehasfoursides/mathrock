package storage

import "os"

var Store Storage

type Storage interface {
	Upload(file *os.File) error
	Get(name string) error
	Delete(name string) error
}
