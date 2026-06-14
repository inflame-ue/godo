package models

import "go.mongodb.org/mongo-driver/v2/bson"

type TODO struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string        `bson:"title" json:"title"`
	Completed bool          `bson:"completed" json:"completed"`
}
