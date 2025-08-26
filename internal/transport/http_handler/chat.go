package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/dto"
	"github.com/rugi123/totem-hub/internal/transport"
)

func (h *HTTPHandler) GetChats(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*transport.Claims)
	userUUID, err := uuid.Parse(claims.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	chats, err := h.chatUC.GetByUserID(ctx, userUUID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, chats)
}

func (h *HTTPHandler) CreateChat(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*transport.Claims)
	userUUID, err := uuid.Parse(claims.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var dto dto.CreateChat
	err = ctx.ShouldBindBodyWithJSON(&dto)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = h.chatUC.Create(ctx, dto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
