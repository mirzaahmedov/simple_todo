package store

import (
	"github.com/mirzaahmedov/simple_todo/internal/model"
)

type Store interface {
	Todo() TodoRepository
}

type TodoRepository interface {
	Create(*model.Todo) (*model.Todo, error)
	GetAll() ([]model.Todo, error)
	GetByID(string) (*model.Todo, error)
	Update(string, *model.Todo) (*model.Todo, error)
	Delete(string) error
}
