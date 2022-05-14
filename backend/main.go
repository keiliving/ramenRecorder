package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/keiliving/ramenRecorder/backend/api"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load("../.env")
	log.Println("Waiting Requests ....")
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/build")))).Methods("GET")
	r.HandleFunc("/upload", upload).Methods("POST")
	r.HandleFunc("/images", getNames).Methods("GET")
	r.HandleFunc("/image", get).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func upload(w http.ResponseWriter, r *http.Request){
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

	objectName := r.URL.Query().Get("name")

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
	body, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(body)
}
