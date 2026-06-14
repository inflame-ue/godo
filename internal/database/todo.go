package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/inflame-ue/godo/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "todos"

func (mc *MongoClient) InsertTODO(ctx context.Context, title string, completed bool) (*models.TODO, error) {
	coll := mc.conn.Collection(collectionName)
	todo := models.TODO{Title: title, Completed: completed}

	result, err := coll.InsertOne(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("inserting a todo: %w", err)
	}
	todo.ID = result.InsertedID.(bson.ObjectID)

	return &todo, nil
}

func (mc *MongoClient) GetAllTODO(ctx context.Context) ([]models.TODO, error) {
	coll := mc.conn.Collection(collectionName)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("making a cursor over all todos: %w", err)
	}

	var todos []models.TODO
	if err := cursor.All(ctx, &todos); err != nil {
		return nil, fmt.Errorf("populating a slice of todos from cursor: %w", err)
	}

	return todos, nil
}

func (mc *MongoClient) GetTODOByID(ctx context.Context, id bson.ObjectID) (*models.TODO, error) {
	coll := mc.conn.Collection(collectionName)
	filter := bson.D{{Key: "_id", Value: id}}

	var todo models.TODO
	if err := coll.FindOne(ctx, filter).Decode(&todo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no TODO with such an ID")
		}
		return nil, fmt.Errorf("finding one todo by object ID: %w", err)
	}

	return &todo, nil
}

func (mc *MongoClient) UpdateTODOByID(ctx context.Context, id bson.ObjectID, title string, completed bool) (*models.TODO, error) {
	coll := mc.conn.Collection(collectionName)
	todo := models.TODO{Title: title, Completed: completed}

	updateResult, err := coll.UpdateByID(ctx, id, bson.M{"$set": todo})
	if err != nil {
		return nil, fmt.Errorf("update todo by id: %w", err)
	}
	if updateResult.MatchedCount == 0 {
		return nil, errors.New("no TODO with the given object id was found in the collection")
	}
	todo.ID = id

	return &todo, nil
}

func (mc *MongoClient) DeleteTODOByID(ctx context.Context, id bson.ObjectID) error {
	coll := mc.conn.Collection(collectionName)
	filter := bson.D{{Key: "_id", Value: id}}

	deleteResult, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete todo by id: %w", err)
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("no TODO with the given object id was found in the collection")
	}

	return nil
}
