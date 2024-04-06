package request

import 
(
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)


type CreateUser struct {
	email     string `json:"email"`
	password  string `json:"password"`
	name      string `json:"name"`
	birthDate string `json:"birth_date"`
}

func (c createUser) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.email, validation.Required, is.Email),
		validation.Field(&c.password, validation.Required, validation.Length(1, 255)),
		validation.Field(&c.name, validation.Required, validation.Length(1, 255)),
		validation.Field(&c.birthDate, validation.Required, validation.Date("2006-01-02")),	
	)
}
