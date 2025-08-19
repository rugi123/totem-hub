package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/chirp/internal/interfaces"
)

func (s *Server) GetChatMessages(ctx *gin.Context) {
	clamis := ctx.MustGet("claims").(*interfaces.Claims)
	chatID := ctx.Param("id")

	messages, err := s.msgUC.GetChatMessages(ctx, chatID, clamis.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
