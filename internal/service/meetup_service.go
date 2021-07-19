package service

import (
	"github.com/mojeico/gqlgen-golang/graph/model"
	"github.com/mojeico/gqlgen-golang/internal/models"
	"github.com/mojeico/gqlgen-golang/internal/repository"
)

type MeetupsService interface {
	GetAllMeetups() ([]*models.Meetup, error)
	CreateMeetup(meetup model.NewMeetup) (*models.Meetup, error)
	GetMeetupByID(id string) *models.Meetup
}

type meetupsService struct {
	repo repository.MeetupsRepo
}

func NewMeetupsRepo(repo repository.MeetupsRepo) MeetupsService {
	return &meetupsService{
		repo: repo,
	}
}

func (service meetupsService) GetAllMeetups() ([]*models.Meetup, error) {
	meetups, err := service.repo.GetAllMeetups()
	return meetups, err
}

func (service meetupsService) CreateMeetup(meetup model.NewMeetup) (*models.Meetup, error) {
	return service.repo.CreateMeetup(meetup)
}

func (service meetupsService) GetMeetupByID(id string) *models.Meetup {
	return service.repo.GetMeetupByID(id)

}
