package app

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_server/internal/closer"
	"github.com/satanaroom/chat_server/internal/config"
	"github.com/satanaroom/chat_server/internal/interceptor"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
	_ "github.com/satanaroom/chat_server/statik"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swagger         *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, fmt.Errorf("init deps: %w", err)
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		if err := a.runGRPCServer(); err != nil {
			logger.Fatalf("run GRPC server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runHTTPServer(); err != nil {
			logger.Fatalf("run HTTP server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runSwaggerServer(); err != nil {
			logger.Fatalf("run swagger server: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
		a.initSwaggerServer,
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

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(a.serviceProvider.TLSServerCredentials(ctx)),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.ValidateInterceptor,
				interceptor.ErrorCodesInterceptor,
				interceptor.NewAuthInterceptor(a.serviceProvider.AuthClient(ctx)).Unary(),
			),
		),
	)

	reflection.Register(a.grpcServer)

	chatV1.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := chatV1.RegisterChatV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Host(), opts); err != nil {
		return fmt.Errorf("register user v1 handler from endpoint: %w", err)
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Host(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return fmt.Errorf("failed to create statik file system: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/chat.swagger.json/", serveSwaggerFile("/chat.swagger.json"))

	a.swagger = &http.Server{
		Addr:    a.serviceProvider.SwaggerConfig().Host(),
		Handler: mux,
	}

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Infof("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Host())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Host())
	if err != nil {
		return fmt.Errorf("failed to get listener: %s", err.Error())
	}

	if err = a.grpcServer.Serve(list); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Host())

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	logger.Infof("Swagger server is running on %s", a.serviceProvider.SwaggerConfig().Host())

	if err := a.swagger.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
