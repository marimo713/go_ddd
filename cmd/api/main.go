package main

import (
	"log"
	"my-app/config"
	"my-app/interface/handler"
	"my-app/interface/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := config.NewConfig("./config/env/", "develop")
	if err != nil {
		log.Fatal(err)
	}
	r, cleanup, err := initializeServer(config)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	r.Run()
}

func newServer(
	bookHandler handler.BookHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware())
	r.GET("/v1/books/:id", bookHandler.GetByID)
	r.GET("/v1/books/", bookHandler.GetAll)

	return r
}
