package main

import (
	"context"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_server/internal/app"
)

func main() {
	ctx := context.Background()

	chatApp, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatalf("failed to initialize app: %s", err.Error())
	}

	logger.Info("service starting up")

	if err = chatApp.Run(); err != nil {
		logger.Fatalf("failed to run app: %s", err.Error())
	}

	logger.Info("service shutting down")
}
