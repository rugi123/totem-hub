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
