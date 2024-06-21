package db

import (
    "context"
    "testing"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/josefilo/caraCracha/config/env"
)

func TestConnect(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongodbConnection := env.NewMongoDBConnection()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbConnection.URI()))
    if err != nil {
        t.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        t.Fatalf("Failed to ping MongoDB: %v", err)
    }
}