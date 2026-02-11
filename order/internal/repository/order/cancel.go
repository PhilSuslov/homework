// В оригинале был update.go, но я не понял как его использовать и для реализации интерфейса нам
// нужен метод Cancel

package order

import (
	"context"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

func (s *OrderRepo) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (*orderV1.OrderDto, bool) {
	s.mu.Lock()
	order, ok := s.orders[params.OrderUUID.String()]
	s.mu.Unlock()
	return order, ok
}
