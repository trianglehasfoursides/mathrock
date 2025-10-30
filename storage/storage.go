package storage

import (
	"io"
	"mime/multipart"
)

type Storage interface {
	Upload(file *multipart.FileHeader) (string, error)
	Get(hash string) (io.ReadCloser, error)
	Delete(hash string) error
}
