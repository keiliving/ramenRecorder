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
	log.Println("Waiting Requests ....")
	http.Handle("/", http.FileServer(http.Dir("../frontend/build")))
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/images", getNames)
	http.ListenAndServe(":8080", nil)
}

func upload(w http.ResponseWriter, r *http.Request){
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
	// TODO: config
	credentialFilePath := os.Getenv("CREDENTIAL_FILEPATH")
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	user := &api.User{Client: client}
	user.Upload(entry, ctx)
}

func get(w http.ResponseWriter, r *http.Request){
	credentialFilePath := os.Getenv("CREDENTIAL_FILEPATH")
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// 後で直す
	objectName := r.Body.name

	user := &api.User{Client: client}
	f := user.Get(objectName, ctx)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(f)
}

func getNames(w http.ResponseWriter, r *http.Request){
	credentialFilePath := os.Getenv("CREDENTIAL_FILEPATH")
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	user := &api.User{Client: client}
	f := user.GetNames(ctx)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(f)
}
