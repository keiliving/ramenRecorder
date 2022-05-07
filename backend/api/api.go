package api

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type User struct {
	Client *storage.Client
}

func (user *User) Upload(ctx context.Context) {
	// mock file
	f, err := os.Open("kasuga_yamagen.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	// uuid にするとか
	objectPath := "test.jpeg" // e.g. foo/var/sample.txt
	writer := user.Client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	if _, err := io.Copy(writer, f); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}
	log.Println("upload succeeded!")
}

func (user *User) Get(ctx context.Context) {
	bucketName := os.Getenv("BUCKET_NAME")
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt
	obj := user.Client.Bucket(bucketName).Object(objectPath)
	objName := obj.ObjectName()
	rc, err := obj.NewReader(ctx)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(os.Stdout, rc); err != nil {
		panic(err)
	}
	log.Println("get succeeded!" + objName)
}

func (user *User) Delete(ctx context.Context) {
	bucketName := os.Getenv("BUCKET_NAME")
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt
	obj := user.Client.Bucket(bucketName).Object(objectPath)
	objName := obj.ObjectName()
	if err := obj.Delete(ctx); err != nil {
		panic(err)
	}

	log.Println("delete succeeded!" + objName)
}
