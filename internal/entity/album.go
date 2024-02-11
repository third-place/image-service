package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Link        string
	AlbumType   string
	Name        string
	Description string
	User        *User
	UserID      uint
	Images      []*Image
	Visibility  string
	Uuid        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}