package handler

import (
	"io"
	"log"
	entity_book "my-app/domain/entity/book"
	usecase_mock "my-app/utils/mockgen"
	"net/http"
	"net/http/httptest"
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
	r.GET("/v1/books/:id", handler.GetBook)

	return r, bu, ctrl.Finish
}

func TestBook_GetBook_ResponseBook(t *testing.T) {
	r, bu, cleanup := setupBookRouter(t)
	defer cleanup()

	seed := entity_book.NewBook(123, "978-4798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	bu.EXPECT().GetBook(uint64(123)).Return(seed)

	res := callAPI(r, "GET", "/v1/books/123", nil)

	assert.Equal(t, 200, res.Code)
	expect := "{" +
		"\"id\":123," +
		"\"isbn\":\"978-4798121963\"," +
		"\"title\":\"エリック・エヴァンスのドメイン駆動設計\"," +
		"\"author\":\"エリック・エヴァンス\"" +
		"}"
	assert.Equal(t, expect, res.Body.String())
}

func callAPI(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	router.ServeHTTP(w, req)
	return w
}
