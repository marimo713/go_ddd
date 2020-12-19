//go:generate mockgen -source=$GOFILE -destination=../utils/mockgen/$GOPACKAGE/mock_$GOFILE -package=$GOPACKAGE
package usecase

import (
	entity_book "my-app/domain/entity/book"
	"my-app/domain/repository"
)

type BookUsecase interface {
	GetByID(uint64) (*entity_book.Book, error)
	GetAll() ([]entity_book.Book, error)
	Create(entity_book.Book) (*entity_book.Book, error)
	Update(entity_book.Book) (*entity_book.Book, error)
	Delete(uint64) error
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

func (usecase bookUsecase) Create(book entity_book.Book) (*entity_book.Book, error) {

	return usecase.bookRepository.Create(book)
}

func (usecase bookUsecase) Update(book entity_book.Book) (*entity_book.Book, error) {

	return usecase.bookRepository.Update(book)
}

func (usecase bookUsecase) Delete(id uint64) error {

	return usecase.bookRepository.Delete(id)
}
