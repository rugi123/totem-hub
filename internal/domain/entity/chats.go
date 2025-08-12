package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        uuid.UUID
	Title     string
	Type      string
	CreatedAt time.Time
	CreatedBy uuid.UUID ///user id
}

type Diolog struct {
	Chat
	User1ID uuid.UUID
	User2ID uuid.UUID
}

type Group struct {
	Chat
	IsPublic bool
}

type Channel struct {
	Chat
	AdminID   uuid.UUID
	IsPrivate bool
}
