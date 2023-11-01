package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseSuccess struct {
	Data any `json:"data,omitempty"`
	Meta any `json:"meta,omitempty"`
}
type ResponseError struct {
	Error any `json:"error,omitempty"`
}

func JSON(c *gin.Context, status int, data any, meta any) {
	c.JSON(status, &ResponseSuccess{
		Data: data,
		Meta: meta,
	})
}
func Error(c *gin.Context, status int, err any) {
	c.JSON(status, &ResponseError{
		Error: err,
	})
}
