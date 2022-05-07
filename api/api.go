package api

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type Client struct {
	client *storage.Client
	ctx    context.Context
}

func (rc *Client) upload(ramenClient Client) {
	// mock file
	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	// オブジェクトのReaderを作成
	bucketName := "ramen-recorder"           // e.g. example-bucket
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt
	writer := rc.client.Bucket(bucketName).Object(objectPath).NewWriter(rc.ctx)
	if _, err := io.Copy(writer, f); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}
}
