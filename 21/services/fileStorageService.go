package services

import (
	"examplestorage/storages"
)

type FileStorage interface {
	Load(path string, data []byte) error
	Save(path string) ([]byte, error)
}

type S3Adapter struct {
	S3 *storages.AWS_S3
}

func (a *S3Adapter) Load(path string, data []byte) error {
	return a.S3.UploadFile(a.S3.Bucket, path, data)
}

func (a *S3Adapter) Save(path string) ([]byte, error) {
	return a.S3.DownloadFile(a.S3.Bucket, path)
}

type LSAdapter struct {
	LS *storages.LocalStorage
}

func (a *LSAdapter) Load(path string, data []byte) error {
	return a.LS.WriteFile(path, data)
}

func (a *LSAdapter) Save(path string) ([]byte, error) {
	return a.LS.ReadFile(path)
}