package model

import (
	"github.com/Josefreitas788/gestao_de_financas/src/configuration/api_errors"
	"github.com/Josefreitas788/gestao_de_financas/src/configuration/logger"
)

func (userDomain *UserDomain) CreateUser() *api_errors.ApiError {

	logger.Info("Init CreateUser model")
	userDomain.EncryptPassword()

	return nil
}
