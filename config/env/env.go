package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}
	return nil
}

type MongodbConnection struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
}

func NewMongoDBConnection() MongodbConnection {

	err := Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}
	return MongodbConnection{
		HOST:     os.Getenv("MONGODB_HOST"),
		PORT:     os.Getenv("MONGODB_PORT"),
		USER:     os.Getenv("MONGODB_USER"),
		PASSWORD: os.Getenv("MONGODB_PASSWORD"),
	}
}

func (mongodbConnection MongodbConnection) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		mongodbConnection.USER,
		mongodbConnection.PASSWORD,
		mongodbConnection.HOST,
		mongodbConnection.PORT,
	)
}
