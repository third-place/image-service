package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/third-place/image-service/internal/service"
)

func main() {
	svc := service.CreateConsumerService()
	svc.InitializeAndRunLoop()
}
