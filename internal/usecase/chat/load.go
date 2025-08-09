package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/dto"
)

func (u *Usecase) LoadChatList(ctx context.Context, user_id uuid.UUID) (*[]dto.GetChatResponse, []uuid.UUID, error) {

	chats, err := u.ChatRepo.GetAllUserChats(ctx, user_id)
	if err != nil {
		return nil, nil, nil
	}
	members, err := u.MemberRepo.GetAllUserMembers(ctx, user_id)
	if err != nil {
		return nil, nil, nil
	}

	return nil, nil, nil
}
