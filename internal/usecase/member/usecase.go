package member

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type MemberRepository interface {
	GetByUserID(ctx context.Context, user_id uuid.UUID) ([]entity.Member, error)
}

type Usecase struct {
	MemberRepo MemberRepository
	Config     config.Config
}

func NewMessageUsecase(cfg config.Config, repo MemberRepository) *Usecase {
	return &Usecase{
		Config:     cfg,
		MemberRepo: repo,
	}
}
