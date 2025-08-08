package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/repository/postgres"
	handler "github.com/rugi123/chirp/internal/transport/http"
	"github.com/rugi123/chirp/internal/usecase/auth"
	"github.com/rugi123/chirp/internal/usecase/chat"
	"github.com/rugi123/chirp/internal/usecase/message"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg, err := config.InitConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Panicln("received shutdown signal")
		cancel()
	}()

	initCtx, initCancel := context.WithTimeout(ctx, 10*time.Second)
	defer initCancel()

	g, gCtx := errgroup.WithContext(initCtx)

	var (
		userRepo   *postgres.UserRepo
		chatRepo   *postgres.ChatRepo
		memberRepo *postgres.MemberRepo
		msgRepo    *postgres.MessageRepo
	)

	g.Go(func() error {
		var err error
		userRepo, err = postgres.NewUserRepo(gCtx, cfg.Postgres)
		if err != nil {
			return fmt.Errorf("failed to init user repo: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		chatRepo, err = postgres.NewChatRepo(gCtx, cfg.Postgres)
		if err != nil {
			return fmt.Errorf("failed to init chat repo: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		memberRepo, err = postgres.NewMemberRepo(gCtx, cfg.Postgres)
		if err != nil {
			return fmt.Errorf("failed to init chat repo: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		msgRepo, err = postgres.NewMessageRepo(gCtx, cfg.Postgres)
		if err != nil {
			return fmt.Errorf("failed to init message repo: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("initialization failed: %v", err)
	}

	//тут юзкейсы

	fmt.Println(memberRepo)

	authUC := auth.NewAuthUsecase(cfg, userRepo)
	chatUC := chat.NewChatUsecase(cfg, chatRepo, memberRepo)
	msgUC := message.NewMessageUsecase(cfg, msgRepo)

	handler := handler.NewHanlder(*authUC, *chatUC, *msgUC)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	handler.RegisterRoutes(router)

	router.Run(":8080")
}
