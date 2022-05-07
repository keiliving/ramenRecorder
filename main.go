package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/keiliving/ramenRecorder/api"
	"google.golang.org/api/option"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
	credentialFilePath := "./key.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	user := &api.User{Client: client}
	user.Upload(ctx)
}
