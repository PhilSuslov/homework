package order

import (
	repo "github.com/PhilSuslov/homework/order/internal/repository/order"
	def "github.com/PhilSuslov/homework/order/internal/service"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
)

var _ def.OrderService = (*OrderService)(nil)

type OrderService struct {
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient

	orderService *repo.OrderRepo
}

func NewOrderService(inventoryClient inventoryV1.InventoryServiceClient,
	paymentClient paymentV1.PaymentServiceClient, repo *repo.OrderRepo) *OrderService {
	return &OrderService{
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
		orderService: repo,
	}
}
