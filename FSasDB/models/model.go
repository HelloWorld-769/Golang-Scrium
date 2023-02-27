package mod

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,

		validation.Field(&u.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Name, validation.Required, is.Email),
	)
}
