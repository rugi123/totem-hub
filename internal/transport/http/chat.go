package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/chirp/internal/transport"
)

func (h *Handler) LoadChats(ctx *gin.Context) {
	clamis := ctx.MustGet("claims").(transport.Claims)
	members, chats, err := h.chatUC.LoadChats(ctx, clamis.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("load chats error: %v", err),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"members": members,
		"chats":   chats,
	})
}
