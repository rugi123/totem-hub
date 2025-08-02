package usecase

import (
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type Repository interface {
	GetUser(id int) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	DeleteUser(id int) error
}

type Usecase struct {
	Repo   Repository
	Config config.Config
}

func NewAuthUsecase(cfg config.Config, repo Repository) *Usecase {
	return &Usecase{
		Config: cfg,
		Repo:   repo,
	}
}
