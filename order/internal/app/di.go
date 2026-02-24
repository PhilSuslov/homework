package app

import (
	"context"
	"log"
	
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/PhilSuslov/homework/order/internal/repository"
	"github.com/PhilSuslov/homework/order/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"

	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/order"

	orderService "github.com/PhilSuslov/homework/order/internal/service/order"
)

type diContainer struct {
	orderService    service.OrderService
	orderRepository repository.OrderRepository
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient

	postgresConn *pgxpool.Pool
}

func NewDiContainer(ctx context.Context) *diContainer {
	d := &diContainer{}
	d.postgresConn = d.PostgresCfg(ctx)
	return d
}

func (d *diContainer) OrderService(ctx context.Context) service.OrderService {
	if d.orderService == nil {
		d.orderService = orderService.NewOrderService(d.inventoryClient, d.paymentClient, d.OrderRepository(ctx))
	}
	return d.orderService
}

func (d *diContainer) OrderRepository(ctx context.Context) repository.OrderRepository {
	if d.orderRepository == nil {
		if d.postgresConn == nil {
			d.postgresConn = d.PostgresCfg(ctx)
		}
		d.orderRepository = orderRepo.NewOrderRepo(d.postgresConn)
	}
	return d.orderRepository
}

func (d *diContainer) PostgresCfg(ctx context.Context) *pgxpool.Pool {
	if d.postgresConn != nil {
		return d.postgresConn
	}

	conn, err := pgxpool.New(ctx, "postgres://demo:demo@localhost:5435/orders?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Проверяем доступность базы
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	// Миграции через Goose
	migrationsDir := "../migrations"
	db := stdlib.OpenDB(*conn.Config().ConnConfig)
	migrationsRunner := orderRepo.NewMigrator(db, migrationsDir)

	err = migrationsRunner.Up()
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	d.postgresConn = conn
	return d.postgresConn
}