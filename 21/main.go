package main

import (
	"examplestorage/services"
	"examplestorage/storages"
	"fmt"
)

func main() {
	s3 := &storages.AWS_S3{Bucket: "my-bucket"}
	ls := &storages.LocalStorage{}

	s3adapter := services.S3Adapter{S3: s3}
	lsadapter := services.LSAdapter{LS: ls}

	allStorages := []services.FileStorage{&s3adapter, &lsadapter}
	data := []byte("Hello world")

	for _, separateStorage := range allStorages {
		separateStorage.Load("example.txt", data)
		separateStorage.Save("example.txt")
		fmt.Println()
	}
}