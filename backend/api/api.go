package api

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type User struct {
	Client *storage.Client
}

type Entry struct {
	File multipart.File
	Name string
}
var Use = User{}

func (user *User) Upload(e *Entry,ctx context.Context) {
	bucketName := os.Getenv("BUCKET_NAME")
	writer := user.Client.Bucket(bucketName).Object(e.Name).NewWriter(ctx)
	if _, err := io.Copy(writer, e.File); err != nil {
		log.Println(err)
	}

	if err := writer.Close(); err != nil {
		log.Println(err)
	}
	log.Println("upload succeeded!")
}

func (user *User) Get(name string, ctx context.Context) []byte{
	bucketName := os.Getenv("BUCKET_NAME")
	obj := user.Client.Bucket(bucketName).Object(name)
	rc, err := obj.NewReader(ctx)
	if err != nil {
		log.Println(err)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, rc); err != nil {
		log.Println(err)
	}
	ret := buf.Bytes()
	return ret
}

func (user *User) Delete(name string, ctx context.Context) {
	bucketName := os.Getenv("BUCKET_NAME")
	obj := user.Client.Bucket(bucketName).Object(name)
	if err := obj.Delete(ctx); err != nil {
		log.Println(err)
	}
}


// Fixme: use listing object 
// https://pkg.go.dev/cloud.google.com/go/storage#hdr-Listing_objects
func (user *User) Ls(ctx context.Context) []string{
	bucketName := os.Getenv("BUCKET_NAME")
	backet := user.Client.Bucket(bucketName)
	it := backet.Objects(ctx, nil)

	var objNames []string
	for {
		objAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
		}
		objNames = append(objNames, objAttrs.Name)
	}
	return objNames
}
