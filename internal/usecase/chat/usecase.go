package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type ChatRepository interface {
	GetAllUserChats(ctx context.Context, user_id uuid.UUID) ([]entity.Chat, error)
	CreateChat(ctx context.Context, chat *entity.Chat) error
	UpdateChat(ctx context.Context, chat *entity.Chat) error
	DeleteChat(ctx context.Context, chat_id uuid.UUID) error
}
type MemberRepository interface {
	GetMemberIDs(ctx context.Context, user_id uuid.UUID) ([]entity.ChatMember, error)
	CreateMember(ctx context.Context, member *entity.ChatMember) error
	UpdateMember(ctx context.Context, member *entity.ChatMember) error
	DeleteMember(ctx context.Context, member_id uuid.UUID) error
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
