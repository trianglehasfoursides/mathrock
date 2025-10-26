package storage

import (
	"os"
	"path/filepath"
)

type filesystem struct {
	path string
}

func (f *filesystem) Setup() (err error) {}

func (f *filesystem) Upload(file *os.File) (err error) {
	defer file.Close()
	f.path = filepath.Join(f.path)
}
