package core

type Category struct {
	Name  string `bson:"name"`
	Notes []Note `bson:"notes"`
}
