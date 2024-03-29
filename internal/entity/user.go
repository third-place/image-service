package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string     `gorm:"unique;not null"`
	Albums   []*Album
	Images   []*Image
}
