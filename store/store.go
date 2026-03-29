// store/store.go
package store

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	DB *sql.DB
}

func New(dbPath string) (*Store, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

func (s *Store) CreateBudget(month time.Time, categoryId int, budget float64) (string, error) {
	id := uuid.New().String() // generate a new UUID
	_, err := s.DB.Exec(
		"INSERT INTO Budget (id, month, category_id, budget) VALUES (?, ?, ?, ?)",
		id, month, categoryId, budget,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}
