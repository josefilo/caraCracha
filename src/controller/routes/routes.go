package routes

import (
	"github.com/Josefreitas788/gestao_de_financas/src/controller/userControlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", userControlers.GetUser)
	r.GET("/userByEmail/:email", userControlers.GetUserByEmail)
	r.POST("/user", userControlers.CreateUser)
	r.PATCH("/user/:id", userControlers.UpdateUser)
	r.DELETE("/user/:id", userControlers.DeleteUser)

}
