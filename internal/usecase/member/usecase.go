package member

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type MemberRepository interface {
	GetMembers(ctx context.Context, user_id uuid.UUID) ([]entity.Member, error)
	GetMember(ctx context.Context, user_id uuid.UUID) (*entity.Member, error)
}

type Usecase struct {
	MemberRepo MemberRepository
}

func NewUsecase(repo MemberRepository) *Usecase {
	return &Usecase{
		MemberRepo: repo,
	}
}

func (u *Usecase) GetMembers(ctx context.Context, user_id string) ([]entity.Member, error) {
	id, err := uuid.Parse(user_id)
	if err != nil {
		return nil, err
	}
	members, err := u.MemberRepo.GetMembers(ctx, id)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (u *Usecase) GetMember(ctx context.Context, user_id string) (*entity.Member, error) {
	id, err := uuid.Parse(user_id)
	if err != nil {
		return nil, err
	}

	member, err := u.MemberRepo.GetMember(ctx, id)
	return member, err
}
