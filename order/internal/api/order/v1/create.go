package v1

import (
	"context"

	orderConv "github.com/PhilSuslov/homework/order/internal/converter"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

func (s *OrderHandler) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	create, err := s.service.CreateOrder(ctx, orderConv.CreateOrderRequestToModel(req))
	return orderConv.CreateOrderResponseToOgen(&create), err

}
