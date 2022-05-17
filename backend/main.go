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

var credentialFilePath string

func main() {
	godotenv.Load("../.env")
	log.Println("Waiting Requests ....")
	configureEnv() 
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/home").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "../frontend/build/index.html")
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/build/static/"))))
	r.HandleFunc("/upload", upload).Methods("POST")
	r.HandleFunc("/images", getNames).Methods("GET")
	r.HandleFunc("/image", get).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func configureEnv() {
	credentialFilePath	= os.Getenv("CREDENTIAL_FILEPATH")
}

func upload(w http.ResponseWriter, r *http.Request){
	file, header, err := r.FormFile("file")
	if (err != nil) {
		log.Println(err)
	}
	defer file.Close()

	entry := &api.Entry{File: file, Name: header.Filename}
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Println(err)
	}

	user := &api.User{Client: client}
	user.Upload(entry, ctx)
}

func get(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Println(err)
	}

	objectName := r.URL.Query().Get("name")

	user := &api.User{Client: client}
	f := user.Get(objectName, ctx)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(f)
}

func getNames(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Println(err)
	}

	user := &api.User{Client: client}
	f := user.GetNames(ctx)
	body, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
