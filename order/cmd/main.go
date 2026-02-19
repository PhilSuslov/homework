package main

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

	orderAPI "github.com/PhilSuslov/homework/order/internal/api/order/v1"
	orderClientInv "github.com/PhilSuslov/homework/order/internal/client/grpc/inventory/v1"
	orderClientPay "github.com/PhilSuslov/homework/order/internal/client/grpc/payment/v1"
	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/order"
	orderService "github.com/PhilSuslov/homework/order/internal/service/order"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

const (
	httpPort = "8080"
	// readHeaderTimeout = 5 * time.Second
	shutdownTimeout = 10 * time.Second
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func main() {
	ctx := context.Background()
	
	conn, err := pgx.Connect(ctx, "postgres://demo:demo@localhost:5432/order-service")
	if err != nil{
		log.Printf("failed to connect to database: %v\n", err)
		return
	}
	defer func(){
		cerr := conn.Close(ctx)
		if cerr != nil{
			log.Printf("failed to close connection: %v\n", err)
			return
		}
	}()

	err = conn.Ping(ctx)
	if err != nil{
		log.Printf("База данных недоступна: %v\n", err)
		return
	}

	migrationsDir := "../migrations"
	migrationsRunner := orderRepo.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), migrationsDir)

	err = migrationsRunner.Up()
	if err != nil{
		log.Printf("Ошибка миграции базы данных: %v\n", err)
	}

	// ---------------- Order service ----------------
	paymentClient, payConn, err := orderClientPay.NewPaymentClient()
	if err != nil {
		log.Println("Ошибка в paymentClient")
	}
	defer payConn.Close()

	inventoryClient, invConn, err := orderClientInv.NewInventoryClient()
	if err != nil {
		log.Println("Ошибка в paymentClient")
	}
	defer invConn.Close()

	repo := orderRepo.NewOrderRepo(conn)
	orderService := orderService.NewOrderService(inventoryClient, paymentClient, repo)
	handler := orderAPI.NewOrderHandler(orderService)

	// ---------------- OGEN server ----------------

	ogenServer, err := orderV1.NewServer(handler)
	if err != nil {
		log.Fatalf("failed to create ogen server: %v", err)
	}

	httpServer := &http.Server{
		Addr: ":" + httpPort,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("REQUEST: %s %s — вход в OGEN", r.Method, r.URL.Path)

			// Создаем обертку ResponseWriter, чтобы логировать статус-код
			lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
			ogenServer.ServeHTTP(lrw, r)

			body, _ := io.ReadAll(r.Body)
			log.Printf("Body: %s, Query: %s", string(body), r.URL.RawQuery)
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			log.Printf("REQUEST: %s %s — обработан OGEN с кодом %d", r.Method, r.URL.Path, lrw.statusCode)
		}),
	}

	// ---------------- Run HTTP server ----------------

	go func() {
		log.Printf("OrderService started on :%s", httpPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %v", err)
		}
	}()

	// ---------------- Graceful shutdown ----------------

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	log.Println("bye 👋")
}

// ("===== CHECK ROUTES =====") 82 str
// 	log.Println("===== CHECK ROUTES =====")
// 	paths := []string{
// 		"/api/v1/order/",
// 		"/api/v1/order/123e4567-e89b-12d3-a456-426614174000",
// 	}
// 	methods := []string{"GET", "POST"}
//
// 	for _, p := range paths {
// 		for _, m := range methods {
// 			_, ok := ogenServer.FindRoute(m, p)
// 			log.Println(m, p, "->", ok)
// 		}
// 	}
