// В оригинале был update.go, но я не понял как его использовать и для реализации интерфейса нам
// нужен метод Cancel

package order

import (
	"context"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

func (s *OrderRepo) CancelOrder(ctx context.Context, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool) {
	s.mu.Lock()
	order, ok := s.orders[orderUUID.String()]
	s.mu.Unlock()
	return order, ok
}
