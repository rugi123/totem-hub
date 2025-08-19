package dto

import "github.com/google/uuid"

type SendMessage struct {
	MemberID uuid.UUID `json:"member_id"`
	Text     string    `json:"text"`
}
