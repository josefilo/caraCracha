package repository

import (
	"context"

	"github.com/josefilo/caraCracha/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		collection: collection
	}
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryMongo) GetUser(ctx context.Context, id uint) (*entity.User, error) {
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
