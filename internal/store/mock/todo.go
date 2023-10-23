package mock

import (
	"time"

	"github.com/mirzaahmedov/simple_todo/internal/generate"
	"github.com/mirzaahmedov/simple_todo/internal/model"
	"github.com/mirzaahmedov/simple_todo/internal/store"
)

type TodoRepository struct {
	store    *MockStore
	todos    []model.Todo
	sequence int
}

func (s *MockStore) Todo() store.TodoRepository {
	if s.repository.todo == nil {
		s.repository.todo = &TodoRepository{
			store: s,
			todos: []model.Todo{},
		}
	}

	return s.repository.todo
}

func (r *TodoRepository) Create(values *model.Todo) (*model.Todo, error) {
	todo := model.Todo{
		ID:         generate.UniqueID(),
		Title:      values.Title,
		Content:    values.Content,
		Completed:  values.Completed,
		CreateDate: time.Now(),
	}

	r.todos = append(r.todos, todo)

	return &todo, nil
}
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	return r.todos, nil
}
func (r *TodoRepository) GetByID(id string) (*model.Todo, error) {
	for _, t := range r.todos {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, store.ErrNotFound
}
func (r *TodoRepository) Update(id string, values *model.Todo) (*model.Todo, error) {
	for _, t := range r.todos {
		if t.ID == id {
			t.Title = values.Title
			t.Content = values.Content
			t.Completed = values.Completed

			return &t, nil
		}
	}

	return nil, store.ErrNotFound
}
func (r *TodoRepository) Delete(id string) error {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}

	return store.ErrNotFound
}
