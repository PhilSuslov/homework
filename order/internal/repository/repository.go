package repository

import (
	"context"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

// Проверить Response и Res!
type OrderRepository interface {
	CreateOrder(order *orderRepoModel.OrderDto)
	PayOrder(orderUUID uuid.UUID, payUUID uuid.UUID, paymentMethod string) (*string, error)
	GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool)
	CancelOrder(ctx context.Context, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool)
}
