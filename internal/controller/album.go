package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/service"
	"github.com/third-place/image-service/internal/util"
	"net/http"
)

// CreateNewAlbumV1 - create a new album
func CreateNewAlbumV1(c *gin.Context) {
	newAlbum := model.DecodeRequestToNewAlbum(c.Request)
	session, err := util.GetSession(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	album := service.CreateDefaultAlbumService().CreateAlbum(uuid.MustParse(session.User.Uuid), newAlbum)
	c.JSON(http.StatusOK, album)
}

// GetAlbumV1 - create a new album
func GetAlbumV1(c *gin.Context) {
	uuidParam, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	album, err := service.CreateDefaultAlbumService().GetAlbum(uuidParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, album)
}

// GetAlbumsForUserV1 - create a new album
func GetAlbumsForUserV1(c *gin.Context) {
	username := c.Param("username")
	albums, err := service.CreateDefaultAlbumService().GetAlbumsForUser(username)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, albums)
}
