package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/transport"
)

func (h *HTTPHandler) GetChatMessages(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*transport.Claims)
	chatID := ctx.Param("id")
	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	userUUID, err := uuid.Parse(claims.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	messages, err := h.messageUC.GetChatMessages(ctx, chatUUID, userUUID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
