package models

type User struct {
	ID         uint64
	TelegramId uint64
	Name       string
	SaltToken  string
}
