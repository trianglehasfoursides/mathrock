package storage

import (
	"errors"
	"os"
)

func Setup() (err error) {
	switch os.Getenv("") {
	case "object":
		Store = new(object)
	case "filesystem":
	case "block":
	default:
		return errors.New("")
	}

	return nil
}
