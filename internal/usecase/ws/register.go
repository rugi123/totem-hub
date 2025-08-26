package ws

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rugi123/totem-hub/internal/domain/entity"
)

type MemberGetter interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Member, error)
}

type WebsocketHub interface {
	RegisterClient(conn *websocket.Conn, memberID uuid.UUID)
}

type RegisterClientUsecase struct {
	memberGetter MemberGetter
	websocketHub WebsocketHub
}

func (u *RegisterClientUsecase) Execute(ctx context.Context, conn *websocket.Conn, userID uuid.UUID) error {
	member, err := u.memberGetter.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	u.websocketHub.RegisterClient(conn, member.ID)
	return nil
}
