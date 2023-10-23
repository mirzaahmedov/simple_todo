package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mirzaahmedov/simple_todo/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestTodoHandlers(t *testing.T) {
	r := NewTestHTTPRouter()

	t.Run("TodoCreateHandler", func(t *testing.T) {
		data := model.TodoCreateRequest{
			Title:     "todo title",
			Content:   "todo content",
			Completed: false,
		}
		req, err := NewTestRequest(http.MethodPost, "/todos", data)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Add("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		res := &SuccessResponse[*model.Todo]{}

		assert.Nil(t, json.NewDecoder(rec.Body).Decode(&res))
		assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)
		assert.NotZero(t, res.Data.ID)
		assert.Equal(t, data.Title, res.Data.Title)
		assert.Equal(t, data.Content, res.Data.Content)
		assert.Equal(t, data.Completed, res.Data.Completed)
	})
}
