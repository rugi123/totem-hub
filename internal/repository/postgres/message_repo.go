package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type MessageRepo struct {
	pool *pgxpool.Pool
}

func NewMessageRepo(ctx context.Context, cfg config.Postgres) (*MessageRepo, error) {
	conn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)
	pool, err := pgxpool.New(ctx, conn)
	return &MessageRepo{
		pool: pool,
	}, err
}

func (r *MessageRepo) GetMessages(ctx context.Context, id uuid.UUID) (*entity.Message, error) {
	query := `
		SELECT m.id, m.text, m.sent_at, m.edited_at, cm.*
		FROM messages m
		JOIN chat_members cm ON m.chat_member_id = cm.id
		WHERE cm.id = $1
		`
	var msg entity.Message
	err := r.pool.QueryRow(ctx, query, id).Scan(
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
func (r *MessageRepo) GetChatMessages(ctx context.Context, chat_id uuid.UUID) (*entity.Message, error) {
	query := `
		SELECT m.id, m.text, m.sent_at, m.edited_at, cm.*
		FROM messages m
		JOIN chat_members cm ON m.chat_member_id = cm.id
		WHERE cm.chat_id = $1
		`
	var msg entity.Message
	err := r.pool.QueryRow(ctx, query, chat_id).Scan(
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
func (r *MessageRepo) GetMemberMessages(ctx context.Context, member_id uuid.UUID) (*entity.Message, error) {
	query := `
		SELECT m.id, m.text, m.sent_at, m.edited_at, cm.*
		FROM messages m
		JOIN chat_members cm ON m.chat_member_id = cm.id
		WHERE cm.id = $1
		`
	var msg entity.Message
	err := r.pool.QueryRow(ctx, query, member_id).Scan(
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
func (r *MessageRepo) CreateMessage(ctx context.Context, msg *entity.Message) error {
	query := `
		INSERT INTO messages (id, chat_member_id, text, sent_at, edited_at)
		VALUES ($1, $2, $3, $4, $5)
		`
	_, err := r.pool.Exec(ctx, query, msg.ID, msg.ChatMember.ID, msg.Text, msg.SentAt, msg.EditedAt)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *MessageRepo) UpdateMessage(ctx context.Context, msg *entity.Message) error {
	return nil
}
func (r *MessageRepo) DeleteMessage(ctx context.Context, id int) error {
	return nil
}
