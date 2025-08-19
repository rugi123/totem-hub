package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rugi123/chirp/pkg/database"
)

type PostgresRepository struct {
	db *database.Postgres
}

func NewPostgresRepository(db *database.Postgres) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// общие методы для всех репозиториев
func (r *PostgresRepository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(tx)
	return err
}
