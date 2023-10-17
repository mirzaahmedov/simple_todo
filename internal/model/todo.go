package model

import (
	"time"
)

type Todo struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Completed  bool      `json:"completed"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}

type TodoCreateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
type TodoUpdateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
