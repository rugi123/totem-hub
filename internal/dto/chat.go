package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type GetChatResponse struct {
	ID        uuid.UUID
	Name      string
	Members   []entity.ChatMember
	CreatedAt time.Time
	CreatedBy uuid.UUID
}

//create chat

type CreateChatRequest struct {
}

type CreateChatResponse struct {
	ID uuid.UUID
}
