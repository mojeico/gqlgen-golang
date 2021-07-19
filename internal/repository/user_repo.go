package repository

import (
	"context"
	"github.com/mojeico/gqlgen-golang/graph/model"
	"github.com/mojeico/gqlgen-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"time"
)

type UserRepo interface {
	GetAllUsers() ([]*models.User, error)
	CreateUser(meetup model.NewUser) (*models.User, error)
	GetUserByID(id string) *models.User
}

type userRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) UserRepo {
	return &userRepo{
		client: client,
	}
}

func (repo *userRepo) GetAllUsers() ([]*models.User, error) {

	collection := repo.client.Database("myapp").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{}}

	var tasks []*models.User

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t *models.User
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil

}

func (repo *userRepo) CreateUser(meetup model.NewUser) (*models.User, error) {

	ctx := context.Background()

	coll := repo.client.Database("myapp").Collection("user")

	_, err := coll.InsertOne(ctx, &meetup)

	return &models.User{}, err
}

func (repo *userRepo) GetUserByID(id string) *models.User {

	collection := repo.client.Database("myapp").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.Find(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal(err)
	}

	meetup := &models.User{}
	result.Decode(meetup)

	return meetup

}
