package api

import (
	"bytes"
	"context"
	"fmt"
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
	// uuid にするとか
	writer := user.Client.Bucket(bucketName).Object(e.Name).NewWriter(ctx)
	if _, err := io.Copy(writer, e.File); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}
	log.Println("upload succeeded!")
}

func (user *User) Get(name string, ctx context.Context) []byte{
	bucketName := os.Getenv("BUCKET_NAME")
	obj := user.Client.Bucket(bucketName).Object(name)
	rc, err := obj.NewReader(ctx)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, rc); err != nil {
		panic(err)
	}
	ret := buf.Bytes()
	return ret
}

func (user *User) Delete(name string, ctx context.Context) {
	bucketName := os.Getenv("BUCKET_NAME")
	obj := user.Client.Bucket(bucketName).Object(name)
	if err := obj.Delete(ctx); err != nil {
		panic(err)
	}
}

func (user *User) GetNames(ctx context.Context) []byte{
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
			panic(err)
		}
		objNames = append(objNames,objAttrs.Name)
	}
	fmt.Println(objNames)
	
	rc, err := backet.Object(objNames[0]).NewReader(ctx)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, rc); err != nil {
		panic(err)
	}
	ret := buf.Bytes()
	return ret
}
