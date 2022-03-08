package services

import (
	"go-todoApp/entities"
	"go-todoApp/libs"
	"go-todoApp/repository"

	"gorm.io/gorm"
)

type UserService struct {
	logger     libs.Logger
	repository repository.UserRepository
}

func NewUserService(logger libs.Logger, repository repository.UserRepository) UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s UserService) GetOneUser(username string) (user entities.User, err error) {
	return user, s.repository.Where("username = ?", username).First(&user).Error
}

func (s UserService) GetAllUser(username string) (users []entities.User, err error) {
	return users, s.repository.Find(&users).Error
}

func (s UserService) CreateUser(user entities.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) UpdateUser(user entities.User) error {
	return s.repository.Save(&user).Error
}
