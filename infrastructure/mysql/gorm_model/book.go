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

func (book Book) ToDomain() entity_book.Book {
	return entity_book.NewBook(
		book.ID,
		book.Isbn,
		book.Title,
		book.Author,
	)
}
