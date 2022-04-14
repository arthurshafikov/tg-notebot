package mongodb

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	collection *mongo.Collection
}

func NewUser(db *mongo.Client) *User {
	return &User{
		collection: db.Database("homestead").Collection("users"),
	}
}

func (u *User) CreateIfNotExists(ctx context.Context, userName string, telegramChatID int64) error {
	filter := bson.M{
		"telegram_chat_id": telegramChatID,
	}
	user := core.User{
		ID:             primitive.NewObjectID(),
		TelegramChatID: telegramChatID,
		SaltToken:      "someToken123", // todo generate token
		Categories:     []core.Category{},
	}
	update := bson.M{
		"$setOnInsert": user,
	}
	opts := options.Update().SetUpsert(true)

	_, err := u.collection.UpdateOne(ctx, filter, update, opts)

	return err
}
