package chat

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/internal/dto"
)

type ChatRepository interface {
	Create(ctx context.Context, chat entity.Chat) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Chat, error)
	GetByIDs(ctx context.Context, id []uuid.UUID) ([]entity.Chat, error)
}

type Usecase struct {
	repo ChatRepository
}

func NewUsecase(repo ChatRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetChats(ctx context.Context, IDs []uuid.UUID) ([]entity.Chat, error) {
	chats, err := u.repo.GetByIDs(ctx, IDs)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (u *Usecase) GetChat(ctx context.Context, user_id string) (*entity.Chat, error) {
	userUUID, err := uuid.Parse(user_id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	chat, err := u.repo.GetByID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("get chat error: %w", err)
	}
	return chat, nil
}

func (u *Usecase) Create(ctx context.Context, req dto.CreateChat, user_id string) error {
	userUUID, err := uuid.Parse(user_id)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	creator, err := u.getChatCreator(req.Type)
	if err != nil {
		return err
	}

	if err := creator.ValidateAttributes(req.Attributes); err != nil {
		return fmt.Errorf("invalid attributes: %w", err)
	}

	chat, err := creator.CreateEntity(req.Type, req.Title, userUUID, req.Attributes)
	if err != nil {
		return fmt.Errorf("failed to create chat entity: %w", err)
	}

	if err := u.repo.Create(ctx, chat); err != nil {
		return fmt.Errorf("repository error: %w", err)
	}

	return nil
}

func (u *Usecase) getChatCreator(chatType string) (ChatCreator, error) {
	switch chatType {
	case "channel":
		return &ChannelCreator{}, nil
	case "group":
		return &GroupCreator{}, nil
	case "diolog":
		return &DiologCreator{}, nil
	default:
		return nil, fmt.Errorf("unsupported chat type: %s", chatType)
	}
}
