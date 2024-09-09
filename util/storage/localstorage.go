package storageutils

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type LocalStorage struct {
}

func (ls *LocalStorage) StoreFile(pathToFile, filename, ext string, data bytes.Buffer) (string, error) {
	outputPath := fmt.Sprintf("tmp/%s/%s.%s", pathToFile, filename, ext)

	err := os.MkdirAll("tmp/"+pathToFile, 0777)
	if err != nil {
		return "", err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err = out.Write(data.Bytes()); err != nil {
		return "", err
	}
	return outputPath, nil
}

func (ls *LocalStorage) GetFile(filepath string) (bytes.Buffer, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return bytes.Buffer{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return *bytes.NewBuffer(data), nil
}
