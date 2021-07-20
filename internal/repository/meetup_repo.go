package repository

import (
	"context"
	"fmt"
	"github.com/mojeico/gqlgen-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MeetupsRepo interface {
	GetAllMeetups(filter *model.MeetupFilter, limit int64, offset int64) ([]*model.Meetup, error)
	CreateMeetup(meetup model.Meetup) (*model.Meetup, error)
	GetMeetupByID(id string) (*model.Meetup, error)
	UpdateMeetup(id string, meetup *model.UpdateMeetup) (*model.Meetup, error)
	DeleteMeetup(id string) (*bool, error)
}

type meetupsRepo struct {
	client *mongo.Client
}

func NewMeetupsRepo(client *mongo.Client) MeetupsRepo {
	return &meetupsRepo{
		client: client,
	}
}

func (repo meetupsRepo) DeleteMeetup(id string) (*bool, error) {

	collection := repo.client.Database("myapp").Collection("meetup")
	mongoId, _ := primitive.ObjectIDFromHex(id)

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": mongoId})

	resultBool := result.DeletedCount != 0
	return &resultBool, err
}

func (repo meetupsRepo) UpdateMeetup(id string, meetup *model.UpdateMeetup) (*model.Meetup, error) {

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

	var updatedModel model.Meetup

	err = collection.FindOne(context.Background(), bson.M{"_id": mongoId}).Decode(&updatedModel)

	if err != nil {
		print(err)
	}

	return &updatedModel, err
}

func (repo *meetupsRepo) GetAllMeetups(filter *model.MeetupFilter, limit int64, offset int64) ([]*model.Meetup, error) {

	collection := repo.client.Database("myapp").Collection("meetup")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoFilter := bson.M{}
	if *filter.Name != "" {
		mongoFilter = bson.M{"name": bson.M{"$regex": primitive.Regex{
			Pattern: fmt.Sprintf("^([%s])\\w+", *filter.Name),
			Options: "i",
		}}}
	}

	var tasks []*model.Meetup

	opts := options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	}

	cur, err := collection.Find(ctx, mongoFilter, &opts)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t *model.Meetup
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

func (repo *meetupsRepo) CreateMeetup(meetup model.Meetup) (*model.Meetup, error) {

	ctx := context.Background()

	coll := repo.client.Database("myapp").Collection("meetup")

	_, err := coll.InsertOne(ctx, &meetup)

	return &model.Meetup{}, err
}

func (repo *meetupsRepo) GetMeetupByID(id string) (*model.Meetup, error) {

	collection := repo.client.Database("myapp").Collection("meetup")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoId, _ := primitive.ObjectIDFromHex(id)

	var meetup model.Meetup
	err := collection.FindOne(ctx, bson.M{"_id": mongoId}).Decode(&meetup)

	if err != nil {
		log.Fatal(err)
	}

	return &meetup, err

}
