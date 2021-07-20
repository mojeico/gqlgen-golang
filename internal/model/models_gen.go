// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccesToken string `json:"accesToken"`
	ExpiredAt  int    `json:"expiredAt"`
}

type MeetupFilter struct {
	Name *string `json:"name"`
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

type UpdateMeetup struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}