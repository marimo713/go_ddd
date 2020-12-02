package handler

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupBookRouter(handler BookHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/v1/books/:id", handler.GetBook)

	return r
}

func TestBook_GetBook_ResponseBook(t *testing.T) {
	handler := NewBookHandler()
	r := setupBookRouter(handler)

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
