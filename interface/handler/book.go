package handler

import (
	"my-app/interface/response"
	"my-app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	GetBook(c *gin.Context)
}

type bookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) BookHandler {
	return bookHandler{
		bookUsecase: bookUsecase,
	}
}

func (handler bookHandler) GetBook(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.Abort()
		return
	}
	book := handler.bookUsecase.GetBook(id)

	c.JSON(200, response.NewBookFromDomain(book))
}
