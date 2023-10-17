package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	databaseURL string
	db          *sql.DB
	repository  struct {
		todo *TodoRepository
	}
}

func NewStore(databaseURL string) *PostgresStore {
	return &PostgresStore{
		databaseURL: databaseURL,
	}
}
func (s *PostgresStore) Open() error {
	db, err := sql.Open("postgres", s.databaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}
func (s *PostgresStore) Close() error {
	return s.db.Close()
}
