package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mojeico/gqlgen-golang/graph/generated"
	"github.com/mojeico/gqlgen-golang/graph/model"
	"github.com/mojeico/gqlgen-golang/internal/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user, err := r.UserService.GetUserByID(obj.UserID)
	return user, err
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupsService.CreateMeetup(input)
	return meetup, err
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input *model.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupsService.UpdateMeetup(id, input)
	return meetup, err
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (*bool, error) {
	result, err := r.MeetupsService.DeleteMeetup(id)
	return result, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user, err := r.UserService.CreateUser(input)
	return user, err
}

func (r *queryResolver) GetAllMeetups(ctx context.Context) ([]*models.Meetup, error) {
	meetup, err := r.MeetupsService.GetAllMeetups()
	return meetup, err
}

func (r *queryResolver) GetMeetupByID(ctx context.Context, id string) (*models.Meetup, error) {
	meetup, err := r.MeetupsService.GetMeetupByID(id)
	return meetup, err
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	user, err := r.UserService.GetAllUsers()
	return user, err
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := r.UserService.GetUserByID(id)
	return user, err
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	panic(fmt.Errorf("not implement ed"))
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
