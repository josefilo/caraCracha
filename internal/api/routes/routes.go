package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/josefilo/caraCracha/internal/api/handler/userhandler"
)

// InitRoutes initializes the routes for the application
func InitRoutes(r *gin.RouterGroup) {
	r.POST("/user", userhandler.CreateUser)
}
