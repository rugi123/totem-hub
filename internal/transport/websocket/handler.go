package websocket

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rugi123/totem-hub/internal/transport"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type RegisterClientUsecase interface {
	Execute(ctx context.Context, conn *websocket.Conn, userID uuid.UUID) error
}

type WSHandler struct {
	registerClientUsecase RegisterClientUsecase
}

func (h *WSHandler) RegisterClient(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*transport.Claims)
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.registerClientUsecase.Execute(ctx, conn, userID)
}
