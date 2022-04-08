package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	SaltToken  string             `bson:"salt_token"`
	Categories []Category         `bson:"categories"`
}
