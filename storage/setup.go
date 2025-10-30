package storage

import (
	"errors"
	"os"
)

var Box Storage

func Setup() (err error) {
	switch os.Getenv("") {
	case "object":
		Box = new(object)
	case "filesystem":
	case "block":
	default:
		return errors.New("")
	}

	return nil
}
