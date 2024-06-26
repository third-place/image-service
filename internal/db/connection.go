package db

import (
	"fmt"
	"github.com/third-place/image-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var dbConn *gorm.DB

func CreateDefaultConnection() *gorm.DB {
	return CreateConnection(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"))
}

func CreateConnection(host string, port string, dbname string, user string, password string) *gorm.DB {
	if dbConn == nil {
		db, err := gorm.Open(
			postgres.Open(
				fmt.Sprintf(
					"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
					host,
					port,
					dbname,
					user,
					password,
				),
			),
			&gorm.Config{},
		)

		if err != nil {
			log.Fatal(err)
		}

		dbConn = db
		sqlConnection, err := dbConn.DB()

		if err != nil {
			log.Fatal(err)
		}

		_, err = sqlConnection.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;")

		if err != nil {
			log.Fatal(err)
		}

		err = db.AutoMigrate(
			&entity.User{},
			&entity.Image{},
			&entity.Album{},
		)

		if err != nil {
			log.Fatal(err)
		}

		sqlConnection.SetMaxOpenConns(20)
		sqlConnection.SetMaxIdleConns(5)
		sqlConnection.SetConnMaxLifetime(time.Hour)
	}
	return dbConn
}
