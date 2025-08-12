package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/pkg/database"
)

type MessageRepository struct {
	*PostgresRepository
}

func NewMessageRepository(db *database.Postgres) *MessageRepository {
	return &MessageRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *MessageRepository) GetAll(ctx context.Context, msg_id uuid.UUID) (*entity.Message, error) {
	query := `
		SELECT m.id, m.text, m.sent_at, m.edited_at, cm.*
		FROM messages m
		JOIN chat_members cm ON m.chat_member_id = cm.id
		WHERE cm.id = $1
		`
	var msg entity.Message
	err := r.db.Pool.QueryRow(ctx, query, msg_id).Scan(
		&msg.ID, &msg.Text, &msg.SentAt, &msg.EditedAt,
		&msg.ChatMember.ID, &msg.ChatMember.UserID, &msg.ChatMember.ChatID, &msg.ChatMember.Role)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("context canceled: %w", err)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return &msg, nil
}
func (r *MessageRepository) Create(ctx context.Context, msg *entity.Message) error {
	query := `
		INSERT INTO messages (id, chat_member_id, text, sent_at, edited_at)
		VALUES ($1, $2, $3, $4, $5)
		`
	_, err := r.db.Pool.Exec(ctx, query, msg.ID, msg.ChatMember.ID, msg.Text, msg.SentAt, msg.EditedAt)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *MessageRepository) Update(ctx context.Context, msg *entity.Message) error {
	return nil
}
func (r *MessageRepository) Delete(ctx context.Context, msg_id uuid.UUID) error {
	return nil
}
