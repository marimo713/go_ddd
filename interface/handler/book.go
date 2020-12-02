package handler

import (
	entity_book "my-app/domain/entity/book"
	"my-app/interface/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	GetBook(c *gin.Context)
}

type bookHandler struct {
}

func NewBookHandler() BookHandler {
	return bookHandler{}
}

func (handler bookHandler) GetBook(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.Abort()
		return
	}
	book := entity_book.NewBook(id, "978-4798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")

	c.JSON(200, response.NewBookFromDomain(book))
}
