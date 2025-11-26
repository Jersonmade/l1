package storages

import (
	"errors"
	"fmt"
)

type S3Interface interface {
	UploadFile(bucket, path string, data []byte) error
	DownloadFile(bucket, path string) ([]byte, error)
}

type AWS_S3 struct {
	Bucket string
}

func (s *AWS_S3) UploadFile(bucket, path string, data []byte) error {
	if bucket == "" || path == "" || data == nil {
		return errors.New("не переданы обязательные параметры")
	}

	fmt.Println("AWS S3: По адресу", path, "в корзину", bucket, "записано -", data)

	return nil
}

func (s *AWS_S3) DownloadFile(bucket, path string) ([]byte, error) {
	if bucket == "" || path == "" {
		return nil, errors.New("не переданы обязательные параметры")
	}

	fmt.Println("AWS S3: Скачано значение по адресу", path, "из корзины", bucket)

	return []byte("file content S3"), nil
}