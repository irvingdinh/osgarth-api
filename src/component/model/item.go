package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID primitive.ObjectID `bson:"_id"`

	Type    string `bson:"type"`
	Payload bson.M `bson:"payload"`

	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
