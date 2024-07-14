package repository

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/josefilo/caraCracha/config/db"
	"github.com/josefilo/caraCracha/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	// Setup
	is := require.New(t)
	ctx := context.Background()
	databaseName := "caraCrach"
	collectionName := "user"
	db, err := db.NewMongoDBCollection(databaseName, collectionName)
	is.Nil(err)
	userRepository := NewUserRepositoryMongo(db)

	t.Run("CreateUser", func(t *testing.T) {
		// Given
		is := require.New(t)
		user := &entity.User{
			ID:        faker.UUIDDigit(),
			Email:     faker.Email(),
			Password:  faker.Password(),
			Name:      faker.Name(),
			BirthDate: faker.Date(),
			CreatedAt: faker.Date(),
			UpdatedAt: faker.Date(),
		}

		// When
		createdUser, err := userRepository.CreateUser(context.Background(), user)

		// Then
		is.Nil(err)
		is.NotNil(createdUser)
		is.Equal(user.Email, createdUser.Email)
		is.Equal(user.Password, createdUser.Password)
	})

	t.Run("CreateUser failed", func(t *testing.T) {
		// Given
		is := require.New(t)
		user := &entity.User{
			ID:        faker.UUIDDigit(),
			Email:     faker.Email(),
			Password:  faker.Password(),
			Name:      faker.Name(),
			BirthDate: faker.Date(),
			CreatedAt: faker.Date(),
			UpdatedAt: faker.Date(),
		}

		// When
		createdUser, err := userRepository.CreateUser(ctx, user)
		is.Nil(err)
		createdUser.ID = faker.UUIDDigit()
		_, err = userRepository.CreateUser(ctx, createdUser)

		// Then
		//fmt.Println("Error: ", err.Error())
		is.NotNil(err)
		is.Contains(err.Error(), "user index: _email_1 dup key")
	})
}
