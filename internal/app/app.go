package app

import (
	"context"
	"fmt"

	"github.com/satanaroom/chat_server/internal/closer"
	"github.com/satanaroom/chat_server/internal/config"
	"google.golang.org/grpc/metadata"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, fmt.Errorf("init deps: %w", err)
	}

	return a, nil
}

const accessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ5OTYwMjAsInVzZXJuYW1lIjoicTJxMnEyIiwicm9sZSI6MX0.2isMeJ0mNhOXF6Z1NSimeP6W7LMfhKmfEcL_MPWFn1M"

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	ctx := context.Background()

	m := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}
	md := metadata.New(m)

	ctx = metadata.NewOutgoingContext(ctx, md)

	ok, err := a.serviceProvider.AuthClient(ctx).Check(ctx, "bla")
	fmt.Println(err, ok)

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,

		a.initServiceProvider,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return fmt.Errorf("init: %w", err)
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
