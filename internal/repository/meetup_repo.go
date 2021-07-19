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

type MeetupsRepo interface {
	GetAllMeetups() ([]*models.Meetup, error)
	CreateMeetup(meetup model.NewMeetup) (*models.Meetup, error)
	GetMeetupByID(id string) *models.Meetup
}

type meetupsRepo struct {
	client *mongo.Client
}

func NewMeetupsRepo(client *mongo.Client) MeetupsRepo {
	return &meetupsRepo{
		client: client,
	}
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

	result, err := collection.Find(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal(err)
	}

	meetup := &models.Meetup{}
	result.Decode(meetup)

	return meetup

}
