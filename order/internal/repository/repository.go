package repository

import (
	"context"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)


type OrderRepository interface {
	CreateOrder(ctx context.Context, order *orderRepoModel.OrderDto)
	PayOrder(ctx context.Context, orderUUID uuid.UUID, uuidPay uuid.UUID, paymentMethod string) (*string, error)
	PayOrderCreate(ctx context.Context, req *orderRepoModel.PayOrderRequest, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool)
	GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool)
	CancelOrder(ctx context.Context, orderUUID uuid.UUID) (bool)
}
