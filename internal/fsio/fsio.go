// Package fsio wraps basic filesystem IO so that I don't have to deal with errors in `cmd/sitegen`.
package fsio

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

func File(path string) *os.File {
	f, err := os.Create(path)

	if err != nil {
		log.Panicln(err)
	}

	return f
}

func Directory(path string, perm os.FileMode) {
	if err := os.Mkdir(path, perm); err != nil {
		log.Panicln(err)
	}
}

func Delete(path string) {
	if err := os.RemoveAll(path); err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Panicln(err)
	}
}
