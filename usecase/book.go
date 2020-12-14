//go:generate mockgen -source=$GOFILE -destination=../utils/mockgen/$GOPACKAGE/mock_$GOFILE -package=$GOPACKAGE
package usecase

import (
	entity_book "my-app/domain/entity/book"
	"my-app/domain/repository"
)

type BookUsecase interface {
	GetByID(uint64) (*entity_book.Book, error)
	GetAll() ([]entity_book.Book, error)
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	return bookUsecase{
		bookRepository: bookRepository,
	}
}

func (usecase bookUsecase) GetByID(id uint64) (*entity_book.Book, error) {

	return usecase.bookRepository.GetByID(id)
}

func (usecase bookUsecase) GetAll() ([]entity_book.Book, error) {

	return usecase.bookRepository.GetAll()
}
