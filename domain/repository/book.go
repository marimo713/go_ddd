//go:generate mockgen -source=$GOFILE -destination=../../utils/mockgen/$GOPACKAGE/mock_$GOFILE -package=$GOPACKAGE
package repository

import entity_book "my-app/domain/entity/book"

type BookRepository interface {
	GetByID(uint64) (*entity_book.Book, error)
}
