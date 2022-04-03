package services

type Categories interface {
	AddCategory() error
	RemoveCategory() error
	RenameCategory() error
	ListCategories() error
}

type Notes interface {
	AddNote() error
	ListNotes() error
	RemoveNotes() error
}

type Services struct {
	Categories
	Notes
}

func NewServices() *Services {
	return &Services{}
}
