package mapper

import (
	"github.com/third-place/image-service/internal/entity"
	"github.com/third-place/image-service/internal/model"
	"github.com/google/uuid"
)

func GetUserEntityFromModel(user *model.User) *entity.User {
	userUuid := uuid.MustParse(user.Uuid)
	return &entity.User{
		Uuid:     &userUuid,
		Username: user.Username,
	}
}

func GetUserModelFromEntity(user *entity.User) model.User {
	return model.User{
		Uuid: user.Uuid.String(),
	}
}
