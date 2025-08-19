package message

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/internal/dto"
)

type MessageRepository interface {
	GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, userID uuid.UUID) ([]entity.Message, error)
	Create(ctx context.Context, msg *entity.Message) error
	Update(ctx context.Context, msg *entity.Message) error
	Delete(ctx context.Context, msg_id uuid.UUID) error
}

type Usecase struct {
	MessageRepo MessageRepository
}

func NewUsecase(repo MessageRepository) *Usecase {
	return &Usecase{
		MessageRepo: repo,
	}
}

func (u *Usecase) GetChatMessages(ctx context.Context, chatID string, userID string) ([]entity.Message, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		return nil, err
	}
	//перенеси в транспорнтый слой уебан
	messages, err := u.MessageRepo.GetMessagesByChatID(ctx, chatUUID, userUUID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (u *Usecase) CreateMessage(ctx context.Context, dto dto.SendMessage) error {
	msg := entity.Message{
		ID:       uuid.New(),
		MemberID: dto.MemberID,
		Text:     dto.Text,
		SentAt:   time.Now(),
		EditedAt: sql.NullTime{},
	}
	err := u.MessageRepo.Create(ctx, &msg)
	return err
}
