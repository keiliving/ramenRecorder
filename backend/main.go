package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	godotenv.Load("../.env")
	message := os.Getenv("SAMPLE_MESSAGE")
	log.Println(message)
	http.Handle("/", http.FileServer(http.Dir("../frontend/build")))
	http.ListenAndServe(":8080", nil)
	// credentialFilePath := "./key.json"
	// ctx := context.Background()
	// client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// user := &api.User{Client: client}
	// user.Upload(ctx)
}
