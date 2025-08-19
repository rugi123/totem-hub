package http

import (
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/usecase/chat"
	"github.com/rugi123/chirp/internal/usecase/member"
	"github.com/rugi123/chirp/internal/usecase/message"
	"github.com/rugi123/chirp/internal/usecase/user"
)

type Server struct {
	config   config.App
	userUC   user.Usecase
	chatUC   chat.Usecase
	memberUC member.Usecase
	msgUC    message.Usecase
}

func NewServer(cfg config.App, userUC user.Usecase, chatUC chat.Usecase, memberUC member.Usecase, msgUC message.Usecase) *Server {
	return &Server{
		config:   cfg,
		userUC:   userUC,
		chatUC:   chatUC,
		memberUC: memberUC,
		msgUC:    msgUC,
	}
}
