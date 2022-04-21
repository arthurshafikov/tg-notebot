package mongodb

import (
	"context"
	"errors"

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

func (u *User) CreateIfNotExists(ctx context.Context, telegramChatID int64) error {
	filter := bson.M{
		"telegram_chat_id": telegramChatID,
	}
	user := core.User{
		ID:             primitive.NewObjectID(),
		TelegramChatID: telegramChatID,
		Categories:     []core.Category{},
	}
	update := bson.M{
		"$setOnInsert": user,
	}
	opts := options.Update().SetUpsert(true)

	_, err := u.collection.UpdateOne(ctx, filter, update, opts)

	return err
}

func (u *User) CheckChatIDExists(ctx context.Context, telegramChatID int64) error {
	filter := bson.M{"telegram_chat_id": telegramChatID}
	res := u.collection.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return core.ErrNotFound
		}

		return err
	}

	var user core.User
	if err := res.Decode(&user); err != nil {
		return err
	}

	if user.TelegramChatID != telegramChatID {
		return core.ErrNotFound
	}

	return nil
}
