package router

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/simple_todo/internal/model"
)

func (r *HTTPRouter) handleTodoCreate(c *gin.Context) {
	req := model.TodoCreateRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		r.logger.Debug("error binding request body", slog.String("error", err.Error()))
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := r.store.Todo().Create(&model.Todo{
		Title:     req.Title,
		Content:   req.Content,
		Completed: req.Completed,
	})
	if err != nil {
		r.logger.Debug("can not save todo in database", slog.String("error", err.Error()))
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(c, http.StatusCreated, todo, nil)
}
func (r *HTTPRouter) handleTodoGetAll(c *gin.Context) {
	todos, err := r.store.Todo().GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}
func (r *HTTPRouter) handleTodoGetByID(c *gin.Context) {
	id := c.Param("id")

	todo, err := r.store.Todo().GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}
func (r *HTTPRouter) handleTodoUpdate(c *gin.Context) {
	id := c.Param("id")
	req := model.TodoUpdateRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		r.logger.Debug("can not bind request body", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := r.store.Todo().Update(id, &model.Todo{
		Title:     req.Title,
		Content:   req.Content,
		Completed: req.Completed,
	})
	if err != nil {
		r.logger.Debug("can not save todo in the database", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}
func (r *HTTPRouter) handleTodoDelete(c *gin.Context) {
	id := c.Param("id")

	err := r.store.Todo().Delete(id)
	if err != nil {
		r.logger.Debug("can not delete todo from todos", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
