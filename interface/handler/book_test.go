package handler

import (
	"errors"
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

	return r, bu, ctrl.Finish
}

func TestBook_GetByID_ReturnsBook(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	seed := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	bu.EXPECT().GetByID(seed.ID()).Return(&seed, nil)

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
