package repository

import (
	"database/sql"
	"linux-docker-web-gui/pkg/models"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) GetTest() (*models.Test, error) {
	return &models.Test{ID: 0}, nil
}
