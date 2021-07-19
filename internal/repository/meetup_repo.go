package repository

import (
	"context"
	"github.com/mojeico/gqlgen-golang/graph/model"
	"github.com/mojeico/gqlgen-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"time"
)

type MeetupsRepo interface {
	GetAllMeetups() ([]*models.Meetup, error)
	CreateMeetup(meetup model.NewMeetup) (*models.Meetup, error)
	GetMeetupByID(id string) *models.Meetup
	UpdateMeetup(id string, meetup *model.UpdateMeetup) (*models.Meetup, error)
}

type meetupsRepo struct {
	client *mongo.Client
}

func NewMeetupsRepo(client *mongo.Client) MeetupsRepo {
	return &meetupsRepo{
		client: client,
	}
}

func (repo meetupsRepo) UpdateMeetup(id string, meetup *model.UpdateMeetup) (*models.Meetup, error) {

	collection := repo.client.Database("myapp").Collection("meetup")
	mongoId, _ := primitive.ObjectIDFromHex(id)

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": mongoId},
		bson.D{
			{"$set", bson.D{
				{"name", meetup.Name},
				{"description", meetup.Description},
			}},
		},
	)

	var updatedModel models.Meetup

	err = collection.FindOne(context.Background(), bson.M{"_id": mongoId}).Decode(&updatedModel)

	if err != nil {
		print(err)
	}

	return &updatedModel, err
}

func (repo *meetupsRepo) GetAllMeetups() ([]*models.Meetup, error) {

	collection := repo.client.Database("myapp").Collection("meetup")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{}}

	var tasks []*models.Meetup

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t *models.Meetup
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	//cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil

}

func (repo *meetupsRepo) CreateMeetup(meetup model.NewMeetup) (*models.Meetup, error) {

	ctx := context.Background()

	coll := repo.client.Database("myapp").Collection("meetup")

	_, err := coll.InsertOne(ctx, &meetup)

	return &models.Meetup{}, err
}

func (repo *meetupsRepo) GetMeetupByID(id string) *models.Meetup {

	collection := repo.client.Database("myapp").Collection("meetup")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoId, _ := primitive.ObjectIDFromHex(id)

	var meetup models.Meetup
	err := collection.FindOne(ctx, bson.M{"_id": mongoId}).Decode(&meetup)

	if err != nil {
		log.Fatal(err)
	}

	return &meetup

}
