package storage

import (
	"os"
)

var Box Storage

type Storage interface {
	Upload(file *os.File) (string, error)
	Get(hash string) error
	Delete(hash string) error
}
