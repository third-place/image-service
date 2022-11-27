package mapper

import (
	"github.com/third-place/image-service/internal/entity"
	"github.com/third-place/image-service/internal/model"
)

func GetImageModelFromEntity(image *entity.Image) *model.Image {
	var userModel model.User
	if image.User != nil {
		userModel = GetUserModelFromEntity(image.User)
	}
	return &model.Image{
		Uuid:      image.Uuid.String(),
		Link:      image.Link,
		S3Key:     image.S3Key,
		CreatedAt: image.CreatedAt,
		User:      userModel,
	}
}

func GetImageModelsFromEntities(images []*entity.Image) []*model.Image {
	imageModels := make([]*model.Image, len(images))
	for i, image := range images {
		imageModels[i] = GetImageModelFromEntity(image)
	}
	return imageModels
}
