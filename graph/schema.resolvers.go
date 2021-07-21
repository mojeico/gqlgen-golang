package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/mojeico/gqlgen-golang/pkg/logger"
	"time"

	"github.com/mojeico/gqlgen-golang/graph/generated"
	"github.com/mojeico/gqlgen-golang/internal/middleware"
	"github.com/mojeico/gqlgen-golang/internal/model"
)

var (
	ErrorBadCredential  = errors.New("email/password combination doesn't work")
	UserUnauthenticated = errors.New("user is not unauthenticated")
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	err := input.IsValidate()

	if err != nil {
		logMessage := fmt.Sprintf("Email %s already registred", input.Email)
		logger.LoginRegisterLoggerWarn("Register", "schema.resolvers.go", logMessage)
		return nil, errors.New("problem with inserting data")
	}

	user, _ := r.UserService.GetUserByEmail(input.Email)

	if user != nil {
		logMessage := fmt.Sprintf("Email %s already registred", input.Email)
		logger.LoginRegisterLoggerWarn("Register", "schema.resolvers.go", logMessage)
		return nil, errors.New("email already use")
	}

	user, _ = r.UserService.GetUserByUserName(input.Username)

	if user != nil {
		logMessage := fmt.Sprintf("Username %s already registred", input.Username)
		logger.LoginRegisterLoggerWarn("Register", "schema.resolvers.go", logMessage)
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

	err = newUser.HashPassword(input.Password)

	if err != nil {
		logMessage := fmt.Sprintf("Problem with password - %s ", err.Error())
		logger.LoginRegisterLoggerError("Register", "schema.resolvers.go", logMessage)
		return nil, errors.New("problem with password")
	}

	id, err := r.UserService.RegistrationUser(newUser)

	if err != nil {
		logMessage := fmt.Sprintf("Something wrong with registration - %s ", err.Error())
		logger.LoginRegisterLoggerError("Register", "schema.resolvers.go", logMessage)
		return nil, errors.New("something wrong with registration")
	}

	newUser.ID = id
	token, err := newUser.GenToken()

	if err != nil {
		logMessage := fmt.Sprintf("Something wrong with token - %s ", err.Error())
		logger.LoginRegisterLoggerError("Register", "schema.resolvers.go", logMessage)
		return nil, errors.New("something wrong with token")
	}

	logMessage := fmt.Sprintf("User was registred successfully with username - %s and email - %s ", input.Username, input.Email)
	logger.LoginRegisterLoggerInfo("Register", "schema.resolvers.go", logMessage)

	return &model.AuthResponse{
		AuthToken: token,
		User:      &newUser,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.UserService.GetUserByEmail(input.Email)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't get user by email %s - %s ", input.Email, err.Error())
		logger.LoginRegisterLoggerError("Login", "schema.resolvers.go", logMessage)
		return nil, ErrorBadCredential
	}

	err = user.ComparePassword(input.Password)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't compare password for %s - %s ", input.Email, err.Error())
		logger.LoginRegisterLoggerError("Login", "schema.resolvers.go", logMessage)
		return nil, ErrorBadCredential
	}

	token, err := user.GenToken()

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't generate token for %s - %s ", input.Email, err.Error())
		logger.LoginRegisterLoggerError("Login", "schema.resolvers.go", logMessage)
		return nil, errors.New("something wrong with token")
	}

	logMessage := fmt.Sprintf("User was logined successfully with email - %s ", input.Email)
	logger.LoginRegisterLoggerInfo("Login", "schema.resolvers.go", logMessage)

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromContext(ctx)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't get user from context - %s", err.Error())
		logger.SystemLoggerError("CreateMeetup", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	newMeetup := model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	meetup, err := r.MeetupsService.CreateMeetup(newMeetup)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't create meetup - %s", err.Error())
		logger.CrudLoggerError(ctx, "CreateMeetup", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Meetup was created successfully with name - %s ", input.Name)
	logger.CrudLoggerInfo(ctx, "CreateMeetup", "schema.resolvers.go", logMessage)

	return meetup, nil
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input *model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := r.MeetupsService.UpdateMeetup(id, input)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't update meetup - %s", err.Error())
		logger.CrudLoggerError(ctx, "UpdateMeetup", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Meetup was updated successfully with name - %s ", input.Name)
	logger.CrudLoggerInfo(ctx, "UpdateMeetup", "schema.resolvers.go", logMessage)

	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (*bool, error) {
	result, err := r.MeetupsService.DeleteMeetup(id)
	if err != nil {
		logMessage := fmt.Sprintf("Couldn't delete meetup - %s", err.Error())
		logger.CrudLoggerError(ctx, "DeleteMeetup", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Meetup was deleted successfully with id - %s ", id)
	logger.CrudLoggerInfo(ctx, "DeleteMeetup", "schema.resolvers.go", logMessage)

	return result, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.UserService.CreateUser(input)
	if err != nil {
		logMessage := fmt.Sprintf("Couldn't create user - %s", err.Error())
		logger.CrudLoggerError(ctx, "CreateUser", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("User was creates successfully with email - %s ", input.Email)
	logger.CrudLoggerInfo(ctx, "CreateUser", "schema.resolvers.go", logMessage)

	return user, nil
}

func (r *queryResolver) GetAllMeetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*model.Meetup, error) {
	meetup, err := r.MeetupsService.GetAllMeetups(filter, *limit, *offset)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't get all meetups - %s", err.Error())
		logger.CrudLoggerError(ctx, "GetAllMeetups", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Get all meetups  successfully ")
	logger.CrudLoggerInfo(ctx, "GetAllMeetups", "schema.resolvers.go", logMessage)

	return meetup, nil
}

func (r *queryResolver) GetMeetupByID(ctx context.Context, id string) (*model.Meetup, error) {
	meetup, err := r.MeetupsService.GetMeetupByID(id)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't get meetup by id  - %s", err.Error())
		logger.CrudLoggerError(ctx, "GetMeetupByID", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Get meetup by id successfully with id %s ", id)
	logger.CrudLoggerInfo(ctx, "GetMeetupByID", "schema.resolvers.go", logMessage)

	return meetup, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	user, err := r.UserService.GetAllUsers()

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't get all users - %s", err.Error())
		logger.CrudLoggerError(ctx, "GetAllUsers", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Get all users successfully")
	logger.CrudLoggerInfo(ctx, "GetAllUsers", "schema.resolvers.go", logMessage)

	return user, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := r.UserService.GetUserByID(id)

	if err != nil {
		logMessage := fmt.Sprintf("Couldn't  get user by id  %s", err.Error())
		logger.CrudLoggerError(ctx, "GetUserByID", "schema.resolvers.go", logMessage)
		return nil, UserUnauthenticated
	}

	logMessage := fmt.Sprintf("Get user by id successfully with id %s ", id)
	logger.CrudLoggerInfo(ctx, "GetUserByID", "schema.resolvers.go", logMessage)

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
