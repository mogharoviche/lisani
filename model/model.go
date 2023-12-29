package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"Name,omitempty" bson:"name,omitempty"`
	Age     int    `json:"age" bson:"age"`
	Address string `json:"Address,omitempty" bson:"address,omitempty"`
}

func CreateStudent(student Student, collection *mongo.Collection) error {
	_, err := collection.InsertOne(context.Background(), student)
	return err
}

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"Username" bson:"Username"`
	Password string `json:"Password" bson:"Password"`
	Email    string `json:"Email,omitempty" bson:"Email,omitempty"`
}
