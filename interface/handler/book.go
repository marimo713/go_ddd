package handler

import (
	"my-app/interface/response"
	"my-app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type bookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) BookHandler {
	return bookHandler{
		bookUsecase: bookUsecase,
	}
}

func (handler bookHandler) GetByID(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.Abort()
		return
	}
	book, err := handler.bookUsecase.GetByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.NewBookFromDomain(*book))
}

func (handler bookHandler) GetAll(c *gin.Context) {
	books, err := handler.bookUsecase.GetAll()
	if err != nil {
		c.Error(err)
		return
	}

	responseBooks := []response.Book{}
	for i := range books {
		responseBooks = append(responseBooks, response.NewBookFromDomain(books[i]))
	}
	c.JSON(200, responseBooks)
}
