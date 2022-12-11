package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/third-place/image-service/internal/service"
	"github.com/third-place/image-service/internal/util"
	"net/http"
)

// CreateNewImageV1 - create a new image
func CreateNewImageV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumUuid := uuid.MustParse(params["uuid"])
	session, err := util.GetSession(r.Header.Get("x-session-token"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tempFile, fileHeader, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	album, err := service.CreateDefaultImageService().CreateNewImageForAlbum(
		uuid.MustParse(session.User.Uuid),
		albumUuid,
		tempFile,
		fileHeader.Filename,
		fileHeader.Size,
	)
	data, err := json.Marshal(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

// UploadNewLivestreamImageV1 - upload a new image
func UploadNewLivestreamImageV1(w http.ResponseWriter, r *http.Request) {
	session, err := util.GetSession(r.Header.Get("x-session-token"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tempFile, fileHeader, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	image, err := service.CreateDefaultImageService().CreateNewLivestreamImage(
		uuid.MustParse(session.User.Uuid),
		tempFile,
		fileHeader.Filename,
		fileHeader.Size,
	)
	data, err := json.Marshal(image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

// UploadNewProfileImageV1 - upload a new profile pic
func UploadNewProfileImageV1(w http.ResponseWriter, r *http.Request) {
	session, err := util.GetSession(r.Header.Get("x-session-token"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tempFile, fileHeader, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	image, err := service.CreateDefaultImageService().
		CreateNewProfileImage(uuid.MustParse(session.User.Uuid), tempFile, fileHeader.Filename, fileHeader.Size)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
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
