package usecase

import (
	"errors"
	"log"
	entity_book "my-app/domain/entity/book"
	repository_mock "my-app/utils/mockgen/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func newBookUsecaseWithMock(t *testing.T) (BookUsecase, *repository_mock.MockBookRepository, func()) {
	ctrl := gomock.NewController(t)
	br := repository_mock.NewMockBookRepository(ctrl)

	usecase := NewBookUsecase(br)
	return usecase, br, ctrl.Finish
}

func TestBook_GetByID_ReturnsBook(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	expect, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	mBookRepository.EXPECT().GetByID(expect.ID()).Return(expect, nil)

	book, err := usecase.GetByID(expect.ID())

	assert.NoError(t, err)
	assert.Equal(t, expect, book)
}

func TestBook_GetByID_ReturnsErrorWhenRepositoryReturnsError(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	mBookRepository.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("repository error"))

	book, err := usecase.GetByID(123)

	assert.Error(t, err)
	assert.Nil(t, book)

}
