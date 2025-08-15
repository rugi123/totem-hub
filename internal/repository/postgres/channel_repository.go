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

type ChannelRepository struct {
	*PostgresRepository
}

func NewChannelRepository(db *database.Postgres) *ChannelRepository {
	return &ChannelRepository{PostgresRepository: NewPostgresRepository(db)}
}
func (r *ChannelRepository) Create(ctx context.Context, channel entity.Channel) error {
	query := `INSERT INTO groups (id, type, title, created_at, created_by, description, is_private)
		VALUES (@id, @type, @title, @created_at, @created_by, @description, @is_private)`

	args := pgx.NamedArgs{
		"id":          channel.ID,
		"type":        channel.Type,
		"title":       channel.Title,
		"created_at":  channel.CreatedAt,
		"created_by":  channel.CreatedBy,
		"description": channel.Description,
		"is_private":  channel.IsPrivate,
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

func GetChannelsTx(ctx context.Context, tx pgx.Tx, ids []uuid.UUID) ([]entity.Channel, error) {
	query := `SELECT id, title, type, created_at, created_by, description, is_private 
		FROM groups WHERE id = ANY($1)`
	rows, err := tx.Query(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	var channels []entity.Channel
	for rows.Next() {
		var channel entity.Channel
		err = rows.Scan(&channel.ID, &channel.Title, &channel.Type,
			&channel.CreatedAt, &channel.CreatedBy, &channel.Description, &channel.IsPrivate)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil, fmt.Errorf("context canceled: %w", err)
			}
			return nil, fmt.Errorf("postgres error: %w", err)
		}
		channels = append(channels, channel)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return channels, nil
}
