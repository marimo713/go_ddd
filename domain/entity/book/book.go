package entity_book

import validation "github.com/go-ozzo/ozzo-validation"

type Book struct {
	id     uint64
	isbn   string
	title  string
	author string
}

func NewBook(isbn string, title string, author string) (*Book, error) {
	book := Book{
		isbn:   isbn,
		title:  title,
		author: author,
	}

	if err := book.validate(); err != nil {
		return nil, err
	}
	return &book, nil
}

func NewBookForRebuild(id uint64, isbn string, title string, author string) (*Book, error) {

	book := Book{
		id:     id,
		isbn:   isbn,
		title:  title,
		author: author,
	}

	if err := book.validate(); err != nil {
		return nil, err
	}

	return &book, nil
}

func (b Book) validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.isbn, validation.Required, validation.Length(13, 13)),
		validation.Field(&b.title, validation.Length(0, 255)),
		validation.Field(&b.author, validation.Length(0, 255)),
	)
}

func (b Book) ID() uint64 {
	return b.id
}

func (b Book) Isbn() string {
	return b.isbn
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Author() string {
	return b.author
}
