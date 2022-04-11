package mongodb

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	collection *mongo.Collection
}

func NewUser(db *mongo.Client) *User {
	return &User{
		collection: db.Database("homestead").Collection("users"),
	}
}

func (u *User) Create(ctx context.Context, userName string) error {
	insert := core.User{
		ID:         primitive.NewObjectID(),
		Name:       userName,
		SaltToken:  "someToken",
		Categories: []core.Category{},
	}
	_, err := u.collection.InsertOne(ctx, insert)

	return err
}
