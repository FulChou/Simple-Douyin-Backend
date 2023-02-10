package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var minioClient *minio.Client

func Init() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "admin"
	secretAccessKey := "12345678"

	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println("minio connect error", err)
		log.Fatalln("minio connect error", err)
	}
	log.Printf("%#v\n", minioClient) // minioClient is now set up
	fmt.Println("minioClient is now set up")
	minioClient = client
	bucketName := "dousheng"

	err = CreateBucket(bucketName)
	if err != nil {
		return
	}
}
