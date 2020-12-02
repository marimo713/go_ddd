package main

import (
	"my-app/interface/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	bookHandler := handler.NewBookHandler()
	r := gin.Default()
	r.GET("/v1/book/:id", bookHandler.GetBook)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
