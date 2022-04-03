package repository

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

type Repository struct {
	Categories
	Notes
}

func NewRepository() *Repository {
	return &Repository{}
}
