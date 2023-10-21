package router

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mirzaahmedov/simple_todo/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestTodoHandlers(t *testing.T) {
	router := NewTestHTTPStore()

	t.Run("TodoCreateHandler", func(t *testing.T) {
		payload := model.TodoCreateRequest{
			Title:     "todo title",
			Content:   "todo content",
			Completed: false,
		}
		r, err := NewTestRequest(http.MethodPost, "/todos", payload)
		if err != nil {
			t.Fatal(err)
		}

		r.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		body := model.Todo{}
		bodyBytes, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Fatal(err)
		}
		json.Unmarshal(bodyBytes, &body)

		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
		assert.Equal(t, payload.Title, body.Title)
		assert.Equal(t, payload.Content, body.Content)
		assert.Equal(t, payload.Completed, body.Completed)
	})
}
