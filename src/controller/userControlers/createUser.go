package userControlers

import (
	"fmt"
	"log"

	"github.com/Josefreitas788/gestao_de_financas/src/configuration/validation"
	"github.com/Josefreitas788/gestao_de_financas/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println("Error trying to bind json", err)
		apiErrors := validation.ValidationUserError(err)
		c.JSON(apiErrors.Code, apiErrors)
		return
	}
	fmt.Println(userRequest)
}
