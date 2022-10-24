package store

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	FileType Type = "file"
)

type Type string

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type fileStore struct {
	filePath string
}

func NewStore(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{filePath: fileName}
	}
	return nil
}

func (f *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		return fmt.Errorf("could not read the file in the path: %s", f.filePath)
	}
	return json.Unmarshal(file, &data)
}
func (f *fileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("could not write to the path file: %s", f.filePath)
	}
	return os.WriteFile(f.filePath, file, os.FileMode(os.O_RDWR))
}
