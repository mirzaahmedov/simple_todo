package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/simple_todo/internal/model"
)

type Model interface {
	*model.Todo
}

type SuccessResponse[T Model] struct {
	Data T   `json:"data"`
	Meta any `json:"meta"`
}
type ErrorResponse struct {
	Error any `json:"error"`
}

func respondJSON[T Model](c *gin.Context, status int, data T, meta any) {
	c.JSON(status, &SuccessResponse[T]{
		Data: data,
		Meta: meta,
	})
}
func respondError(c *gin.Context, status int, error any) {
	c.JSON(status, &ErrorResponse{
		Error: error,
	})
}
