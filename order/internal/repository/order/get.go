package order

import (
	"context"


	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

func (s *OrderRepo) GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool) {
	s.mu.Lock()
	order, ok := s.orders[orderUUID.String()]
	s.mu.Unlock()

	return order, ok
}
