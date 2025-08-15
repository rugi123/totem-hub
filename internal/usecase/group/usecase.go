package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type GroupRepository interface {
	Create(ctx context.Context, chat *entity.Chat) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	Update(ctx context.Context, chat *entity.Chat) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Usecase struct {
	repo *GroupRepository
}

func NewGroupUsecase(repo GroupRepository) *Usecase {
	return &Usecase{
		repo: &repo,
	}
}
