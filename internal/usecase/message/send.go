package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/internal/dto"
)

type MessageSaver interface {
	Save(ctx context.Context, msg *entity.Message) error
}

type MemberGetter interface {
	GetByID(ctx context.Context, ID uuid.UUID) (*entity.Member, error)
}

type WebsocketHub interface {
	BroadcastToChat(memberID uuid.UUID, message dto.BroadcastMessage)
}

type SendMessageUsecase struct {
	websocketHub WebsocketHub
	messageSaver MessageSaver
	memberGetter MemberGetter
}

func NewUsecase(messageSaver MessageSaver, memberGetter MemberGetter, websocketHub WebsocketHub) *SendMessageUsecase {
	return &SendMessageUsecase{messageSaver: messageSaver, memberGetter: memberGetter, websocketHub: websocketHub}
}

func (u *SendMessageUsecase) Execute(ctx context.Context, userID uuid.UUID, sendDto dto.SendMessage) error {
	member, err := u.memberGetter.GetByID(ctx, sendDto.MemberID)
	if err != nil {
		return err
	}
	if member.ID != userID || member.IsMuted {
		return domain.ErrMessageSentRefused
	}

	message := entity.NewMessage(sendDto.MemberID, sendDto.Text)

	err = u.messageSaver.Save(ctx, message)
	if err != nil {
		return err
	}

	broadcastDto := dto.BroadcastMessage{
		ID:       message.ID,
		MemberID: message.MemberID,
		Text:     message.Text,
		SentAt:   message.SentAt,
	}

	u.websocketHub.BroadcastToChat(sendDto.MemberID, broadcastDto)

	return nil
}
