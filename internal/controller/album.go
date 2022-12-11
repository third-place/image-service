package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/service"
	"github.com/third-place/image-service/internal/util"
	"log"
	"net/http"
)

// CreateNewAlbumV1 - create a new album
func CreateNewAlbumV1(w http.ResponseWriter, r *http.Request) {
	newAlbum := model.DecodeRequestToNewAlbum(r)
	session, err := util.GetSession(r.Header.Get("x-session-token"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	album := service.CreateDefaultAlbumService().CreateAlbum(uuid.MustParse(session.User.Uuid), newAlbum)
	data, err := json.Marshal(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
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
