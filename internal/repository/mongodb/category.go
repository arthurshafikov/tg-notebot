package mongodb

import (
	"context"
	"errors"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Category struct {
	collection *mongo.Collection
}

func NewCategory(db *mongo.Client) *Category {
	return &Category{
		collection: db.Database("homestead").Collection("users"),
	}
}

func (c *Category) AddCategory(ctx context.Context, telegramChatID int64, name string) error {
	filter := bson.M{"$and": []interface{}{
		bson.M{"telegram_chat_id": telegramChatID},
		bson.M{"categories.name": name},
	}}
	err := c.collection.FindOne(ctx, filter).Err()
	if err == nil {
		return core.ErrCategoryExists
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	match := bson.M{"telegram_chat_id": telegramChatID}
	change := bson.M{"$push": bson.M{"categories": core.Category{
		Name:  name,
		Notes: []core.Note{},
	}}}

	return c.collection.FindOneAndUpdate(ctx, match, change).Err()
}

func (c *Category) RemoveCategory(ctx context.Context, telegramChatID int64, name string) error {
	match := bson.M{"telegram_chat_id": telegramChatID}
	change := bson.M{"$pull": bson.M{"categories": bson.M{
		"name": name,
	}}}

	err := c.collection.FindOneAndUpdate(ctx, match, change).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return core.ErrNotFound
	}

	return err
}

func (c *Category) RenameCategory(ctx context.Context, telegramChatID int64, name, newName string) error {
	filter := bson.M{"$and": []interface{}{
		bson.M{"telegram_chat_id": telegramChatID},
		bson.M{"categories.name": newName},
	}}
	err := c.collection.FindOne(ctx, filter).Err()
	if err == nil {
		return core.ErrCategoryExists
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	match := bson.M{"$and": []interface{}{
		bson.M{"telegram_chat_id": telegramChatID},
		bson.M{"categories.name": name},
	}}
	change := bson.M{"$set": bson.M{"categories.$.name": newName}}

	err = c.collection.FindOneAndUpdate(ctx, match, change).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return core.ErrNotFound
	}

	return err
}

func (c *Category) ListCategories(ctx context.Context, telegramChatID int64) ([]core.Category, error) {
	var user core.User

	res := c.collection.FindOne(ctx, bson.M{"telegram_chat_id": telegramChatID})
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []core.Category{}, core.ErrNotFound
		}

		return []core.Category{}, err
	}

	if err := res.Decode(&user); err != nil {
		return []core.Category{}, err
	}

	return user.Categories, nil
}
