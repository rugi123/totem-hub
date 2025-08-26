package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain/entity"
)

type MessageGetter interface {
	GetByChatID(ctx context.Context, chatID uuid.UUID, userID uuid.UUID) ([]entity.Message, error)
}

type GetMessageUsecase struct {
	messageGetter MessageGetter
}

func (u *GetMessageUsecase) GetChatMessages(ctx context.Context, chatID uuid.UUID, userID uuid.UUID) ([]entity.Message, error) {
	messages, err := u.messageGetter.GetByChatID(ctx, chatID, userID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
