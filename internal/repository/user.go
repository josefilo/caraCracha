package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserReporitory interface {
	CreateUser(ctx context.Context,user *model.User) (*model.User,error)
	GetUser(ctx context.Context, id uint) (*model.User,error)
	GetUserByEmail(ctx context.Context, email string) (*model.User,error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User,error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserReporitory {
	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User,error) {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}


