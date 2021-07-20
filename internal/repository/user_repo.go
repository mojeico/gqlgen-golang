package repository

import (
	"context"
	"github.com/mojeico/gqlgen-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepo interface {
	GetAllUsers() ([]*model.User, error)
	CreateUser(meetup model.NewUser) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUserName(userName string) (*model.User, error)
	RegistrationUser(user model.User) (string, error)
}

type userRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) UserRepo {
	return &userRepo{
		client: client,
	}
}

func (repo *userRepo) GetAllUsers() ([]*model.User, error) {

	collection := repo.client.Database("myapp").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{}}

	var tasks []*model.User

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t *model.User
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

func (repo *userRepo) CreateUser(newUser model.NewUser) (*model.User, error) {

	ctx := context.Background()

	coll := repo.client.Database("myapp").Collection("user")

	_, err := coll.InsertOne(ctx, &newUser)

	return &model.User{}, err
}

func (repo *userRepo) RegistrationUser(user model.User) (string, error) {
	ctx := context.Background()

	coll := repo.client.Database("myapp").Collection("user")

	result, err := coll.InsertOne(ctx, &user)

	id := result.InsertedID.(primitive.ObjectID)

	return id.Hex(), err
}

func (repo *userRepo) GetUserByID(id string) (*model.User, error) {

	collection := repo.client.Database("myapp").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoId, _ := primitive.ObjectIDFromHex(id)

	var user model.User
	err := collection.FindOne(ctx, bson.M{"_id": mongoId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, err

}

func (repo *userRepo) GetUserByEmail(email string) (*model.User, error) {

	collection := repo.client.Database("myapp").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var user model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, err

}

func (repo *userRepo) GetUserByUserName(userName string) (*model.User, error) {

	collection := repo.client.Database("myapp").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var user model.User
	err := collection.FindOne(ctx, bson.M{"username": userName}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, err

}
