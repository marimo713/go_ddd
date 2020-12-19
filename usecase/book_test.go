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

func TestBook_GetByID_ReturnsBookWhenRepositorySuccess(t *testing.T) {
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

func TestBook_GetAll_ReturnsBooksWhenRepositorySuccess(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	seed1, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	seed2, err := entity_book.NewBookForRebuild(234, "1234567890123", "実践ドメイン駆動設計", "	ヴァーン・ヴァーノン")
	if err != nil {
		log.Fatal(err)
	}
	expect := []entity_book.Book{*seed1, *seed2}
	mBookRepository.EXPECT().GetAll().Return(expect, nil)

	books, err := usecase.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expect, books)
}

func TestBook_GetAll_ReturnsErrorWhenRepositoryReturnsError(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	mBookRepository.EXPECT().GetAll().Return(nil, errors.New("repository error"))

	book, err := usecase.GetAll()

	assert.Error(t, err)
	assert.Nil(t, book)
}

func TestBook_Create_ReturnsBookWhenRepositorySuccess(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	expect, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	mBookRepository.EXPECT().Create(*expect).Return(expect, nil)

	book, err := usecase.Create(*expect)

	assert.NoError(t, err)
	assert.Equal(t, expect, book)
}

func TestBook_Create_ReturnsErrorWhenRepositoryReturnsError(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	mBookRepository.EXPECT().Create(gomock.Any()).Return(nil, errors.New("repository error"))

	book, err := usecase.Create(entity_book.Book{})

	assert.Error(t, err)
	assert.Nil(t, book)
}

func TestBook_Update_ReturnsBookWhenRepositorySuccess(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	expect, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	mBookRepository.EXPECT().Update(*expect).Return(expect, nil)

	book, err := usecase.Update(*expect)

	assert.NoError(t, err)
	assert.Equal(t, expect, book)
}

func TestBook_Update_ReturnsErrorWhenRepositoryReturnsError(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	mBookRepository.EXPECT().Update(gomock.Any()).Return(nil, errors.New("repository error"))

	book, err := usecase.Update(entity_book.Book{})

	assert.Error(t, err)
	assert.Nil(t, book)
}

func TestBook_Delete_NotReturnErrorWhenRepositorySuccess(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	mBookRepository.EXPECT().Delete(uint64(123)).Return(nil)

	err := usecase.Delete(123)

	assert.NoError(t, err)
}

func TestBook_Delete_ReturnsErrorWhenRepositoryReturnsError(t *testing.T) {
	usecase, mBookRepository, cleanup := newBookUsecaseWithMock(t)
	defer cleanup()

	mBookRepository.EXPECT().Delete(gomock.Any()).Return(errors.New("repository error"))

	err := usecase.Delete(123)

	assert.Error(t, err)
}
