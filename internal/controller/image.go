package controller

import (
	"encoding/json"
	"github.com/third-place/image-service/internal/auth/model"
	"github.com/third-place/image-service/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateNewImageV1 - create a new image
func CreateNewImageV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumUuid := uuid.MustParse(params["uuid"])
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model.Session) (interface{}, error) {
		tempFile, fileHeader, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		return service.CreateDefaultImageService().CreateNewImageForAlbum(
			uuid.MustParse(session.User.Uuid),
			albumUuid,
			tempFile,
			fileHeader.Filename,
			fileHeader.Size,
		)
	})
}

// UploadNewLivestreamImageV1 - upload a new image
func UploadNewLivestreamImageV1(w http.ResponseWriter, r *http.Request) {
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model.Session) (interface{}, error) {
		tempFile, fileHeader, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		return service.CreateDefaultImageService().CreateNewLivestreamImage(
			uuid.MustParse(session.User.Uuid),
			tempFile,
			fileHeader.Filename,
			fileHeader.Size,
		)
	})
}

// UploadNewProfileImageV1 - upload a new profile pic
func UploadNewProfileImageV1(w http.ResponseWriter, r *http.Request) {
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model.Session) (interface{}, error) {
		tempFile, fileHeader, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		return service.CreateDefaultImageService().
			CreateNewProfileImage(uuid.MustParse(session.User.Uuid), tempFile, fileHeader.Filename, fileHeader.Size)
	})
}

// GetImageV1 - get an image
func GetImageV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := uuid.MustParse(params["uuid"])
	imageModel, err := service.CreateDefaultImageService().GetImage(uuidParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(imageModel)
	_, _ = w.Write(data)
}

// GetImagesForAlbumV1 - get images for album
func GetImagesForAlbumV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := params["uuid"]

	imageModels := service.CreateDefaultImageService().GetAllImagesForAlbum(uuid.MustParse(uuidParam))
	data, _ := json.Marshal(imageModels)
	_, _ = w.Write(data)

}
