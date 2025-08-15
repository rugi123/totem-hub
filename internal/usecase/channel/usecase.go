package channel

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type ChannelRepository interface {
	Create(ctx context.Context, chat *entity.Chat) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	Update(ctx context.Context, chat *entity.Chat) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Usecase struct {
	repo *ChannelRepository
}

func NewGroupUsecase(repo ChannelRepository) *Usecase {
	return &Usecase{
		repo: &repo,
	}
}
