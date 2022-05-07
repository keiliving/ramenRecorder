package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

type Client struct {
	client *storage.Client
	ctx    context.Context
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
	credentialFilePath := "./key.json"

	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// GCSオブジェクトを書き込む空のsample.txtファイルをローカルで作成
	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// オブジェクトのReaderを作成
	bucketName := "ramen-recorder"           // e.g. example-bucket
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt

	writer := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	if _, err := io.Copy(writer, f); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	log.Println("done")
}
