package main

import (
	"github.com/Josefreitas788/gestao_de_financas/src/configuration/logger"
	"github.com/Josefreitas788/gestao_de_financas/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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
