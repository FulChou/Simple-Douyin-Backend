package minio

import (
	"Simple-Douyin-Backend/utils"
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/url"
	"time"
)

func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		return errors.New("bucketName invalid")
	}
	location := "guangzhou"
	ctx := context.Background()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("bucket %s already exists", bucketName)
			return nil
		} else {
			return err
		}
	} else {
		log.Printf("bucket %s create successfully", bucketName)
	}
	return nil
}

// FileUploader upload local file
func FileUploader(ctx context.Context, bucketName string, objectName string, filePath string) error {
	object, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Println("upload failed：", err)
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
	return nil
}

// GetFileUrl get shareUrl of file from minio
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()

	// get ip
	ip, err := utils.GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	endpoint := ip + ":9000" // "172.26.41.217:9000"
	accessKeyID := "admin"
	secretAccessKey := "12345678"

	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("minio connect error", err)
	}

	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := client.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		log.Printf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	// TODO: url is quite long, or need to shorten
	return presignedUrl, nil
}

func RemoveFile(bucketName string, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := minioClient.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Success")
	return nil
}
