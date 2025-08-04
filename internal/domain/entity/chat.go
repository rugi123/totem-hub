package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        uuid.UUID
	Name      string
	IsPrivate bool
	CreatedAt time.Time
	CreatedBy uuid.UUID ///user id
}

type ChatMember struct {
	ID     uuid.UUID
	UserID uuid.UUID
	ChatID uuid.UUID
	Role   string
}
