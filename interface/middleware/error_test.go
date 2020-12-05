package middleware

import (
	"errors"
	myerror "my-app/error"
	test_util "my-app/utils/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestErrorMiddleware_ResponseErrorErrorIsSetted(t *testing.T) {
	cases := []struct {
		name         string
		err          error
		expectStatus int
		expectBody   string
	}{
		{
			name:         "ReturnsErrorResponseWhenGeneralErrorIsSetted",
			err:          myerror.NewNotFoundError(errors.New("data not found"), "data not found"),
			expectStatus: 404,
			expectBody: "{" +
				"\"Code\":\"1\"," +
				"\"Messages\":[\"data not found\"]" +
				"}",
		},
		{
			name:         "Response500WhenNonGeneralErrorIsSetted",
			err:          errors.New("data not found"),
			expectStatus: 500,
			expectBody:   "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := gin.Default()
			r.Use(ErrorMiddleware())
			r.GET("/test", func(ctx *gin.Context) {
				ctx.Error(c.err)
				return
			})

			res := test_util.CallAPI(r, "GET", "/test", nil)
			assert.Equal(t, c.expectStatus, res.Code)
			assert.Equal(t, c.expectBody, res.Body.String())
		})
	}
}

func TestErrorMiddleware_DoesNotDoAnythingWhenErrorIsNotSetted(t *testing.T) {
	r := gin.Default()
	r.Use(ErrorMiddleware())
	r.GET("/test", func(ctx *gin.Context) {
		return
	})

	res := test_util.CallAPI(r, "GET", "/test", nil)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "", res.Body.String())
}
