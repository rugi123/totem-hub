package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowProfilePage(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "profile.html", nil)
}
