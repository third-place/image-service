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
	imageModel := &model.Image{
		Uuid:        image.Uuid.String(),
		Link:        image.Link,
		Key:         image.Key,
		CreatedAt:   image.CreatedAt,
		User:        userModel,
		ContentType: image.ContentType,
	}
	if image.Album != nil {
		imageModel.Album = model.Album{
			Uuid:       image.Album.Uuid.String(),
			Visibility: model.Visibility(image.Album.Visibility),
		}
	}
	return imageModel
}

func GetImageModelsFromEntities(images []*entity.Image) []*model.Image {
	imageModels := make([]*model.Image, len(images))
	for i, image := range images {
		imageModels[i] = GetImageModelFromEntity(image)
	}
	return imageModels
}
