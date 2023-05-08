package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_server/internal/app"
)

func main() {
	ctx := context.Background()

	chatApp, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatalf("failed to initialize app: %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	logger.Info("service starting up")

	if err = chatApp.Run(); err != nil {
		logger.Fatalf("failed to run app: %s", err.Error())
	}

	<-quit

	logger.Info("service shutting down")
}
