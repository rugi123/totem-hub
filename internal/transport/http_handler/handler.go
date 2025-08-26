package http_handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/internal/dto"
)

type MessageUsecase interface {
	GetChatMessages(ctx context.Context, chatID uuid.UUID, userID uuid.UUID) ([]entity.Message, error)
}
type MemberUsecase interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Member, error)
}
type ChatUsecase interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Chat, error)
	Create(ctx context.Context, dto dto.CreateChat) error
}

type HTTPHandler struct {
	messageUC MessageUsecase
	memberUC  MemberUsecase
	chatUC    ChatUsecase
}
