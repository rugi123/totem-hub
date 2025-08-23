package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/pkg/database"
)

type MessageRepository struct {
	*PostgresRepository
}

func NewMessageRepository(db *database.Postgres) *MessageRepository {
	return &MessageRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *MessageRepository) GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, userID uuid.UUID) ([]entity.Message, error) {

	query := `
        SELECT m.id, m.chat_member_id, m.text, m.sent_at, m.edited_at 
        FROM messages m 
        JOIN chat_members cm ON m.chat_member_id = cm.id 
        WHERE cm.chat_id = $1 AND cm.user_id = $2
    `

	var messages []entity.Message

	rows, err := r.db.Pool.Query(ctx, query, chatID, userID)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var message entity.Message
		err := rows.Scan(
			&message.ID,
			&message.MemberID,
			&message.Text,
			&message.SentAt,
			&message.EditedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return messages, nil
}
func (r *MessageRepository) Save(ctx context.Context, msg *entity.Message) error {
	query := `
		INSERT INTO messages (id, chat_member_id, text, sent_at, edited_at)
		VALUES ($1, $2, $3, $4, $5)
		`
	_, err := r.db.Pool.Exec(ctx, query, msg.ID, msg.MemberID, msg.Text, msg.SentAt, msg.EditedAt)
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
