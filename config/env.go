package config

import "github.com/joho/godotenv"

err := godotenv.Load()
if err != nil {
	logger.Error("Error loading .env file",
		zap.Any("error", err))
}