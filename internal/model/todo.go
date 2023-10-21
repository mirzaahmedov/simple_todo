package model

import (
	"time"
)

type Todo struct {
	ID         string     `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	Content    string     `json:"content,omitempty"`
	Completed  bool       `json:"completed"`
	UpdateDate *time.Time `json:"update_date,omitempty"`
	CreateDate time.Time  `json:"create_date,omitempty"`
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
