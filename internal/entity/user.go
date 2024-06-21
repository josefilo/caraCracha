package entity

import "context"

type User struct {
	ID        uint   `bson:"_id"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	Name      string `bson:"name"`
	BirthDate string `bson:"birthDate"`
	CreatedAt string `bson:"createdAt"`
	UpdatedAt string `bson:"updatedAt"`
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id uint) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
}
