// +build wireinject

package main

import (
	"my-app/config"
	"my-app/infrastructure/mysql"
	"my-app/interface/handler"
	"my-app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeServer(conf config.Config) (*gin.Engine, func(), error) {
	wire.Build(
		handler.NewBookHandler,
		usecase.NewBookUsecase,
		mysql.NewBookRepository,
		mysql.Open,
		wire.FieldsOf(&conf, "DB"),
		newServer,
	)
	return nil, nil, nil
}
