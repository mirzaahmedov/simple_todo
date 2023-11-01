package router

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mirzaahmedov/simple_todo/internal/http/response"
	"github.com/mirzaahmedov/simple_todo/internal/model"
)

func (r *HTTPRouter) handleTodoCreate(c *gin.Context) {
	req := model.TodoCreateRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			response.Error(c, 400, gin.H{
				"message": e.Field + " must be a " + e.Type.Name(),
			})
			return
		}
		response.Error(c, 400, err.Error())
		return
	}

	err = validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.Content, validation.Required, validation.Length(2, 50)),
	)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	todo, err := r.store.Todo().Create(&model.Todo{
		Title:     req.Title,
		Content:   req.Content,
		Completed: req.Completed,
	})
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.JSON(c, 201, todo, nil)
}
func (r *HTTPRouter) handleTodoGetAll(c *gin.Context) {
	todos, err := r.store.Todo().GetAll()
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.JSON(c, 200, todos, nil)
}
func (r *HTTPRouter) handleTodoGetByID(c *gin.Context) {
	id := c.Param("id")

	todo, err := r.store.Todo().GetByID(id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.JSON(c, 200, todo, nil)
}
func (r *HTTPRouter) handleTodoUpdate(c *gin.Context) {
	id := c.Param("id")
	req := model.TodoUpdateRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); !ok {
			response.Error(c, 400, gin.H{
				"message": e.Field + " " + e.Type.Name(),
			})
			return
		}
		response.Error(c, 400, err.Error())
		return
	}

	err = validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.Content, validation.Required, validation.Length(2, 50)),
	)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	todo, err := r.store.Todo().Update(id, &model.Todo{
		Title:     req.Title,
		Content:   req.Content,
		Completed: req.Completed,
	})
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.JSON(c, 200, todo, nil)
}
func (r *HTTPRouter) handleTodoDelete(c *gin.Context) {
	id := c.Param("id")

	err := r.store.Todo().Delete(id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.JSON(c, 200, true, nil)
}
