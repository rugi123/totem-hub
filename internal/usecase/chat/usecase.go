package chat

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/internal/usecase/member"
)

type ChatRepository interface {
	GetByIDs(ctx context.Context, ids []uuid.UUID) (*entity.DataResult, error)
}

type Usecase struct {
	repo       ChatRepository
	memberRepo member.MemberRepository
}

func NewChatUsecase(repo ChatRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) LoadChats(ctx context.Context, user_id string) ([]entity.Member, *entity.DataResult, error) {
	id, err := uuid.Parse(user_id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed parse id: %w", err)
	}
	members, err := u.memberRepo.GetByUserID(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("get members error: %w", err)
	}
	IDs := entity.ExtractMemberIDs(members)
	chats, err := u.repo.GetByIDs(ctx, IDs)
	if err != nil {
		return nil, nil, fmt.Errorf("get chats error: %w", err)
	}
	return members, chats, nil
}
