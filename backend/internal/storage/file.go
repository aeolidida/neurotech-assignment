package storage

import (
	"neurotech-assignment/backend/internal/errs"
	"os"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{filePath: filePath}
}

func (fs *FileStorage) Save(data []byte) error {
	op := "FileStorage.Save"
	err := os.WriteFile(fs.filePath, data, 0644)
	if err != nil {
		return errs.WrapError(op, "error writing file", err)
	}

	return nil
}

func (fs *FileStorage) Load() ([]byte, error) {
	op := "FileStorage.Load"
	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, errs.WrapError(op, "error opening file", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, errs.WrapError(op, "error getting file info", err)
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, errs.WrapError(op, "error reading file", err)
	}

	return data, nil
}
