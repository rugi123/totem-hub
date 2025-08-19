package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (S *Server) ShowProfilePage(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "profile.html", nil)
}
