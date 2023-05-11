package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
)

func Init(_ context.Context) error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	return nil
}
