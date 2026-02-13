package v1

import (
	"context"

	orderConv "github.com/PhilSuslov/homework/order/internal/converter"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

func (s *OrderHandler) GetOrderByUUID(ctx context.Context, params orderV1.GetOrderByUUIDParams) (orderV1.GetOrderByUUIDRes, error) {
	get, err := s.service.GetOrderByUUID(ctx, orderConv.GetOrderByUUIDParamsToModel(params))

	getConv := orderConv.OrderDtoToOgen(*get)
	return &getConv, err
}
