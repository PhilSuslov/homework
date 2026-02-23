package app

import (
	"context"

	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"

	api "github.com/PhilSuslov/homework/order/internal/api/order/v1"
	"github.com/PhilSuslov/homework/order/internal/repository"
	"github.com/PhilSuslov/homework/order/internal/service"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	"github.com/jackc/pgx/v5/pgxpool"

	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/order"

	orderService "github.com/PhilSuslov/homework/order/internal/service/order"
)

type diContainer struct {
	orderV1API      orderV1.Handler
	orderService    service.OrderService
	orderRepository repository.OrderRepository
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient

	postgresConn *pgxpool.Pool
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderV1API(ctx context.Context) orderV1.Handler {
	if d.orderV1API == nil {
		d.orderV1API = api.NewAPI(d.PartService(ctx))
	}
	return d.orderV1API
}

func (d *diContainer) PartService(ctx context.Context) service.OrderService {
	if d.orderService == nil {
		d.orderService = orderService.NewOrderService(d.inventoryClient, d.paymentClient, d.PartRepository(ctx))
	}
	return d.orderService
}

func (d *diContainer) PartRepository(ctx context.Context) repository.OrderRepository {
	if d.orderRepository == nil {
		d.orderRepository = orderRepo.NewOrderRepo(d.postgresConn)
	}
	return d.orderRepository
}
