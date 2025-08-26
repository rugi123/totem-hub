package http_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/totem-hub/internal/dto"
	"github.com/rugi123/totem-hub/pkg/validator"
)

func (h *HTTPHandler) Login(ctx *gin.Context) {
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

	token, err := s.userUC.Login(ctx, dto, s.config.JWTKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "get user: " + err.Error(),
		})
		return
	}

	ctx.SetCookie("token", token, int(15*time.Minute), "/", "localhost", true, true)

	ctx.JSON(http.StatusOK, nil)

	ctx.Redirect(http.StatusPermanentRedirect, "/profile")
}

func (h *HTTPHandler) Register(ctx *gin.Context) {
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

	id, err := s.userUC.Register(ctx, &dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}
