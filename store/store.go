// store/store.go
package store

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
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

func (s *Store) CreateBudget(month, category string, amount int) (string, error) {
    id := uuid.New().String() // generate a new UUID
    _, err := s.DB.Exec(
        "INSERT INTO budgets (id, month, category, budget) VALUES (?, ?, ?, ?)",
        id, month, category, amount,
    )
    if err != nil {
        return "", err
    }
    return id, nil
}
