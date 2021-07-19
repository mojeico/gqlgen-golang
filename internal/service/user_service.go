package service

import (
	"github.com/mojeico/gqlgen-golang/graph/model"
	"github.com/mojeico/gqlgen-golang/internal/models"
	"github.com/mojeico/gqlgen-golang/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]*models.User, error)
	CreateUser(meetup model.NewUser) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (service userService) GetAllUsers() ([]*models.User, error) {
	return service.repo.GetAllUsers()
}

func (service userService) CreateUser(user model.NewUser) (*models.User, error) {
	return service.repo.CreateUser(user)
}

func (service userService) GetUserByID(id string) (*models.User, error) {
	return service.repo.GetUserByID(id)
}
