package db

import (
	"context"
	"time"

	"github.com/josefilo/caraCracha/config/env"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDBClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbConnection := env.NewMongoDBConnection()
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(mongodbConnection.URI()).
			SetMaxPoolSize(10).
			SetMinPoolSize(2).
			SetRetryWrites(true),
	)
	if err != nil {
		return nil, errors.Wrap(err, "connection to MongoDB failed")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "ping to MongoDB failed")
	}
	return client, nil
}

func NewMongoDBCollection(databaseName, collectionName string) (*mongo.Collection, error) {
	client, err := NewMongoDBClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil, errors.New("creation of collection failed")
	}
	return collection, nil
}
