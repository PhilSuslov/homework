package service

import (
	"context"

	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request *orderServiceModel.CreateOrderRequest) (orderServiceModel.CreateOrderResponse, error)
	PayOrder(ctx context.Context,req *orderServiceModel.PayOrderRequest, orderUUID uuid.UUID) (orderServiceModel.PayOrderResponse, error)
	GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderServiceModel.OrderDto, error)
	CancelOrder(ctx context.Context, orderUUID uuid.UUID) ( error)
	NewError(ctx context.Context, err error) *orderServiceModel.GenericErrorStatusCode 
}
