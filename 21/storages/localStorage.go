package storages

import (
	"errors"
	"fmt"
)

type LSInterface interface {
	WriteFile(path string, data []byte) error
	ReadFile(path string) ([]byte, error)
}

type LocalStorage struct{}

func (ls *LocalStorage) WriteFile(path string, data []byte) error {
	if path == "" || data == nil {
		return errors.New("не переданы обязательные параметры")
	}

	fmt.Println("Local storage: по адресу", path, "записано значение -", data)
	return nil
}

func (ls *LocalStorage) ReadFile(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("не передан обязательный параметр")
	}

	fmt.Println("Local storage: прочитано значение из адреса", path)
	return []byte("file content local storage"), nil
}
