package store

import (
	"database/sql"
)

type SqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &SqlStore{
		db: db,
	}
}
