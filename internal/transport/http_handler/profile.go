package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) ShowProfilePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "profile.html", nil)
}
