// +build wireinject
package main

import (
	"my-app/interface/handler"
	"my-app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeServer() *gin.Engine {
	wire.Build(
		handler.NewBookHandler,
		usecase.NewBookUsecase,
		newServer,
	)
	return nil
}
