// store/store.go
package store

import (
    "database/sql"
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
