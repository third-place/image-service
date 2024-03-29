package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/third-place/image-service/internal/db"
	"github.com/third-place/image-service/internal/entity"
	"github.com/third-place/image-service/internal/kafka"
	"github.com/third-place/image-service/internal/mapper"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/repository"
	"github.com/third-place/image-service/internal/util"
	"log"
	"mime/multipart"
)

type ImageService struct {
	imageRepository *repository.ImageRepository
	albumRepository *repository.AlbumRepository
	userRepository  *repository.UserRepository
	uploadService   UploadService
	kafkaProducer   kafka.Producer
}

func CreateImageService() *ImageService {
	conn := db.CreateDefaultConnection()
	return &ImageService{
		repository.CreateImageRepository(conn),
		repository.CreateAlbumRepository(conn),
		repository.CreateUserRepository(conn),
		CreateLocalFSUploadService(),
		kafka.CreateProducer(),
	}
}

func CreateTestImageService() *ImageService {
	conn := util.SetupTestDatabase()
	producer, err := util.CreateTestProducer()
	if err != nil {
		log.Fatal(err)
	}
	return &ImageService{
		repository.CreateImageRepository(conn),
		repository.CreateAlbumRepository(conn),
		repository.CreateUserRepository(conn),
		CreateTestUploadService(),
		producer,
	}
}

func (i *ImageService) GetAllImagesForAlbum(albumUuid uuid.UUID) []*model.Image {
	images := i.imageRepository.FindByAlbumUuid(&albumUuid)
	return mapper.GetImageModelsFromEntities(images)
}

func (i *ImageService) CreateNewImageForAlbum(userUuid uuid.UUID, albumUuid uuid.UUID, file multipart.File, filename string, filesize int64) (imageModel *model.Image, err error) {
	user, err := i.userRepository.FindOneByUuid(userUuid)
	if user.Uuid == nil || err != nil {
		log.Print("error finding user :: ", err)
		return
	}
	album, err := i.albumRepository.FindOne(albumUuid)
	if err != nil {
		return
	}
	imageEntity := i.uploadAndCreateImageEntity(user, album, file, filename, filesize)
	imageModel = mapper.GetImageModelFromEntity(imageEntity)
	return
}

func (i *ImageService) CreateNewLivestreamImage(userUuid uuid.UUID, file multipart.File, filename string, filesize int64) (imageModel *model.Image, err error) {
	user, err := i.userRepository.FindOneByUuid(userUuid)
	if user.Uuid == nil || err != nil {
		log.Print("error finding user :: ", err)
		return
	}
	album := i.albumRepository.FindOrCreateLivestreamAlbumForUser(user)
	imageEntity := i.uploadAndCreateImageEntity(user, album, file, filename, filesize)
	imageModel = mapper.GetImageModelFromEntity(imageEntity)
	return
}

func (i *ImageService) CreateNewProfileImage(userUuid uuid.UUID, file multipart.File, filename string, filesize int64) (imageModel *model.Image, err error) {
	user, err := i.userRepository.FindOneByUuid(userUuid)
	if user.Uuid == nil || err != nil {
		log.Print("error finding user :: ", err)
		return
	}
	album := i.albumRepository.FindOrCreateProfileAlbumForUser(user)
	imageEntity := i.uploadAndCreateImageEntity(user, album, file, filename, filesize)
	imageModel = mapper.GetImageModelFromEntity(imageEntity)
	data, _ := json.Marshal(imageModel)
	log.Print("publishing image to kafka: ", string(data))
	topic := "images"
	_ = i.kafkaProducer.Produce(kafka.CreateMessage(data, topic), nil)
	return
}

func (i *ImageService) GetImage(imageUuid uuid.UUID) (*model.Image, error) {
	imageEntity := i.imageRepository.FindByUuid(&imageUuid)
	if imageEntity.ID == 0 {
		return nil, errors.New("image not found")
	}
	return mapper.GetImageModelFromEntity(imageEntity), nil
}

func (i *ImageService) uploadAndCreateImageEntity(user *entity.User, album *entity.Album, file multipart.File, filename string, filesize int64) *entity.Image {
	key, contentType, err := i.uploadService.UploadImage(file, filename, filesize)
	if err != nil {
		log.Print("error occurred in image service upload", err)
		return nil
	}
	imageEntity := i.createNewImageEntity(user, album)
	imageEntity.Key = key
	imageEntity.ContentType = contentType.String()
	i.imageRepository.Create(imageEntity)
	return imageEntity
}

func (i *ImageService) findOrCreateProfileImage(user *entity.User, album *entity.Album) (imageEntity *entity.Image) {
	imageEntity = i.imageRepository.FindByUserAndAlbum(user.Uuid, album.Uuid)
	if imageEntity.Uuid == nil {
		log.Print("profile pic not found, creating new one, user :: ", user.Uuid)
		imageEntity = &entity.Image{
			Filename: "",
			User:     user,
			UserID:   user.ID,
			Album:    album,
			AlbumID:  album.ID,
		}
		i.imageRepository.Create(imageEntity)
	}
	return
}

func (i *ImageService) createNewImageEntity(user *entity.User, album *entity.Album) *entity.Image {
	return &entity.Image{
		Filename: "",
		User:     user,
		UserID:   user.ID,
		Album:    album,
		AlbumID:  album.ID,
	}
}
