package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/service"
	"github.com/third-place/image-service/internal/util"
	"log"
	"net/http"
	"os"
)

// CreateNewImageV1 - create a new image
func CreateNewImageV1(c *gin.Context) {
	albumUuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	session, err := util.GetSession(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	tempFile, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	image, err := service.CreateImageService().CreateNewImageForAlbum(
		uuid.MustParse(session.User.Uuid),
		albumUuid,
		tempFile,
		fileHeader.Filename,
		fileHeader.Size,
	)
	c.JSON(http.StatusCreated, image)
}

// UploadNewLivestreamImageV1 - upload a new image
func UploadNewLivestreamImageV1(c *gin.Context) {
	session, err := util.GetSession(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	tempFile, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	image, err := service.CreateImageService().CreateNewLivestreamImage(
		uuid.MustParse(session.User.Uuid),
		tempFile,
		fileHeader.Filename,
		fileHeader.Size,
	)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, image)
}

// UploadNewProfileImageV1 - upload a new profile pic
func UploadNewProfileImageV1(c *gin.Context) {
	session, err := util.GetSession(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	tempFile, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	image, err := service.CreateImageService().
		CreateNewProfileImage(uuid.MustParse(session.User.Uuid), tempFile, fileHeader.Filename, fileHeader.Size)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, image)
}

// GetAssetV1 - get the image binary
func GetAssetV1(c *gin.Context) {
	log.Print("GetAssetV1")
	keyParam := c.Param("key")
	log.Print(keyParam)
	imageModel, err := service.CreateImageService().GetImageFromKey(keyParam)
	if err != nil {
		log.Print("image model not found")
		c.Status(http.StatusNotFound)
		return
	}
	if imageModel.Album.Visibility != model.PUBLIC {
		session, err := util.GetSession(c)
		if err != nil {
			log.Print("not logged in, trying to view private image")
			c.Status(http.StatusBadRequest)
			return
		}
		if session.User.Uuid != imageModel.User.Uuid {
			log.Print("not owner, trying to view private image")
			c.Status(http.StatusForbidden)
			return
		}
	}
	log.Print("serving static asset")
	r := gin.Default()
	r.Static("/asset", os.Getenv("IMAGE_DIR"))
}

// GetImageV1 - get an image
func GetImageV1(c *gin.Context) {
	log.Print("GetImageV1")
	uuidParam, err := uuid.Parse(c.Param("uuid"))
	log.Print(uuidParam.String())
	imageModel, err := service.CreateImageService().GetImage(uuidParam)
	log.Print(imageModel.Uuid)
	if err != nil {
		log.Print(fmt.Sprintf("image not found -- %s", uuidParam))
		c.Status(http.StatusNotFound)
		return
	}
	log.Print("ok!")
	c.JSON(http.StatusOK, imageModel)
}

// GetImagesForAlbumV1 - get images for album
func GetImagesForAlbumV1(c *gin.Context) {
	uuidParam, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	imageModels := service.CreateImageService().GetAllImagesForAlbum(uuidParam)
	c.JSON(http.StatusOK, imageModels)
}
