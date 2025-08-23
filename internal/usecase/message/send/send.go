package send

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

type Usecase struct {
	messageSaver MessageSaver
	memberGetter MemberGetter
}

func NewUsecase(messageSaver MessageSaver, memberGetter MemberGetter) (*Usecase, error) {
	if messageSaver == nil {
		return nil, domain.ErrEmptyField
	}
	if memberGetter == nil {
		return nil, domain.ErrEmptyField
	}
	return &Usecase{messageSaver: messageSaver, memberGetter: memberGetter}, nil
}

func (u *Usecase) Execute(ctx context.Context, userID uuid.UUID, dto dto.SendMessage) error {
	member, err := u.memberGetter.GetByID(ctx, dto.MemberID)
	if err != nil {
		return err
	}
	if member.ID != userID || member.IsMuted {
		return domain.ErrMessageSentRefused
	}

	msg := entity.NewMessage(dto.MemberID, dto.Text)

	err = u.messageSaver.Save(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
