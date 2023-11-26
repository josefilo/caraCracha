package userControlers

import (
	"fmt"

	"github.com/Josefreitas788/gestao_de_financas/src/configuration/logger"
	"github.com/Josefreitas788/gestao_de_financas/src/configuration/validation"
	"github.com/Josefreitas788/gestao_de_financas/src/controller/model/request"
	"github.com/Josefreitas788/gestao_de_financas/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to bind json",
			zap.Any("error", err),
		)
		apiErrors := validation.ValidationUserError(err)
		c.JSON(apiErrors.Code, apiErrors)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.BirthDate,
	)

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error trying to create user",
			zap.Any("error", err),
		)
		c.JSON(err.Code, err)
		return
	}

	fmt.Println(userRequest)
}
