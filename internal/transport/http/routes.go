package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/chirp/internal/transport"
)

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/profile", transport.AuthMiddleware(h.authUC.Config.App.JWTKey), h.ShowProfilePage)

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "login.html", nil)
		})
		authGroup.GET("/register", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "register.html", nil)
		})

		authGroup.POST("/login", h.Login)
		authGroup.POST("/register", h.Register)
	}

	apiGroup := r.Group("/api")
	{

		chatGroup := apiGroup.Group("/chats")
		{
			//все чаты
			chatGroup.GET("/", transport.AuthMiddleware(h.authUC.Config.App.JWTKey), h.LoadChats)

			//конкретный чат
			chatGroup.GET("/:id")
			chatGroup.POST("/:id")
			chatGroup.DELETE("/:id")
		}
	}
}
