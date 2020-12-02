package main

import (
	"my-app/interface/handler"
	"my-app/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	bookUsecase := usecase.NewBookUsecase()
	bookHandler := handler.NewBookHandler(bookUsecase)
	r := gin.Default()
	r.GET("/v1/book/:id", bookHandler.GetBook)
	r.Run()
}
