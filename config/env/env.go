package env

import (
	"fmt"
	"os"
)

type MongodbConnection struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DB       string
}

func NewMongoDBConnection() MongodbConnection {
	return MongodbConnection{
		HOST:     os.Getenv("MONGO_HOST"),
        PORT:     os.Getenv("MONGO_PORT"),
        USER:     os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
        PASSWORD: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}
}

func (MongodbConnection) URI() string {
	connection := getMongodbConnection()
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		connection.USER,
		connection.PASSWORD,
		connection.HOST,
		connection.PORT
	)
}
