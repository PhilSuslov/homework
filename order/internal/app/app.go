package app

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PhilSuslov/homework/order/internal/config"
	orderProducer "github.com/PhilSuslov/homework/order/internal/service/producer/order_producer"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"

	orderAPI "github.com/PhilSuslov/homework/order/internal/api/order/v1"
	orderClientInv "github.com/PhilSuslov/homework/order/internal/client/grpc/inventory/v1"
	orderClientPay "github.com/PhilSuslov/homework/order/internal/client/grpc/payment/v1"
	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/order"
	orderService "github.com/PhilSuslov/homework/order/internal/service/order"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"

	"google.golang.org/grpc"
)

type App struct {
	diContainer *diContainer

	httpServer *http.Server

	payConn       *grpc.ClientConn
	inventoryConn *grpc.ClientConn
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	// Запускаем graceful shutdown в фоне
	go a.handleSignals(ctx)

	logger.Info(ctx, "HTTP server started on "+config.AppConfig().OrderHTTP.Address())
	return a.httpServer.ListenAndServe()
}

func (a *App) handleSignals(ctx context.Context) {
	// ловим SIGINT/SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	logger.Info(ctx, "🚦 Начинаем процесс graceful shutdown...")
	if err := a.Shutdown(ctx); err != nil {
		logger.Error(ctx, "Ошибка при shutdown", zap.Error(err))
	}
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initApp,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDI(ctx context.Context) error {
	a.diContainer = NewDiContainer(ctx)
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	cfg := config.AppConfig()
	if cfg == nil || cfg.Logger == nil {
		log.Fatal("config or config.Logger is nil!")
	}
	return logger.Init(cfg.Logger.Level(), cfg.Logger.AsJson())
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initApp(ctx context.Context) error {
	// --- clients ---
	paymentClient, payConn, err := orderClientPay.NewPaymentClient()
	if err != nil {
		return err
	}

	inventoryClient, invConn, err := orderClientInv.NewInventoryClient()
	if err != nil {
		return err
	}

	// --- producer ---
	producer := orderProducer.NewService(a.diContainer.OrderPaidProducer())

	// --- repository ---
	repo := orderRepo.NewOrderRepo(a.diContainer.postgresConn)

	// --- service ---
	orderSvc := orderService.NewOrderService(inventoryClient, paymentClient, producer, repo)

	// --- handler ---
	handler := orderAPI.NewOrderHandler(orderSvc)

	// --- ogen ---
	ogenServer, err := orderV1.NewServer(handler)
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Addr:    config.AppConfig().OrderHTTP.Address(),
		Handler: loggerMiddleware(ogenServer),
	}

	a.payConn = payConn
	a.inventoryConn = invConn

	// Регистрируем graceful shutdown для всех ресурсов
	closer.AddNamed("HTTP server", func(ctx context.Context) error {
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return a.httpServer.Shutdown(shutdownCtx)
	})

	closer.AddNamed("Payment GRPC conn", func(ctx context.Context) error {
		return a.payConn.Close()
	})

	closer.AddNamed("Inventory GRPC conn", func(ctx context.Context) error {
		return a.inventoryConn.Close()
	})

	return nil
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQUEST: %s %s — вход в OGEN", r.Method, r.URL.Path)

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(lrw, r)

		body, _ := io.ReadAll(r.Body)
		log.Printf("Body: %s, Query: %s", string(body), r.URL.RawQuery)
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		log.Printf("REQUEST: %s %s — обработан OGEN с кодом %d",
			r.Method, r.URL.Path, lrw.statusCode)
	})
}

func (a *App) Shutdown(ctx context.Context) error {
	return closer.CloseAll(ctx)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
