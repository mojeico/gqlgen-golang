package model

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	tokenTime = 24 * 7 * time.Hour // 7 days for token
	signInKey = "registerKey"
)

type User struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Meetups   []*Meetup `json:"meetups"`
	CratedAt  int64     `json:"crated_at"`
	UpdatedAt int64     `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
}

func (u *User) HashPassword(password string) error {

	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Role   string `json:"user_role"`
}

func (u *User) GenToken() (*AuthToken, error) {

	expirationAt := time.Now().Add(tokenTime).UnixNano() / int64(time.Millisecond)
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	hmacSampleSecret := []byte(signInKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expirationAt,
		Id:        u.ID,
		IssuedAt:  currentTime,
		Issuer:    "registerUser",
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccesToken: tokenString,
		ExpiredAt:  int(expirationAt),
	}, nil
}
