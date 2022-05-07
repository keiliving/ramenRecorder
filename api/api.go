package api

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type a interface {
	upload()
}

type User struct {
	Client *storage.Client
}

func (user *User) Upload(ctx context.Context) {
	// mock file
	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// オブジェクトのReaderを作成
	bucketName := "ramen-recorder"           // e.g. example-bucket
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt
	writer := user.Client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	if _, err := io.Copy(writer, f); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}
	log.Println("upload succeeded!")
}
