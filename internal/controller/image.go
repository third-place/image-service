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
	keyParam := c.Param("key")
	imageModel, err := service.CreateImageService().GetImageFromKey(keyParam)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if imageModel.Album.Visibility != model.PUBLIC {
		session, err := util.GetSession(c)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if session.User.Uuid != imageModel.User.Uuid {
			c.Status(http.StatusForbidden)
			return
		}
	}
	file := fmt.Sprintf("%s/%s", os.Getenv("IMAGE_DIR"), imageModel.Key)
	buf, err := os.ReadFile(file)
	if err != nil {
		log.Print(err)
		return
	}
	c.Status(http.StatusOK)
	c.Writer.Header().Set("Content-Type", imageModel.ContentType)
	_, err = c.Writer.Write(buf)
	if err != nil {
		log.Print(err)
		return
	}
}

// GetImageV1 - get an image
func GetImageV1(c *gin.Context) {
	uuidParam, err := uuid.Parse(c.Param("uuid"))
	imageModel, err := service.CreateImageService().GetImage(uuidParam)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
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
