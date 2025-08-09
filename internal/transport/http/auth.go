package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/chirp/internal/dto"
	"github.com/rugi123/chirp/pkg/validator"
)

func (h *Handler) Login(ctx *gin.Context) {
	var dto dto.LoginRequest
	err := ctx.ShouldBindBodyWithJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "parse json: " + err.Error(),
		})
		return
	}

	if err := validator.Validate(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "validate error: " + err.Error(),
		})
		return
	}

	token, err := h.authUC.Login(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "get user: " + err.Error(),
		})
		return
	}

	ctx.SetCookie("token", token, int(time.Now().Add(15+time.Minute).Unix()), "/", "localhost", true, true)

	ctx.JSON(http.StatusOK, nil)

	ctx.Redirect(http.StatusPermanentRedirect, "/profile")
}

func (h *Handler) Register(ctx *gin.Context) {
	var dto dto.RegisterRequest
	err := ctx.ShouldBindBodyWithJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "parse json: " + err.Error(),
		})
		return
	}

	if err := validator.Validate(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "validate error: " + err.Error(),
		})
		return
	}

	id, err := h.authUC.Register(ctx, &dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}
