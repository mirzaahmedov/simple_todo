package postgres

import (
	"github.com/mirzaahmedov/simple_todo/internal/model"
	"github.com/mirzaahmedov/simple_todo/internal/store"
)

type TodoRepository struct {
	store *PostgresStore
}

func (s *PostgresStore) Todo() store.TodoRepository {
	if s.repository.todo == nil {
		s.repository.todo = &TodoRepository{
			store: s,
		}
	}

	return s.repository.todo
}

func (t *TodoRepository) Create(values *model.Todo) (*model.Todo, error) {
	todo := &model.Todo{}

	err := t.store.db.QueryRow(
		`
                  INSERT INTO todos (title, content, completed) VALUES ($1, $2, $3);
                `,
		&values.Title,
		&values.Content,
		&values.Completed,
	).Scan(&todo)
	if err != nil {
		return nil, err
	}

	return todo, err
}
func (t *TodoRepository) GetAll() ([]model.Todo, error) {
	todos := []model.Todo{}

	rows, err := t.store.db.Query(`SELECT id, title, content, completed, create_date, update_date FROM todos;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		todo := model.Todo{}
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.Completed,
			&todo.CreateDate,
			&todo.UpdateDate,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
func (t *TodoRepository) GetByID(id string) (*model.Todo, error) {
	todo := &model.Todo{}

	err := t.store.db.QueryRow(
		`SELECT id, title, content, completed, create_date, update_date FROM todos WHERE id = $1`,
		id,
	).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.Completed,
		&todo.CreateDate,
		&todo.UpdateDate,
	)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
func (t *TodoRepository) Update(id string, values *model.Todo) (*model.Todo, error) {
	todo := model.Todo{}

	err := t.store.db.QueryRow(
		`
                  UPDATE todos 
                  SET title = $2, content = $3, completed = $4 
                  WHERE id = $1
                  RETURNING id, title, content, completed, create_date, update_date
                `,
		id,
		values.Title,
		values.Content,
		values.Completed,
	).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.Completed,
		&todo.CreateDate,
		&todo.UpdateDate,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
func (t *TodoRepository) Delete(id string) error {
	_, err := t.store.db.Exec(
		`DELETE FROM todos WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
