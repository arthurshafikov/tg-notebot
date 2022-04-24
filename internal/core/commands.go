package core

const (
	StartCommand = "start"

	AddCategoryCommand    = "addcategory"
	RemoveCategoryCommand = "removecategory"
	RenameCategoryCommand = "renamecategory"
	ListCategoriesCommand = "listcategories"

	AddNoteCommand     = "addnote"
	RemoveNotesCommand = "removenotes"
	ListNotes          = "listnotes"
	ListAllNotes       = "listallnotes"

	// CallbackQuery.
	RemoveNotesChooseCategoryCallbackQuery = "removenoteschoosecategory"
	ListNotesChooseCategoryCallbackQuery   = "listnoteschoosecategory"
	SpecialDelimeterInQueryCallback        = "_#_"
)
