package storageutils

import (
	"bytes"
)

// create interface to easier to migreate using strategy pattern
type Storage interface {
	StoreFile(pathToFile, filename, ext string, data bytes.Buffer) (string, error)
	GetFile(filepath string) (bytes.Buffer, error)
}

func GetStorageAccessor(storageType string) Storage {
	switch storageType {
	default:
		return &LocalStorage{}
	}
}
