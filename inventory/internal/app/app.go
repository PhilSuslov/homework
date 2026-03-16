package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/PhilSuslov/homework/inventory/internal/config"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	"github.com/PhilSuslov/homework/platform/pkg/grpc/health"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	authGRPC "github.com/PhilSuslov/homework/platform/pkg/middleware/grpc"
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	diContainer *diContainer
	grpcServer  *grpc.Server
	listener    net.Listener
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.diContainer.InitFakeData(ctx); err != nil {
		log.Printf("Failed to init fake data: %v", err)
		return err
	}
	return a.runGRPCServer(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initListener,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initListener(_ context.Context) error {
	listener, err := net.Listen("tcp", config.AppConfig().InventoryGRPC.Address())
	if err != nil {
		return err
	}
	closer.AddNamed("TCP Listener", func(ctx context.Context) error {
		lerr := listener.Close()
		if lerr != nil && errors.Is(lerr, net.ErrClosed) {
			return lerr
		}
		return nil
	})
	a.listener = listener
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	// --- Создаем gRPC клиента для AuthService ---
	authClientConn, err := grpc.NewClient(config.AppConfig().IamGRPC.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to dial AuthService: %v", err)
		return err
	}
	authClient := authV1.NewAuthServiceClient(authClientConn)

	// --- Создаем interceptor ---
	authInterceptor := authGRPC.NewAuthInterceptor(authClient)

	// --- Создаем gRPC сервер с interceptor ---
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)

	// --- Добавляем graceful shutdown ---
	closer.AddNamed("gRPC server", func(ctx context.Context) error {
		a.grpcServer.GracefulStop()
		return nil
	})

	// --- Reflection и Health ---
	reflection.Register(a.grpcServer)
	health.RegisterService(a.grpcServer)

	// --- Регистрируем InventoryService ---
	srv, err := a.diContainer.InventoryV1API(ctx)
	if err != nil {
		log.Printf("Failed to get InventoryV1API: %v", err)
		return err
	}
	inventoryV1.RegisterInventoryServiceServer(a.grpcServer, srv)

	return nil
}

func (a *App) runGRPCServer(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("🚀 gRPC InventoryService server listening on %s", config.AppConfig().InventoryGRPC.Address()))

	err := a.grpcServer.Serve(a.listener)
	if err != nil {
		return err
	}

	return nil
}
