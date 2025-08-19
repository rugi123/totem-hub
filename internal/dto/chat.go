package dto

import "github.com/google/uuid"

type CreateChat struct {
	Type       string                 `json:"type"`
	Title      string                 `json:"title"`
	Attributes map[string]interface{} `json:"attributes"`
}

type ChannelAttributes struct {
	Description string `validate:"max=512"`
	IsPrivate   bool   `validate:"required,boolean"`
}
type DiologAttributes struct {
	User1ID uuid.UUID `validate:"required,uuid"`
	User2ID uuid.UUID `validate:"required,uuid"`
}
type GroupAttributes struct {
	IsPublic bool `vaildate:"required,boolean"`
}
