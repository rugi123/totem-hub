package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type MessageRepository interface {
	GetAll(ctx context.Context, msg_id uuid.UUID) (*entity.Message, error)
	Create(ctx context.Context, msg *entity.Message) error
	Update(ctx context.Context, msg *entity.Message) error
	Delete(ctx context.Context, msg_id uuid.UUID) error
}

type Usecase struct {
	MessageRepo MessageRepository
	Config      config.Config
}

func NewMessageUsecase(cfg config.Config, repo MessageRepository) *Usecase {
	return &Usecase{
		Config:      cfg,
		MessageRepo: repo,
	}
}
