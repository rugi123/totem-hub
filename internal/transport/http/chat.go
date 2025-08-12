package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/transport"
)

func (h *Handler) LoadChatList(ctx *gin.Context) {
	user_id := ctx.MustGet("claims").(*transport.Claims)
	id, err := uuid.Parse(user_id.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token or claims",
		})
	}

	chats, err := h.chatUC.LoadChatList(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("load chats %v", err),
		})
	}

	ctx.JSON(http.StatusOK, chats)
}
