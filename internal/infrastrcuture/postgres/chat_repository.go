package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/pkg/database"
)

type ChatRepository struct {
	*PostgresRepository
}

func NewChatRepository(db *database.Postgres) *ChatRepository {
	return &ChatRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *ChatRepository) Create(ctx context.Context, chat entity.Chat) error {
	query := `
        INSERT INTO chats (id, type, title, created_at, created_by, attributes)
        VALUES (@id, @type, @title, @created_at, @created_by, @attributes)
    `

	args := pgx.NamedArgs{
		"id":         chat.ID,
		"type":       chat.Type,
		"title":      chat.Title,
		"created_at": chat.CreatedAt,
		"created_by": chat.CreatedBy,
		"attributes": chat.Attributes,
	}

	_, err := r.db.Pool.Exec(ctx, query, args)
	return err
}

func (r *ChatRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Chat, error) {
	query := `
        SELECT id, type, title, created_at, created_by, attributes
        FROM chats WHERE id = $1
    `

	var chat entity.Chat
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&chat.ID,
		&chat.Type,
		&chat.Title,
		&chat.CreatedAt,
		&chat.CreatedBy,
		&chat.Attributes, // pgx автоматически парсит JSONB в map
	)

	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (r *ChatRepository) GetByIDs(ctx context.Context, id []uuid.UUID) ([]entity.Chat, error) {
	query := `
        SELECT id, type, title, created_at, created_by, attributes
        FROM chats WHERE id = ANY($1)
    `

	var chats []entity.Chat
	rows, err := r.db.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var chat entity.Chat
		err := rows.Scan(
			&chat.ID,
			&chat.Type,
			&chat.Title,
			&chat.CreatedAt,
			&chat.CreatedBy,
			&chat.Attributes)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return chats, nil
}
