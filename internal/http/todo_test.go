package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mirzaahmedov/simple_todo/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestTodoHandlers(t *testing.T) {
	id := ""
	router := MakeTestHTTPRouter()

	t.Run("TodoCreateHandler", func(t *testing.T) {
		todo := model.TodoCreateRequest{
			Title:     "todo title",
			Content:   "todo content",
			Completed: false,
		}

		r, err := MakeTestRequest(http.MethodPost, "/todos", todo)
		assert.Nil(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		res, err := ParseJSON[*model.Todo](w.Body)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
		assert.Equal(t, todo.Title, res.Data.Title)
		assert.Equal(t, todo.Content, res.Data.Content)
		assert.Equal(t, todo.Completed, res.Data.Completed)

		assert.NotZero(t, res.Data.ID)
		assert.NotZero(t, res.Data.CreateDate)

		id = res.Data.ID
	})
	t.Run("TodoGetByIDHandler", func(t *testing.T) {
		r, err := MakeTestRequest(http.MethodGet, "/todos/"+id, nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		res, err := ParseJSON[*model.Todo](w.Body)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.NotZero(t, res.Data.Title)
	})
}
