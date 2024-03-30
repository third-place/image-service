/*
 * Otto image service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package internal

import (
	"github.com/third-place/image-service/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"CreateNewAlbumV1",
		http.MethodPost,
		"/album",
		controller.CreateNewAlbumV1,
	},

	{
		"CreateNewImageV1",
		http.MethodPost,
		"/album/:uuid/image",
		controller.CreateNewImageV1,
	},

	{
		"GetAlbumV1",
		http.MethodGet,
		"/album/:uuid",
		controller.GetAlbumV1,
	},

	{
		"GetAlbumsForUserV1",
		http.MethodGet,
		"/albums/:username",
		controller.GetAlbumsForUserV1,
	},

	{
		"GetAssetV1",
		http.MethodGet,
		"/asset/:key",
		controller.GetAssetV1,
	},

	{
		"GetImageV1",
		http.MethodGet,
		"/image/:uuid",
		controller.GetImageV1,
	},

	{
		"GetImagesForAlbumV1",
		http.MethodGet,
		"/album/:uuid/image",
		controller.GetImagesForAlbumV1,
	},

	{
		"UploadNewLivestreamImageV1",
		http.MethodPost,
		"/album/livestream",
		controller.UploadNewLivestreamImageV1,
	},

	{
		"UploadNewProfileImageV1",
		http.MethodPost,
		"/album/profile",
		controller.UploadNewProfileImageV1,
	},
}
