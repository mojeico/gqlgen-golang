package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meetup struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	UserID      string             `json:"user_id"`
}
