package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/totem-hub/internal/config"
	"github.com/rugi123/totem-hub/internal/interfaces"
	"github.com/rugi123/totem-hub/internal/interfaces/websocket"
)

func (s *Server) RegisterRoutes(r *gin.Engine, cfg config.App) {
	r.GET("/ws", websocket.WSHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/profile", interfaces.AuthMiddleware(cfg.JWTKey), s.ShowProfilePage)

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "login.html", nil)
		})
		authGroup.GET("/register", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "register.html", nil)
		})

		authGroup.POST("/login", s.Login)
		authGroup.POST("/register", s.Register)
	}

	apiGroup := r.Group("/api")
	{

		chatGroup := apiGroup.Group("/chats")
		{
			//все чаты
			chatGroup.GET("/", interfaces.AuthMiddleware(cfg.JWTKey), s.GetChatsWithMembers)

			//создать чат
			chatGroup.POST("/", interfaces.AuthMiddleware(cfg.JWTKey), s.CreateChat)

			//по id
			chatGroup.GET("/:id", interfaces.AuthMiddleware(cfg.JWTKey))
			chatGroup.DELETE("/:id", interfaces.AuthMiddleware(cfg.JWTKey))

			messageGroup := chatGroup.Group("/:id/messages")
			{
				//получить все сообщения чата
				messageGroup.GET("/", interfaces.AuthMiddleware(cfg.JWTKey), s.GetChatMessages)
				//получить конкретное
				messageGroup.GET("/:id")
				messageGroup.POST("/:id")
			}
		}
	}
}
