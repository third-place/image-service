package mapper

import (
	"github.com/third-place/image-service/internal/entity"
	"github.com/third-place/image-service/internal/model"
)

func GetAlbumEntityFromNewModel(album *model.NewAlbum) *entity.Album {
	return &entity.Album{
		Link:        album.Link,
		AlbumType:   string(model.UserCreated),
		Name:        album.Name,
		Description: album.Description,
		Visibility:  string(album.Visibility),
	}
}

func GetAlbumModelFromEntity(album *entity.Album) *model.Album {
	return &model.Album{
		Uuid:        album.Uuid.String(),
		Link:        album.Link,
		Name:        album.Name,
		Description: album.Description,
		Visibility:  model.Visibility(album.Visibility),
		AlbumType:   model.AlbumType(album.AlbumType),
		CreatedAt:   album.CreatedAt,
		UpdatedAt:   album.UpdatedAt,
		User:        GetUserModelFromEntity(album.User),
	}
}

func GetAlbumModelsFromEntities(albums []*entity.Album) []*model.Album {
	albumModels := make([]*model.Album, len(albums))
	for i, album := range albums {
		albumModels[i] = GetAlbumModelFromEntity(album)
	}
	return albumModels
}
