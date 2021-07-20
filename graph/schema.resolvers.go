package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/mojeico/gqlgen-golang/graph/generated"
	"github.com/mojeico/gqlgen-golang/internal/middleware"
	"github.com/mojeico/gqlgen-golang/internal/model"
)

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	user, err := r.UserService.GetUserByID(obj.UserID)
	return user, err
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	user, _ := r.UserService.GetUserByEmail(input.Email)

	if user != nil {
		return nil, errors.New("email already use")
	}

	user, _ = r.UserService.GetUserByUserName(input.Username)

	if user != nil {
		return nil, errors.New("username already use")
	}

	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	newUser := model.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CratedAt:  currentTime,
		UpdatedAt: currentTime,
	}

	err := newUser.HashPassword(input.Password)

	if err != nil {
		return nil, errors.New("problem with password")
	}

	id, err := r.UserService.RegistrationUser(newUser)

	if err != nil {
		return nil, errors.New("something wrong with registration")
	}

	newUser.ID = id
	token, err := newUser.GenToken()

	if err != nil {
		return nil, errors.New("something wrong with token")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      &newUser,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.UserService.GetUserByEmail(input.Email)

	if err != nil {
		return nil, ErrorBadCredential
	}

	err = user.ComparePassword(input.Password)

	if err != nil {
		return nil, ErrorBadCredential
	}

	token, err := user.GenToken()

	if err != nil {
		return nil, errors.New("something wrong with token")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromContext(ctx)

	if err != nil {
		return nil, UserUnauthenticated
	}

	newMeetup := model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return r.MeetupsService.CreateMeetup(newMeetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input *model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := r.MeetupsService.UpdateMeetup(id, input)
	return meetup, err
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (*bool, error) {
	result, err := r.MeetupsService.DeleteMeetup(id)
	return result, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.UserService.CreateUser(input)
	return user, err
}

func (r *queryResolver) GetAllMeetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*model.Meetup, error) {
	meetup, err := r.MeetupsService.GetAllMeetups(filter, *limit, *offset)
	return meetup, err
}

func (r *queryResolver) GetMeetupByID(ctx context.Context, id string) (*model.Meetup, error) {
	meetup, err := r.MeetupsService.GetMeetupByID(id)
	return meetup, err
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	user, err := r.UserService.GetAllUsers()
	return user, err
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := r.UserService.GetUserByID(id)
	return user, err
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrorBadCredential  = errors.New("email/password combination doesn't work")
	UserUnauthenticated = errors.New("user is not unauthenticated")
)
