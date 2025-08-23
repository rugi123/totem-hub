package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Usecase struct {
	UserRepo UserRepository
}

func NewUsecase(repo UserRepository) *Usecase {
	return &Usecase{
		UserRepo: repo,
	}
}
