package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUser struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

func (c CreateUser) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Password, validation.Required, validation.Length(1, 255)),
		validation.Field(&c.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&c.BirthDate, validation.Required, validation.Date("2006-01-02")),
	)
}
