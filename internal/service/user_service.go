package service

import (
	"github.com/google/uuid"
	"github.com/third-place/image-service/internal/db"
	"github.com/third-place/image-service/internal/mapper"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/repository"
	"github.com/third-place/image-service/internal/util"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func CreateUserService() *UserService {
	conn := db.CreateDefaultConnection()
	return &UserService{
		repository.CreateUserRepository(conn),
	}
}

func CreateTestUserService() *UserService {
	conn := util.SetupTestDatabase()
	return &UserService{
		repository.CreateUserRepository(conn),
	}
}

func (u *UserService) UpsertUser(userModel *model.User) {
	userEntity, err := u.userRepository.FindOneByUuid(uuid.MustParse(userModel.Uuid))
	if err != nil {
		userEntity = mapper.GetUserEntityFromModel(userModel)
		u.userRepository.Create(userEntity)
	}
}
