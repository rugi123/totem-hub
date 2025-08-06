package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type Repository interface {
	GetMessages(ctx context.Context, id uuid.UUID) (*entity.Message, error)
	CreateMessage(ctx context.Context, msg *entity.Message) error
	UpdateMessage(ctx context.Context, msg *entity.Message) error
	DeleteMessage(ctx context.Context, id int) error
}

type Usecase struct {
	Repo   Repository
	Config config.Config
}

func NewMessageUsecase(cfg config.Config, repo Repository) *Usecase {
	return &Usecase{
		Config: cfg,
		Repo:   repo,
	}
}
