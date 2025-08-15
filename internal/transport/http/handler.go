package http

import (
	"github.com/rugi123/chirp/internal/usecase/auth"
	"github.com/rugi123/chirp/internal/usecase/channel"
	"github.com/rugi123/chirp/internal/usecase/chat"
	"github.com/rugi123/chirp/internal/usecase/diolog"
	"github.com/rugi123/chirp/internal/usecase/group"
	"github.com/rugi123/chirp/internal/usecase/member"
	"github.com/rugi123/chirp/internal/usecase/message"
)

type Handler struct {
	authUC    auth.Usecase
	chatUC    chat.Usecase
	channelUC channel.Usecase
	diologUC  diolog.Usecase
	groupUC   group.Usecase
	memberUC  member.Usecase
	msgUC     message.Usecase
}

func NewHanlder(authUC auth.Usecase, chatUC chat.Usecase, channelUC channel.Usecase,
	diologUC diolog.Usecase, groupUC group.Usecase, memberUC member.Usecase, msgUC message.Usecase) *Handler {
	return &Handler{
		authUC:    authUC,
		chatUC:    chatUC,
		channelUC: channelUC,
		diologUC:  diologUC,
		groupUC:   groupUC,
		memberUC:  memberUC,
		msgUC:     msgUC,
	}
}
