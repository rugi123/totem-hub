package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/pkg/database"
)

type ChatRepository struct {
	*PostgresRepository
}

func NewChatRepository(db *database.Postgres) *ChatRepository {
	return &ChatRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *ChatRepository) GetAll(ctx context.Context, user_id uuid.UUID) (*[]entity.Chat, error) {
	return nil, nil
}
func (r *ChatRepository) Create(ctx context.Context, chat *entity.Chat) error {
	query := `
		INSERT INTO chats (id, is_private, created_at, created_by)
		VALUES($1, $2, $3, $4)
		`
	_, err := r.db.Pool.Exec(ctx, query, chat.ID, chat.CreatedAt, chat.CreatedBy)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *ChatRepository) Update(ctx context.Context, chat *entity.Chat) error {
	return nil
}
func (r *ChatRepository) Delete(ctx context.Context, chat_id uuid.UUID) error {
	return nil
}
