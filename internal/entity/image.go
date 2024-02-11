package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Link     string
	Filename string
	S3Key    string
	User     *User
	UserID   uint
	Album    *Album
	AlbumID  uint
	Uuid     *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
