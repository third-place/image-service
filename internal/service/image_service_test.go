package service

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/third-place/image-service/internal/test"
	"github.com/third-place/image-service/internal/util"
	"mime/multipart"
	"os"
	"testing"
)

func createFile() (multipart.File, error) {
	filePath := "sample.jpeg"
	return os.Open(filePath)
}

func TestMain(m *testing.M) {
	if os.Getenv("CI") == "" {
		_ = godotenv.Load()
	}
	os.Exit(m.Run())
}

func Test_Can_UploadImage(t *testing.T) {
	// setup
	svc := CreateTestService()
	user := util.CreateTestUser()
	svc.UpsertUser(user)

	// given
	filename := "sample.jpeg"
	file, _ := createFile()
	fstat, _ := os.Stat(filename)

	// when
	imageModel, err := svc.CreateNewProfileImage(uuid.MustParse(user.Uuid), file, filename, fstat.Size())

	// then
	test.Assert(t, imageModel != nil)
	test.Assert(t, err == nil)
}
