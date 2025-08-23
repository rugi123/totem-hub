package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID
	MemberID uuid.UUID
	Text     string
	SentAt   time.Time
	EditedAt sql.NullTime
}

func NewMessage(memberID uuid.UUID, text string) *Message {
	return &Message{
		ID:       uuid.New(),
		MemberID: memberID,
		Text:     text,
		SentAt:   time.Now(),
		EditedAt: sql.NullTime{},
	}
}
