package repository

import (
	"context"
	"fmt"

	"github.com/josefilo/caraCracha/internal/entity"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) entity.IUserRepository {
	return &UserRepositoryMongo{
		collection: collection,
	}
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	InsertResult, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "insert user failed")
	}
	zap.L().Info(fmt.Sprintf("User created with ID: %s", InsertResult.InsertedID))
	return user, nil
}

func (r *UserRepositoryMongo) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user := &entity.User{}
	err := r.collection.FindOne(ctx, entity.User{ID: id}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryMongo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.collection.FindOne(ctx, entity.User{Email: email}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryMongo) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := r.collection.ReplaceOne(ctx, entity.User{ID: user.ID}, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
