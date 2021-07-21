package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

func (r RegisterInput) IsValidate() error {

	err := validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required, validation.Length(5, 10)),
		validation.Field(&r.Email, validation.Required, is.Email),
		//validation.Field(&r.Password, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]$"))),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 15)),
		validation.Field(&r.FirstName, validation.Required, validation.Length(5, 15)),
		validation.Field(&r.LastName, validation.Required, validation.Length(5, 15)),
	)

	return err

}
