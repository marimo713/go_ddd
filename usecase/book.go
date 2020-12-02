//go:generate mockgen -source=$GOFILE -destination=../utils/mockgen/mock_$GOFILE -package=$GOPACKAGE
package usecase

import entity_book "my-app/domain/entity/book"

type BookUsecase interface {
	GetBook(uint64) entity_book.Book
}

type bookUsecase struct {
}

func NewBookUsecase() BookUsecase {
	return bookUsecase{}
}

func (usecase bookUsecase) GetBook(id uint64) entity_book.Book {

	return entity_book.NewBook(123, "978-4798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
}
