package message

import (
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type Repository interface {
	GetMessages(id int) (*entity.User, error)
	CreateMessage(user *entity.User) error
	UpdateMessage(user *entity.User) error
	DeleteMessage(id int) error
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
