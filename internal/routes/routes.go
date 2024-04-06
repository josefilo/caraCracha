package router

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id")
	r.GET("/userByEmail/:email")
	r.POST("/user")
}
