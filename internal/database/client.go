package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	conn *mongo.Database
}

func NewMongoClient(ctx context.Context, uri, dbName string) (*MongoClient, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongodb connection configuration: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mongodb deployment: %w", err)
	}

	mongoClient := MongoClient{
		conn: client.Database(dbName),
	}

	return &mongoClient, nil
}
