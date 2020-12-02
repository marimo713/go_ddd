package main

import (
	"my-app/interface/handler"
	"my-app/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := initializeServer()
	r.Run()
}

func newServer(
	bookHandler handler.BookHandler,
	bookUsecase usecase.BookUsecase,
) *gin.Engine {
	r := gin.Default()
	r.GET("/v1/book/:id", bookHandler.GetBook)

	return r
}
