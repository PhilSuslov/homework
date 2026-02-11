package repository


import (
	"context"

	"github.com/google/uuid"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

// Проверить Response и Res!
type OrderRepository interface {
	CreateOrder(order *orderV1.OrderDto)
	PayOrder(id string, transactionuuid uuid.UUID, paymentMethod orderV1.PaymentMethod) (*orderV1.PayOrderResponse, error)
	GetOrderByUUID(ctx context.Context, params orderV1.GetOrderByUUIDParams) (*orderV1.OrderDto, bool)
	CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (*orderV1.OrderDto, bool)
}