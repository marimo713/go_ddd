//go:generate mockgen -source=$GOFILE -destination=../../utils/mockgen/$GOPACKAGE/mock_$GOFILE -package=$GOPACKAGE
package repository

import entity_book "my-app/domain/entity/book"

type BookRepository interface {
	GetByID(uint64) (*entity_book.Book, error)
	GetAll() ([]entity_book.Book, error)
	Create(entity_book.Book) (*entity_book.Book, error)
	Update(entity_book.Book) (*entity_book.Book, error)
	Delete(uint64) error
}
