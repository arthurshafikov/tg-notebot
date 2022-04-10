package mongodb

import (
	"context"

	"github.com/arthurshafikov/wolf-notebot/internal/core"
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

func (c *Category) AddCategory(ctx context.Context, userName, name string) error {
	match := bson.M{"name": userName}
	change := bson.M{"$push": bson.M{"categories": core.Category{
		Name:  name,
		Notes: []core.Note{},
	}}}

	return c.collection.FindOneAndUpdate(ctx, match, change).Err()
}

func (c *Category) RemoveCategory(ctx context.Context, userName, name string) error {
	match := bson.M{"name": userName}
	change := bson.M{"$pull": bson.M{"categories": bson.M{
		"name": name,
	}}}

	return c.collection.FindOneAndUpdate(ctx, match, change).Err()
}

func (c *Category) RenameCategory(ctx context.Context, userName, name, newName string) error {
	match := bson.M{"$and": []interface{}{bson.M{"name": userName}, bson.M{"categories.name": name}}}
	change := bson.M{"$set": bson.M{"categories.$.name": newName}}

	return c.collection.FindOneAndUpdate(ctx, match, change).Err()
}

func (c *Category) ListCategories(ctx context.Context, userName string) ([]core.Category, error) {
	var user core.User

	res := c.collection.FindOne(ctx, bson.M{"name": userName})
	if res.Err() != nil {
		return []core.Category{}, res.Err()
	}

	if err := res.Decode(&user); err != nil {
		return []core.Category{}, err
	}

	return user.Categories, nil
}
