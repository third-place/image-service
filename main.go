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
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"github.com/third-place/image-service/internal"
	"github.com/third-place/image-service/internal/middleware"
	"github.com/third-place/image-service/internal/service"
	"log"
	"net/http"
	"os"
	"strconv"
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

func getServicePort() int {
	servicePort, err := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return servicePort
}

func serveHttp() {
	router := internal.NewRouter()
	handler := cors.AllowAll().Handler(router)
	port := getServicePort()
	router.Static("/assets", os.Getenv("IMAGE_DIR"))
	log.Printf("listening on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port),
		middleware.FileSizeLimit(handler)))
}
