package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/pkg/database"
)

type DiologRepository struct {
	*PostgresRepository
}

func NewDiologRepository(db *database.Postgres) *DiologRepository {
	return &DiologRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *DiologRepository) Create(ctx context.Context, diolog entity.Diolog) error {
	query := `INSERT INTO groups (id, type, title, created_at, created_by, user1_id, user2_id)
		VALUES (@id, @type, @title, @created_at, @created_by, @user1_id, @user2_id)`

	args := pgx.NamedArgs{
		"id":         diolog.ID,
		"type":       diolog.Type,
		"title":      diolog.Title,
		"created_at": diolog.CreatedAt,
		"created_by": diolog.CreatedBy,
		"user1_id":   diolog.User1ID,
		"user2_id":   diolog.User2ID,
	}
	_, err := r.db.Pool.Exec(ctx, query, args)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}

func GetDiologsTx(ctx context.Context, tx pgx.Tx, ids []uuid.UUID) ([]entity.Diolog, error) {
	query := `SELECT id, title, type, created_at, created_by, user1_id, user2_id 
		FROM groups WHERE id = ANY($1)`
	rows, err := tx.Query(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	var diologs []entity.Diolog
	for rows.Next() {
		var diolog entity.Diolog
		err = rows.Scan(&diolog.ID, &diolog.Title, &diolog.Type,
			&diolog.CreatedAt, &diolog.CreatedBy, &diolog.User1ID, &diolog.User2ID)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil, fmt.Errorf("context canceled: %w", err)
			}
			return nil, fmt.Errorf("postgres error: %w", err)
		}
		diologs = append(diologs, diolog)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return diologs, nil
}
