package controller

import (
	"encoding/json"
	model2 "github.com/third-place/image-service/internal/auth/model"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CreateNewAlbumV1 - create a new album
func CreateNewAlbumV1(w http.ResponseWriter, r *http.Request) {
	newAlbum := model.DecodeRequestToNewAlbum(r)
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model2.Session) (interface{}, error) {
		album := service.CreateDefaultAlbumService().CreateAlbum(uuid.MustParse(session.User.Uuid), newAlbum)
		data, err := json.Marshal(album)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}
		return data, err
	})
}

// GetAlbumV1 - create a new album
func GetAlbumV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := params["uuid"]
	album, err := service.CreateDefaultAlbumService().GetAlbum(uuid.MustParse(uuidParam))
	if err != nil {
		log.Print("error received from GetAlbumV1 :: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(album)
	_, _ = w.Write(data)
}

// GetAlbumsForUserV1 - create a new album
func GetAlbumsForUserV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	albums, err := service.CreateDefaultAlbumService().GetAlbumsForUser(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(albums)
	_, _ = w.Write(data)
}
