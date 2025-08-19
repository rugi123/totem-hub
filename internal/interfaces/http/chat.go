package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/dto"
	"github.com/rugi123/chirp/internal/interfaces"
)

func (s *Server) GetChatsWithMembers(ctx *gin.Context) {
	clamis := ctx.MustGet("claims").(*interfaces.Claims)

	members, err := s.memberUC.GetMembers(ctx, clamis.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("get members error: %v", err),
		})
	}
	var IDs []uuid.UUID
	for _, member := range members {
		IDs = append(IDs, member.ChatID)
	}

	chats, err := s.chatUC.GetChats(ctx, IDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("get chats error: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"chats":   chats,
		"members": members,
	})
}

func (s *Server) GetChatWithMember(ctx *gin.Context) {
	clamis := ctx.MustGet("claims").(*interfaces.Claims)

	chat, err := s.chatUC.GetChat(ctx, clamis.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	member, err := s.memberUC.GetMember(ctx, clamis.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

	}

	ctx.JSON(http.StatusOK, gin.H{
		"chat":   chat,
		"member": member,
	})

}

func (s *Server) CreateChat(ctx *gin.Context) {
	clamis := ctx.MustGet("claims").(*interfaces.Claims)

	var dto dto.CreateChat
	err := ctx.ShouldBindBodyWithJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = s.chatUC.Create(ctx, dto, clamis.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
