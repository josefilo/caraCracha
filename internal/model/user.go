package model

type User struct {
    ID        uint          `bson:"_id"`
    Email     string        `bson:"email"`
    Password  string        `bson:"password"`
    Name      string        `bson:"name"`
    BirthDate string        `bson:"birthDate"`
    CreatedAt string        `bson:"createdAt"`
    UpdatedAt string        `bson:"updatedAt"`
}

