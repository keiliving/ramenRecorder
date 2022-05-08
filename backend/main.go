package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"github.com/keiliving/ramenRecorder/backend/api"
	"google.golang.org/api/option"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func upload(w http.ResponseWriter, r *http.Request) {
	if  (r.Method != "POST") {
		// TODO: POST のみ受け付けるようにする。
		log.Fatal("only POST")
	}
	file, _, e := r.FormFile("test");
	if (e != nil) {
		log.Fatal(e)
	}
	defer file.Close()

	credentialFilePath := "./key.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	user := &api.User{Client: client}
	user.Upload(file, ctx)
}

func main() {
	godotenv.Load("../.env")
	message := os.Getenv("SAMPLE_MESSAGE")
	log.Println(message)
	http.Handle("/", http.FileServer(http.Dir("../frontend/build")))
	http.HandleFunc("/health", handler)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
