package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type GetChatResponse struct {
	ID        uuid.UUID           `json:"id"`
	Name      string              `json:"name"`
	Members   []entity.ChatMember `json:"chat_members"`
	CreatedAt time.Time           `json:"created_at"`
	CreatedBy uuid.UUID           `json:"created_by"`
}

type CreateChatRequest struct {
	Name      string    `json:"name" validate:"max=32"`
	CreatedBy uuid.UUID `json:"created_by"`
}
