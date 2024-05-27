package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"github.com/third-place/image-service/internal"
	"github.com/third-place/image-service/internal/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getServicePort() int {
	port, ok := os.LookupEnv("SERVICE_PORT")
	if !ok {
		port = "8082"
	}
	servicePort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}
	return servicePort
}

func main() {
	router := internal.NewRouter()
	handler := cors.AllowAll().Handler(router)
	port := getServicePort()
	router.Static("/assets", os.Getenv("IMAGE_DIR"))
	log.Printf("listening on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port),
		middleware.FileSizeLimit(handler)))
}
