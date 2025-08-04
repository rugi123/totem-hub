package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	ChatMember ChatMember
	Text       string
	SentAt     time.Time
	EditedAt   time.Time
}
