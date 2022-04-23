package mongodb

import (
	"context"
	"errors"

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

func (n *Note) AddNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	match := bson.M{
		"$and": []interface{}{
			bson.M{"telegram_chat_id": telegramChatID},
			bson.M{"categories.name": categoryName},
		},
	}
	change := bson.M{"$push": bson.M{"categories.$.notes": core.Note{
		Content: content,
	}}}

	return n.collection.FindOneAndUpdate(ctx, match, change).Err()
}

func (n *Note) ListNotesFromCategory(
	ctx context.Context,
	telegramChatID int64,
	categoryName string,
) ([]core.Note, error) {
	var notes []core.Note

	filter := bson.M{"telegram_chat_id": telegramChatID}

	res := n.collection.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return notes, core.ErrNotFound
		}

		return notes, err
	}

	var user core.User
	if err := res.Decode(&user); err != nil {
		return notes, err
	}

	for _, cat := range user.Categories {
		if cat.Name == categoryName {
			notes = append(notes, cat.Notes...)
		}
	}

	return notes, nil
}

func (n *Note) RemoveNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	match := bson.M{
		"$and": []interface{}{
			bson.M{"telegram_chat_id": telegramChatID},
			bson.M{"categories.name": categoryName},
		},
	}
	change := bson.M{"$pull": bson.M{"categories.$.notes": bson.M{
		"content": content,
	}}}

	err := n.collection.FindOneAndUpdate(ctx, match, change).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return core.ErrNotFound
	}

	return err
}
