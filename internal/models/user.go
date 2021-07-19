package models

type User struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Meetups   []*Meetup `json:"meetups"`
}
