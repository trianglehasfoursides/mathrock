package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"go.etcd.io/bbolt"
)

type filesystem struct {
}

func (f *filesystem) Setup() (err error) {}

func (f *filesystem) Upload(filehead *multipart.FileHeader) (hashstr string, err error) {
	file, err := filehead.Open()
	if err != nil {
		return
	}
	defer file.Close()

	fname := filehead.Filename
	temp := filepath.Join(xdg.DataHome, "uploads", fname+".tmp")

	if err = os.MkdirAll(filepath.Dir(temp), 0755); err != nil {
		return
	}

	newfile, err := os.Create(temp)
	if err != nil {
		return
	}
	defer newfile.Close()

	hash := sha256.New()
	writer := io.MultiWriter(newfile, hash)

	if _, err = file.Seek(0, 0); err != nil {
		return
	}

	bytescop, err := io.Copy(writer, file)
	if err != nil || bytescop == 0 {
		os.Remove(temp)
		return
	}

	hashsum := hash.Sum(nil)
	hashstr = hex.EncodeToString(hashsum)
	finalpath := filepath.Join(xdg.DataHome, "files", hashstr)

	if err = metatable.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("hashes"))
		if err != nil {
			return err
		}

		if val := b.Get([]byte(hashstr)); val != nil {
			return nil
		}

		return b.Put([]byte(hashstr), []byte(fname))
	}); err != nil {
		os.Remove(temp)
		return
	}

	if err = os.Rename(temp, finalpath); err != nil {
		// fallback copy
		if err2 := f.copy(temp, finalpath); err2 != nil {
			os.Remove(temp)
			return
		}
		os.Remove(temp)
	}

	return
}

func (f *filesystem) Get(name string) (err error) {}

func (f *filesystem) copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	return out.Sync()
}
