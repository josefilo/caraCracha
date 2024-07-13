package env

import (
	"os"
	"testing"
)

func TestNewMongoDBConnection(t *testing.T) {
	// Given

	err := Load()
	if err != nil {
		t.Errorf("Error loading .env file: %v", err)
	}

	expectedMongodbConnection := MongodbConnection{
		HOST:     os.Getenv("MONGODB_HOST"),
		PORT:     os.Getenv("MONGODB_PORT"),
		USER:     os.Getenv("MONGODB_USER"),
		PASSWORD: os.Getenv("MONGODB_PASSWORD"),
	}

	// When
	mongodbConnection := NewMongoDBConnection()

	// Then
	if mongodbConnection.HOST != expectedMongodbConnection.HOST {
		t.Errorf("Expected HOST to be %s, got %s", expectedMongodbConnection.HOST, mongodbConnection.HOST)
	}
	if mongodbConnection.PORT != expectedMongodbConnection.PORT {
		t.Errorf("Expected PORT to be %s, got %s", expectedMongodbConnection.PORT, mongodbConnection.PORT)
	}

	if mongodbConnection.USER != expectedMongodbConnection.USER {
		t.Errorf("Expected USER to be %s, got %s", expectedMongodbConnection.USER, mongodbConnection.USER)
	}

	if mongodbConnection.PASSWORD != expectedMongodbConnection.PASSWORD {
		t.Errorf("Expected PASSWORD to be %s, got %s", expectedMongodbConnection.PASSWORD, mongodbConnection.PASSWORD)
	}
}
