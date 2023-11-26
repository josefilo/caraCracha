package model

import (
	"github.com/Josefreitas788/gestao_de_financas/src/configuration/api_errors"
	"github.com/Josefreitas788/gestao_de_financas/src/configuration/logger"
)

func (u *UserDomain) DeleteUser() *api_errors.ApiError {

	logger.Info("Init DeleteUser")
	return nil
}
