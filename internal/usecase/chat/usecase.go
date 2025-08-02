package chat

import (
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type Repository interface {
	GetChat(id int) (*entity.User, error)
	CreateChat(user *entity.User) error
	UpdateChat(user *entity.User) error
	DeleteChat(id int) error
}

type Usecase struct {
	Repo   Repository
	Config config.Config
}

func NewChatUsecase(cfg config.Config, repo Repository) *Usecase {
	return &Usecase{
		Config: cfg,
		Repo:   repo,
	}
}
