package service

import (
	"github.com/mojeico/gqlgen-golang/internal/model"
	"github.com/mojeico/gqlgen-golang/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]*model.User, error)
	CreateUser(meetup model.NewUser) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUserName(userName string) (*model.User, error)
	RegistrationUser(user model.User) (string, error)
}

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (service userService) GetAllUsers() ([]*model.User, error) {
	return service.repo.GetAllUsers()
}

func (service userService) CreateUser(user model.NewUser) (*model.User, error) {
	return service.repo.CreateUser(user)
}

func (service userService) GetUserByID(id string) (*model.User, error) {
	return service.repo.GetUserByID(id)
}

func (service userService) GetUserByEmail(email string) (*model.User, error) {
	return service.repo.GetUserByEmail(email)
}

func (service userService) GetUserByUserName(userName string) (*model.User, error) {
	return service.repo.GetUserByUserName(userName)
}

func (service userService) RegistrationUser(user model.User) (string, error) {
	return service.repo.RegistrationUser(user)
}
