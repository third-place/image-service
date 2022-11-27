package main

import (
	"github.com/third-place/image-service/internal/db"
	"github.com/third-place/image-service/internal/entity"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	conn := db.CreateDefaultConnection()
	conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;")
	conn.AutoMigrate(
		&entity.User{},
		&entity.Album{},
		&entity.Image{})
	conn.Model(&entity.Album{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	conn.Model(&entity.Image{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	conn.Model(&entity.Image{}).
		AddForeignKey("album_id", "albums(id)", "RESTRICT", "RESTRICT")
}