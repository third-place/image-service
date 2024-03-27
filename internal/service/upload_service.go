package service

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UploadService interface {
	UploadImage(file multipart.File, filename string, filesize int64) (key string, err error)
}

type TestUploadService struct{}

func CreateTestUploadService() UploadService {
	return &TestUploadService{}
}

func (t *TestUploadService) UploadImage(file multipart.File, filename string, filesize int64) (key string, err error) {
	return uuid.New().String(), nil
}

type LocalFSUploadService struct {
	dir string
}

func CreateLocalFSUploadService() UploadService {
	return &LocalFSUploadService{
		dir: os.Getenv("IMAGE_DIR"),
	}
}

func (l *LocalFSUploadService) UploadImage(file multipart.File, filename string, filesize int64) (key string, err error) {
	key = uuid.New().String() + filepath.Ext(filename)
	fullPath := fmt.Sprintf("%s/%s", l.dir, key)
	dst, err := os.Create(fullPath)
	if err != nil {
		log.Println("error creating file", err)
		return
	}
	_, err = io.Copy(dst, file)
	return
}

type S3UploadService struct {
	s3Client *s3.S3
	bucket   string
}

func CreateS3UploadService() *S3UploadService {
	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(err)
	}
	return &S3UploadService{
		s3.New(s),
		os.Getenv("S3_BUCKET"),
	}
}

func (u *S3UploadService) UploadImage(file multipart.File, filename string, filesize int64) (key string, err error) {
	log.Print("upload image to s3 :: ", filename)
	buffer := make([]byte, filesize)
	file.Read(buffer)
	key = uuid.New().String() + filepath.Ext(filename)
	_, err = u.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucket),
		Key:                  aws.String(key),
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
