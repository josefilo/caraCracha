package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josefilo/internal/routes"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

func main() {

	err := godotenv.Load()
if err != nil {
	logger.Error("Error loading .env file",
		zap.Any("error", err))
}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	logger.Info("Starting the application")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to start the application",
			zap.Any("error", err))
	}
}
