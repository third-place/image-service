package service

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UploadService struct {
	s3Client *s3.S3
	bucket   string
}

func CreateUploadService(s3Client *s3.S3, bucket string) *UploadService {
	return &UploadService{
		s3Client,
		bucket,
	}
}

func CreateDefaultUploadService() *UploadService {
	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(err)
	}
	return CreateUploadService(s3.New(s), os.Getenv("S3_BUCKET"))
}

func (u *UploadService) UploadImage(file multipart.File, filename string, filesize int64) (s3Key string, err error) {
	log.Print("upload image to s3 :: ", filename)
	buffer := make([]byte, filesize)
	file.Read(buffer)
	s3Key = uuid.New().String() + filepath.Ext(filename)
	_, err = u.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucket),
		Key:                  aws.String(s3Key),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(filesize),
		ContentType:          aws.String(getContentType(filename)),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Print(err)
	}
	return
}

func getContentType(file string) string {
	ext := filepath.Ext(file)
	if ext == ".png" {
		return "image/png"
	} else if ext == ".jpg" || ext == ".jpeg" {
		return "image/jpeg"
	} else {
		return "application/octet-stream"
	}
}
