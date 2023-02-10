package minio

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
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
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
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
