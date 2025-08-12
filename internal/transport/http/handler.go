package http

import (
	"github.com/rugi123/chirp/internal/usecase/auth"
	"github.com/rugi123/chirp/internal/usecase/chat"
	"github.com/rugi123/chirp/internal/usecase/member"
	"github.com/rugi123/chirp/internal/usecase/message"
)

type Handler struct {
	authUC   auth.Usecase
	chatUC   chat.Usecase
	memberUC member.Usecase
	msgUC    message.Usecase
}

func NewHanlder(authUC auth.Usecase, chatUC chat.Usecase, memberUC member.Usecase, msgUC message.Usecase) *Handler {
	return &Handler{
		authUC:   authUC,
		chatUC:   chatUC,
		memberUC: memberUC,
		msgUC:    msgUC,
	}
}
