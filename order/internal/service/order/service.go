package order

import (
	repo "github.com/PhilSuslov/homework/order/internal/repository"
	def "github.com/PhilSuslov/homework/order/internal/service"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
	userV1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"



)

var _ def.OrderService = (*OrderService)(nil)

type OrderService struct {
	inventoryClient inventoryV1.InventoryServiceClient // Для красоты надо было использовать inventoryV1.ServiceClient Так как по пакету уже понятно какой это сервис
	paymentClient   paymentV1.PaymentServiceClient // Для красоты надо было использовать paymentV1.ServiceClient Так как по пакету уже понятно какой это сервис
	authClient  authV1.AuthServiceClient
	userClient userV1.UserServiceClient

	orderProducerService def.OrderProducerService
	orderService repo.OrderRepository
}

func NewOrderService(inventoryClient inventoryV1.InventoryServiceClient,
	paymentClient paymentV1.PaymentServiceClient, authClient  authV1.AuthServiceClient, 
	userClient userV1.UserServiceClient, orderProducerService def.OrderProducerService, 
	repo repo.OrderRepository) *OrderService {
	return &OrderService{
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
		authClient: authClient,
		userClient: userClient,
		orderProducerService: orderProducerService,
		orderService:    repo,
	}
}
