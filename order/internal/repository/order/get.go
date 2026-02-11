package order

import (
	"context"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

func (s *OrderRepo) GetOrderByUUID(_ context.Context, params orderV1.GetOrderByUUIDParams) (*orderV1.OrderDto, bool) {
	s.mu.Lock()
	order, ok := s.orders[params.OrderUUID.String()]
	s.mu.Unlock()

	return order, ok
}
