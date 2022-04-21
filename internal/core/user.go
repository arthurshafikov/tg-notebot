package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	TelegramChatID int64              `bson:"telegram_chat_id"`
	Categories     []Category         `bson:"categories"`
}
