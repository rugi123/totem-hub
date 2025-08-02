package message

import (
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type Repository interface {
	RecieveMessages(id int) (*entity.User, error)
	SendMessage(user *entity.User) error
	EditMessage(user *entity.User) error
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
