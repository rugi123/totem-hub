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

type GroupRepository struct {
	*PostgresRepository
}

func NewGroupRepository(db *database.Postgres) *GroupRepository {
	return &GroupRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *GroupRepository) Create(ctx context.Context, group entity.Group) error {
	query := `INSERT INTO groups (id, type, title, created_at, created_by, is_public)
		VALUES (@id, @type, @title, @created_at, @created_by, @is_public)`

	args := pgx.NamedArgs{
		"id":         group.ID,
		"type":       group.Type,
		"title":      group.Title,
		"created_at": group.CreatedAt,
		"created_by": group.CreatedBy,
		"is_public":  group.IsPublic,
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

func GetGroupsTx(ctx context.Context, tx pgx.Tx, ids []uuid.UUID) ([]entity.Group, error) {
	query := `SELECT id, title, type, created_at, created_by, is_public 
		FROM groups WHERE id = ANY($1)`
	rows, err := tx.Query(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	var groups []entity.Group
	for rows.Next() {
		var group entity.Group
		err = rows.Scan(&group.ID, &group.Title, &group.Type,
			&group.CreatedAt, &group.CreatedBy, &group.IsPublic)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil, fmt.Errorf("context canceled: %w", err)
			}
			return nil, fmt.Errorf("postgres error: %w", err)
		}
		groups = append(groups, group)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return groups, nil
}
