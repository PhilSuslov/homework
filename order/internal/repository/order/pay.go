package order

import (
	"context"
	"log"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	orderErr "github.com/PhilSuslov/homework/order/internal/model"

	"github.com/google/uuid"
)

func (s *OrderRepo) PayOrderCreate(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (*orderV1.OrderDto, bool) {
	s.mu.Lock()
	order, ok := s.orders[params.OrderUUID.String()]
	s.mu.Unlock()
	return order, ok
}

func (s *OrderRepo) PayOrder(orderUUID string, transactionuuid uuid.UUID, paymentMethod orderV1.PaymentMethod) (*orderV1.PayOrderResponse, error) {
	s.mu.Lock()
	order, ok := s.orders[orderUUID]
	if !ok {
		return nil, orderErr.ErrNotFound
	}

	order.Status = orderV1.OrderStatusPAID
	order.TransactionUUID.Value = transactionuuid
	order.PaymentMethod.Value = paymentMethod
	log.Println("=== Статус оплаты должен быть ===", s.orders)
	s.mu.Unlock()

	resp := &orderV1.PayOrderResponse{TransactionUUID: transactionuuid}

	return resp, nil

}
