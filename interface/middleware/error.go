package middleware

import (
	"log"
	myerror "my-app/error"
	"my-app/interface/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		ge, ok := err.Err.(myerror.GeneralError)
		if ok {
			httpStatus := myerror.GetHTTPStatus(ge.Code())
			response := response.ErrorResponse{
				Code:     ge.Code(),
				Messages: ge.Messages(),
			}
			c.AbortWithStatusJSON(httpStatus, response)
		}

		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
