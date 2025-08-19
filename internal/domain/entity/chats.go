package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID         uuid.UUID
	Type       string
	Title      string
	CreatedAt  time.Time
	CreatedBy  uuid.UUID
	Attributes map[string]interface{}
}

func NewChat(chat_type string, title string, created_by uuid.UUID, attributes map[string]interface{}) Chat {
	return Chat{
		ID:         uuid.New(),
		Type:       chat_type,
		Title:      title,
		CreatedAt:  time.Now(),
		CreatedBy:  created_by,
		Attributes: attributes,
	}
}
