package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/josefilo/caraCracha/internal/routes"
	"go.uber.org/zap"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		zap.L().Error("Error loading .env file",
			zap.Any("error", err))
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	zap.L().Info("Starting the application")
	if err := router.Run(":8080"); err != nil {
		zap.L().Error("Error starting the application",
			zap.Any("error", err))
	}
}
