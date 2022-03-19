package service

import (
	"golang.com/m/entity"
	"golang.com/m/repository"
)

type UserService interface {
	GetAllUsers(id int64) entity.User
	CreateUser(user *entity.User) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (u *userService) GetAllUsers(id int64) entity.User {
	return u.userRepository.GetAllUsers(id)
}

func (u *userService) CreateUser(user *entity.User) entity.User {
	u.userRepository.CreateUser(user)
	return *user
}
