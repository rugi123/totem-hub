package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *entity.Chat) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	Update(ctx context.Context, chat *entity.Chat) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Usecase struct {
	ChatRepo ChatRepository
	Config   config.Config
}

func NewChatUsecase(cfg config.Config, chatRepo ChatRepository) *Usecase {
	return &Usecase{
		Config:   cfg,
		ChatRepo: chatRepo,
	}
}
