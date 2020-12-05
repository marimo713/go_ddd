package test_util

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CallAPI(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	router.ServeHTTP(w, req)
	return w
}
