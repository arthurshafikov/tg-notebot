package mongodb

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Category struct {
	collection *mongo.Collection
}

func NewCategory(db *mongo.Client) *Category {
	return &Category{
		collection: db.Database("homestead").Collection("categories"),
	}
}

func (c *Category) AddCategory(ctx context.Context, userId int, name string) error {
	_, err := c.collection.InsertOne(ctx, bson.D{
		{Key: "user_id", Value: userId},
		{Key: "name", Value: name},
		{Key: "posts", Value: []core.Note{}},
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Category) RemoveCategory(ctx context.Context, userId int, name string) error {
	res, err := c.collection.DeleteOne(ctx, bson.D{{"user_id", userId}, {"name", name}})
	if err != nil {
		return err
	}

	if res.DeletedCount < 1 {
		return core.ErrNothingWasDeleted
	}

	return nil
}

func (c *Category) RenameCategory(ctx context.Context, userId int, name, newName string) error {
	res, err := c.collection.UpdateOne(
		ctx,
		bson.D{{"user_id", userId}, {"name", name}},
		bson.D{
			{"$set", bson.D{{"name", newName}}},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount < 1 {
		return core.ErrNothingWasUpdated
	}

	return nil
}

func (c *Category) ListCategories(ctx context.Context, userId int) ([]core.Category, error) {
	var categories []core.Category

	cursor, err := c.collection.Find(ctx, bson.D{{"user_id", userId}})
	if err != nil {
		return categories, err
	}

	if err = cursor.All(ctx, &categories); err != nil {
		return categories, err
	}

	return categories, nil
}
