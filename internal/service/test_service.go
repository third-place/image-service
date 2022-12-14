package service

import (
	"github.com/google/uuid"
	"github.com/third-place/image-service/internal/model"
	"mime/multipart"
)

type TestService struct {
	userService  *UserService
	imageService *ImageService
}

func CreateTestService() *TestService {
	return &TestService{
		CreateTestUserService(),
		CreateTestImageService(),
	}
}

func (t *TestService) UpsertUser(userModel *model.User) {
	t.userService.UpsertUser(userModel)
}

func (t *TestService) CreateNewProfileImage(userUuid uuid.UUID, file multipart.File, filename string, filesize int64) (imageModel *model.Image, err error) {
	return t.imageService.CreateNewProfileImage(userUuid, file, filename, filesize)
}
