package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/totem-hub/internal/config"
	"github.com/rugi123/totem-hub/internal/repository/postgres"
	"github.com/rugi123/totem-hub/internal/transport/http"
	"github.com/rugi123/totem-hub/internal/usecase/chat"
	"github.com/rugi123/totem-hub/internal/usecase/member"
	"github.com/rugi123/totem-hub/internal/usecase/message"
	"github.com/rugi123/totem-hub/internal/usecase/user"
	"github.com/rugi123/totem-hub/pkg/database"
)

func main() {
	// Инициализация конфигурации
	cfg, err := config.InitConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация контекста с graceful shutdown
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	// Инициализация базы данных с таймаутом
	initCtx, initCancel := context.WithTimeout(ctx, 10*time.Second)
	defer initCancel()

	db, err := database.NewPostgres(initCtx, config.CreateConn(cfg.Postgres))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Закрываем соединение при завершении

	// Инициализация репозиториев
	repos := struct {
		user    *postgres.UserRepository
		chat    *postgres.ChatRepository
		member  *postgres.MemberRepository
		message *postgres.MessageRepository
	}{
		user:    postgres.NewUserRepository(db),
		chat:    postgres.NewChatRepository(db),
		member:  postgres.NewMemberRepository(db),
		message: postgres.NewMessageRepository(db),
	}

	// Инициализация use cases
	useCases := struct {
		user    *user.Usecase
		chat    *chat.Usecase
		member  *member.Usecase
		message *message.Usecase
	}{
		user:    user.NewUsecase(repos.user),
		chat:    chat.NewUsecase(repos.chat),
		member:  member.NewUsecase(repos.member),
		message: message.NewUsecase(repos.message),
	}

	// Инициализация HTTP сервера
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	httpServer := http.NewServer(
		cfg.App,
		*useCases.user,
		*useCases.chat,
		*useCases.member,
		*useCases.message,
	)
	httpServer.RegisterRoutes(router, cfg.App)

	// Запуск сервера в отдельной goroutine
	serverErr := make(chan error, 1)
	go func() {
		log.Printf("Starting server on :%s", cfg.App.Port)
		if err := router.Run(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
			serverErr <- fmt.Errorf("server failed: %w", err)
		}
	}()

	// Ожидание shutdown сигнала или ошибки сервера
	select {
	case <-ctx.Done():
		log.Println("Shutting down gracefully...")
	case err := <-serverErr:
		log.Printf("Server error: %v", err)
		cancel()
	}

	// Дополнительное время для graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer shutdownCancel()

	// Здесь можно добавить закрытие других ресурсов (например, соединений с БД)
	<-shutdownCtx.Done()
	log.Println("Server exited")
}
