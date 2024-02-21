package repository

import (
	"awesome-backend-golang/internal/adapter/repository/sqlmap"
	"database/sql"
)

type Repository struct {
	db  *sql.DB
	sql *sqlmap.Map
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db, sql: sqlmap.New()}
}
