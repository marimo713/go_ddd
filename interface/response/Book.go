package response

import entity_book "my-app/domain/entity/book"

type Book struct {
	ID     uint64 `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewBookFromDomain(book entity_book.Book) Book {
	return Book{
		ID:     book.ID(),
		Isbn:   book.Isbn(),
		Title:  book.Title(),
		Author: book.Author(),
	}
}
