/*
 * Otto image service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"github.com/third-place/image-service/internal"
	"github.com/third-place/image-service/internal/middleware"
	"github.com/third-place/image-service/internal/service"
	"log"
	"net/http"
)

func main() {
	go readKafka()
	serveHttp()
}

func readKafka() {
	log.Print("connecting to kafka")
	svc := service.CreateConsumerService()
	svc.InitializeAndRunLoop()
	log.Print("exit kafka loop")
}

func serveHttp() {
	router := internal.NewRouter()
	handler := cors.AllowAll().Handler(router)
	log.Print("listening on 8082")
	log.Fatal(http.ListenAndServe(":8082",
		middleware.FileSizeLimit(middleware.ContentTypeMiddleware(handler))))
}
