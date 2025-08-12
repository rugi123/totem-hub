package postgres

import (
	"github.com/rugi123/chirp/pkg/database"
)

type PostgresRepository struct {
	db *database.Postgres
}

func NewPostgresRepository(db *database.Postgres) *PostgresRepository {
	return &PostgresRepository{db: db}
}

//общие методы для всех репозиториев
