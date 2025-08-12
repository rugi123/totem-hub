package entity

import "github.com/google/uuid"

type Member struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	ChatID  uuid.UUID
	Role    string
	IsMuted bool
}
