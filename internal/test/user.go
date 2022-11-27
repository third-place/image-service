package test

import (
	"github.com/third-place/image-service/internal/entity"
	"github.com/google/uuid"
)

func CreateTestUser() *entity.User {
	userUuid := uuid.New()
	return &entity.User{
		Uuid: &userUuid,
	}
}
