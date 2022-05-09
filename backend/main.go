package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"github.com/keiliving/ramenRecorder/backend/api"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load("../.env")
	message := os.Getenv("SAMPLE_MESSAGE")
	log.Println(message)
	http.Handle("/", http.FileServer(http.Dir("../frontend/build")))
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if  (r.Method != "POST") {
		// TODO: POST のみ受け付けるようにする。
		log.Fatal("only POST")
	}
	file, header, e := r.FormFile("file")
	if (e != nil) {
		log.Fatal(e)
	}
	defer file.Close()

	entry := &api.Entry{File: file, Name: header.Filename}
	credentialFilePath := "../key.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	user := &api.User{Client: client}
	user.Upload(entry, ctx)
}
