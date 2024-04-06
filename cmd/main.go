package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josefilo/internal/routes"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

func main() {

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	logger.Info("Starting the application")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to start the application",
			zap.Any("error", err))
	}
}
