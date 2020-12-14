package handler

import (
	"errors"
	"log"
	entity_book "my-app/domain/entity/book"
	"my-app/interface/middleware"
	usecase_mock "my-app/utils/mockgen/usecase"
	test_util "my-app/utils/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupBookRouter(t *testing.T) (*gin.Engine, *usecase_mock.MockBookUsecase, func()) {
	ctrl := gomock.NewController(t)
	bu := usecase_mock.NewMockBookUsecase(ctrl)

	handler := NewBookHandler(bu)
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware())
	r.GET("/v1/books/:id", handler.GetByID)
	r.GET("/v1/books/", handler.GetAll)

	return r, bu, ctrl.Finish
}

func TestBook_GetByID_ReturnsBook(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	bu.EXPECT().GetByID(seed.ID()).Return(seed, nil)

	res := test_util.CallAPI(r, "GET", "/v1/books/123", nil)

	assert.Equal(t, 200, res.Code)
	expect := "{" +
		"\"id\":123," +
		"\"isbn\":\"9784798121963\"," +
		"\"title\":\"エリック・エヴァンスのドメイン駆動設計\"," +
		"\"author\":\"エリック・エヴァンス\"" +
		"}"
	assert.Equal(t, expect, res.Body.String())
}

func TestBook_GetByID_Response500WhenUsecaseReturnsError(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	bu.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("request failed"))

	res := test_util.CallAPI(r, "GET", "/v1/books/123", nil)

	assert.Equal(t, 500, res.Code)
}

func TestBook_GetAll_ReturnsBooks(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	seed1, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	seed2, err := entity_book.NewBookForRebuild(234, "1234567890123", "実践ドメイン駆動設計", "	ヴァーン・ヴァーノン")
	if err != nil {
		log.Fatal(err)
	}
	seeds := []entity_book.Book{*seed1, *seed2}
	bu.EXPECT().GetAll().Return(seeds, nil)

	res := test_util.CallAPI(r, "GET", "/v1/books/", nil)

	assert.Equal(t, 200, res.Code)
	expect := "[{" +
		"\"id\":123," +
		"\"isbn\":\"9784798121963\"," +
		"\"title\":\"エリック・エヴァンスのドメイン駆動設計\"," +
		"\"author\":\"エリック・エヴァンス\"" +
		"},{" +
		"\"id\":234," +
		"\"isbn\":\"1234567890123\"," +
		"\"title\":\"実践ドメイン駆動設計\"," +
		"\"author\":\"\\tヴァーン・ヴァーノン\"" +
		"}]"
	assert.Equal(t, expect, res.Body.String())
}

func TestBook_GetAll_ReturnsEmptyArrayWhenDataNotFound(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	bu.EXPECT().GetAll().Return([]entity_book.Book{}, nil)

	res := test_util.CallAPI(r, "GET", "/v1/books/", nil)

	assert.Equal(t, 200, res.Code)
	expect := "[]"
	assert.Equal(t, expect, res.Body.String())
}

func TestBook_GetAll_Response500WhenUsecaseReturnsError(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	bu.EXPECT().GetAll().Return(nil, errors.New("request failed"))

	res := test_util.CallAPI(r, "GET", "/v1/books/", nil)

	assert.Equal(t, 500, res.Code)
}
