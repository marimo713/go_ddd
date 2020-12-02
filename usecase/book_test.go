package usecase

import (
	entity_book "my-app/domain/entity/book"
	"testing"

	"github.com/go-playground/assert"
)

func TestBook_GetBook_ReturnsBook(t *testing.T) {
	usecase := NewBookUsecase()

	expect := entity_book.NewBook(123, "978-4798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	book := usecase.GetBook(123)

	assert.Equal(t, expect, book)
}
