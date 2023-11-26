package model

import (
	"crypto/md5"

	"github.com/Josefreitas788/gestao_de_financas/src/configuration/api_errors"
)

func NewUserDomain(email, password, name, birthDate string) UserDomainInterface {
	return &UserDomain{
		Email:     email,
		Password:  password,
		Name:      name,
		BirthDate: birthDate,
	}
}

type UserDomain struct {
	Email     string
	Password  string
	Name      string
	BirthDate string
}

func (userDomain *UserDomain) EncryptPassword() {

	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(userDomain.Password))
	userDomain.Password = string(hash.Sum(nil))

}

type UserDomainInterface interface {
	CreateUser() *api_errors.ApiError
	UpdateUser(string) *api_errors.ApiError
	FindUser(string) (*UserDomain, *api_errors.ApiError)
	DeleteUser() *api_errors.ApiError
}
