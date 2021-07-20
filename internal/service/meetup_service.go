package service

import (
	"github.com/mojeico/gqlgen-golang/internal/model"
	"github.com/mojeico/gqlgen-golang/internal/repository"
)

type MeetupsService interface {
	GetAllMeetups(filter *model.MeetupFilter, limit int, offset int) ([]*model.Meetup, error)
	CreateMeetup(meetup model.Meetup) (*model.Meetup, error)
	GetMeetupByID(id string) (*model.Meetup, error)
	UpdateMeetup(id string, meetup *model.UpdateMeetup) (*model.Meetup, error)
	DeleteMeetup(id string) (*bool, error)
}

type meetupsService struct {
	repo repository.MeetupsRepo
}

func NewMeetupsRepo(repo repository.MeetupsRepo) MeetupsService {
	return &meetupsService{
		repo: repo,
	}
}

func (service meetupsService) DeleteMeetup(id string) (*bool, error) {
	return service.repo.DeleteMeetup(id)
}

func (service meetupsService) UpdateMeetup(id string, meetup *model.UpdateMeetup) (*model.Meetup, error) {
	return service.repo.UpdateMeetup(id, meetup)
}

func (service meetupsService) GetAllMeetups(filter *model.MeetupFilter, limit int, offset int) ([]*model.Meetup, error) {

	meetups, err := service.repo.GetAllMeetups(filter, int64(limit), int64(offset))
	return meetups, err
}

func (service meetupsService) CreateMeetup(meetup model.Meetup) (*model.Meetup, error) {
	return service.repo.CreateMeetup(meetup)
}

func (service meetupsService) GetMeetupByID(id string) (*model.Meetup, error) {
	return service.repo.GetMeetupByID(id)

}
