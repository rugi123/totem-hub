package usecase

import (
	"context"

	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
}

type Usecase struct {
	repo   Repository
	config config.Config
}

func NewAuthUsecase(cfg config.Config, repo Repository) *Usecase {
	return &Usecase{
		config: cfg,
		repo:   repo,
	}
}
