package repository

import (
	"errors"
	"github.com/third-place/image-service/internal/entity"
	"github.com/third-place/image-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"log"
)

type AlbumRepository struct {
	conn *gorm.DB
}

func CreateAlbumRepository(conn *gorm.DB) *AlbumRepository {
	return &AlbumRepository{conn}
}

func (a *AlbumRepository) Create(album *entity.Album) {
	a.conn.Create(album)
}

func (a *AlbumRepository) FindOrCreateProfileAlbumForUser(user *entity.User) *entity.Album {
	return a.findOrCreateAlbumByType(
		user,
		model.ProfilePics,
		user.Username+"'s profile pictures",
		"Profile pictures for "+user.Username,
	)
}

func (a *AlbumRepository) FindOrCreateLivestreamAlbumForUser(user *entity.User) *entity.Album {
	return a.findOrCreateAlbumByType(
		user,
		model.Livestream,
		user.Username+"'s livestream",
		"Livestream for "+user.Username,
	)
}

func (a *AlbumRepository) findOrCreateAlbumByType(
	user *entity.User,
	albumType model.AlbumType,
	albumName string,
	albumDescription string) *entity.Album {
	log.Print("find or create profile album, user :: ", user.ID)
	album := &entity.Album{}
	a.conn.
		Table("albums").
		Where("user_id = ? AND album_type = ?", user.ID, albumType).
		Scan(&album)
	if album.Uuid == nil {
		album = &entity.Album{
			Link:        user.Username,
			AlbumType:   string(albumType),
			Name:        albumName,
			Description: albumDescription,
			User:        user,
			UserID:      user.ID,
			Visibility:  string(model.PUBLIC),
		}
		a.conn.Create(album)
	}
	return album
}

func (a *AlbumRepository) FindAllByUser(userEntity *entity.User) []*entity.Album {
	var albumEntities []*entity.Album
	a.conn.Preload("User").
		Table("albums").
		Where("user_id = ?", userEntity.ID).
		Find(&albumEntities)
	return albumEntities
}

func (a *AlbumRepository) FindOne(albumUuid uuid.UUID) (*entity.Album, error) {
	albumEntity := &entity.Album{}
	a.conn.Preload("User").
		Preload("Images").
		Table("albums").
		Where("uuid = ?", albumUuid).
		Find(albumEntity)
	if albumEntity.ID == 0 {
		return nil, errors.New("album not found")
	}
	// yuck -- reverse slice
	for i, j := 0, len(albumEntity.Images)-1; i < j; i, j = i+1, j-1 {
		albumEntity.Images[i], albumEntity.Images[j] = albumEntity.Images[j], albumEntity.Images[i]
	}
	return albumEntity, nil
}
