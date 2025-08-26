package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/config"
	"github.com/rugi123/totem-hub/internal/transport"
	"github.com/rugi123/totem-hub/internal/transport/http_handler"
	"github.com/rugi123/totem-hub/internal/transport/websocket"
)

type Server struct {
	router      *gin.Engine
	config      config.App
	httpHandler http_handler.HTTPHandler
	wsHandler   websocket.WSHandler
}

func (s *Server) RegisterRoutes() {
	s.router.GET("/ws", s.wsHandler.RegisterClient)

	s.router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	s.router.GET("/profile", transport.AuthMiddleware(s.config.JWTKey), s.httpHandler.ShowProfilePage)

	authGroup := s.router.Group("/auth")
	{
		authGroup.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "login.html", nil)
		})
		authGroup.GET("/register", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "register.html", nil)
		})

		authGroup.POST("/login", s.httpHandler.Login)
		authGroup.POST("/register", s.httpHandler.Register)
	}

	apiGroup := s.router.Group("/api")
	{

		chatGroup := apiGroup.Group("/chats")
		{
			//все чаты
			chatGroup.GET("/", transport.AuthMiddleware(s.config.JWTKey), s.httpHandler.GetChatsWithMembers)

			//создать чат
			chatGroup.POST("/", transport.AuthMiddleware(s.config.JWTKey), s.httpHandler.CreateChat)

			//по id
			chatGroup.GET("/:id", transport.AuthMiddleware(s.config.JWTKey))
			chatGroup.DELETE("/:id", transport.AuthMiddleware(s.config.JWTKey))

			//вот тут сделай братанчик отдельный обработчик для отправки сообщения который вызывает
			// broadcast который отсылает клиентом чата сообщение без повторного запроса

			messageGroup := chatGroup.Group("/:id/messages")
			{
				//получить все сообщения чата
				messageGroup.GET("/", transport.AuthMiddleware(s.config.JWTKey), s.httpHandler.GetChatMessages)
				//получить конкретное
				messageGroup.GET("/:id")
				//отправить
				messageGroup.POST("/:id", s.SendMessage)
			}
		}
	}
}

func (s *Server) SendMessage(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*transport.Claims)
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	s.httpHandler.
}
