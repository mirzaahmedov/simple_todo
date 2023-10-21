package mock

import (
	_ "github.com/lib/pq"
)

type MockStore struct {
	repository struct {
		todo *TodoRepository
	}
}

func NewStore() *MockStore {
	return &MockStore{}
}
