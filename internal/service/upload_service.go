package service

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UploadService interface {
	UploadImage(file multipart.File, filename string, filesize int64) (key string, contentType *mimetype.MIME, err error)
}

type TestUploadService struct{}

func CreateTestUploadService() UploadService {
	return &TestUploadService{}
}

func (t *TestUploadService) UploadImage(file multipart.File, filename string, filesize int64) (key string, contentType *mimetype.MIME, err error) {
	return uuid.New().String(), &mimetype.MIME{}, nil
}

type LocalFSUploadService struct {
	dir string
}

func CreateLocalFSUploadService() UploadService {
	return &LocalFSUploadService{
		dir: os.Getenv("IMAGE_DIR"),
	}
}

func (l *LocalFSUploadService) UploadImage(file multipart.File, filename string, filesize int64) (key string, contentType *mimetype.MIME, err error) {
	key = uuid.New().String() + filepath.Ext(filename)
	fullPath := fmt.Sprintf("%s/%s", l.dir, key)
	dst, err := os.Create(fullPath)
	if err != nil {
		log.Println("error creating file", err)
		return
	}
	contentType, err = mimetype.DetectReader(file)
	if err != nil {
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

func (u *S3UploadService) UploadImage(file multipart.File, filename string, filesize int64) (key string, contentType *mimetype.MIME, err error) {
	log.Print("upload image to s3 :: ", filename)
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		return
	}
	contentType, err = mimetype.DetectReader(file)
	if err != nil {
		return
	}
	key = uuid.New().String() + filepath.Ext(filename)
	_, err = u.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucket),
		Key:                  aws.String(key),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(filesize),
		ContentType:          aws.String(contentType.String()),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Print(err)
	}
	return
}
