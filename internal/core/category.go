package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Notes []Note             `bson:"notes"`
}
