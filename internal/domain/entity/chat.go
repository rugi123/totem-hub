package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	CreatedBy uuid.UUID
}

type ChatMember struct {
	UserID uuid.UUID
	ChatID uuid.UUID
	Role   string
}
