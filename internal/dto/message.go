package dto

import (
	"time"

	"github.com/google/uuid"
)

type SendMessage struct {
	MemberID uuid.UUID `json:"member_id"`
	Text     string    `json:"text"`
}

type BroadcastMessage struct {
	ID       uuid.UUID `json:"id"`
	MemberID uuid.UUID `json:"member_id"`
	Text     string    `json:"text"`
	SentAt   time.Time `json:"sent_at"`
}
