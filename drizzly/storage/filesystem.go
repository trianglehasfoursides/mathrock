package storage

import (
	"io"
	"os"
	"path/filepath"
)

type filesystem struct {
	path string
}

func (f *filesystem) Setup() (err error) {}

func (f *filesystem) Upload(id string, file *os.File) (err error) {
	defer file.Close()

	// uuid
	f.path = filepath.Join(f.path, id)

	newfile, err := os.Create(f.path)
	switch err {
	}

	io.Copy(file, newfile)
}

func (f *filesystem) Get(name string) (err error) {}
