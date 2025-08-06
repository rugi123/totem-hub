package chat

import (
	"context"

	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type ChatRepository interface {
	GetChat(ctx context.Context, id int) (*entity.Chat, error)
	CreateChat(ctx context.Context, chat *entity.Chat) error
	UpdateChat(ctx context.Context, chat *entity.Chat) error
	DeleteChat(ctx context.Context, id int) error
}
type MemberRepository interface {
	GetMember(ctx context.Context, id int) (*entity.ChatMember, error)
	CreateMember(ctx context.Context, member *entity.ChatMember) error
	UpdateMember(ctx context.Context, member *entity.ChatMember) error
	DeleteMember(ctx context.Context, id int) error
}

type Usecase struct {
	ChatRepo   ChatRepository
	MemberRepo MemberRepository
	Config     config.Config
}

func NewChatUsecase(cfg config.Config, chatRepo ChatRepository, memberRepo MemberRepository) *Usecase {
	return &Usecase{
		Config:     cfg,
		ChatRepo:   chatRepo,
		MemberRepo: memberRepo,
	}
}
