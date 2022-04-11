package mongodb

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	collection *mongo.Collection
}

func NewNote(db *mongo.Client) *Note {
	return &Note{
		collection: db.Database("homestead").Collection("users"),
	}
}

func (n *Note) AddNote(ctx context.Context, userName string, categoryName string, note string) error {
	match := bson.M{"$and": []interface{}{bson.M{"name": userName}, bson.M{"categories.name": categoryName}}}
	change := bson.M{"$push": bson.M{"categories.$.notes": core.Note{
		Content: note,
	}}}

	return n.collection.FindOneAndUpdate(ctx, match, change).Err()
}

func (n *Note) ListNotes(ctx context.Context, userName string, categoryName string) ([]core.Note, error) {
	var notes []core.Note

	filter := bson.M{"$and": []interface{}{bson.M{"name": userName}, bson.M{"categories.name": categoryName}}}

	res := n.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return notes, res.Err()
	}

	var user core.User
	if err := res.Decode(&user); err != nil {
		return notes, err
	}

	for _, cat := range user.Categories {
		notes = append(notes, cat.Notes...)
	}

	return notes, nil
}

func (n *Note) RemoveNotes(ctx context.Context, userName, categoryName, noteContent string) error {
	match := bson.M{"$and": []interface{}{bson.M{"name": userName}, bson.M{"categories.name": categoryName}}}
	change := bson.M{"$pull": bson.M{"categories.$.notes": bson.M{
		"content": noteContent,
	}}}

	return n.collection.FindOneAndUpdate(ctx, match, change).Err()
}
