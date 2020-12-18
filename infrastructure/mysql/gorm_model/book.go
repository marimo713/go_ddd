package gorm_model

import (
	entity_book "my-app/domain/entity/book"
)

type Book struct {
	ID     uint64 `gorm:"primaryKey"`
	Isbn   string `gorm:"type:char(13)"`
	Title  string `gorm:"type:varchar(255)"`
	Author string `gorm:"type:varchar(255)"`
}

func NewBookFromDomain(book entity_book.Book) Book {
	return Book{
		ID:     book.ID(),
		Isbn:   book.Isbn(),
		Title:  book.Title(),
		Author: book.Author(),
	}
}

func (book Book) ToDomain() (*entity_book.Book, error) {
	return entity_book.NewBookForRebuild(
		book.ID,
		book.Isbn,
		book.Title,
		book.Author,
	)
}
