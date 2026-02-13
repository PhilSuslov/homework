package order

import (
	"context"
	"log"

	// orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	orderErr "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"

	"github.com/google/uuid"
)

func (s *OrderRepo) PayOrderCreate(ctx context.Context, req *orderRepoModel.PayOrderRequest, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool) {
	s.mu.Lock()
	order, ok := s.orders[orderUUID.String()]
	s.mu.Unlock()
	return order, ok
}

func (s *OrderRepo) PayOrder(orderUUID uuid.UUID, uuidPay uuid.UUID, paymentMethod string) (*string, error) {
	s.mu.Lock()
	order, ok := s.orders[orderUUID.String()]
	if !ok {
		return nil, orderErr.ErrNotFound
	}
	var err error

	order.Status = orderRepoModel.OrderStatusPAID
	order.TransactionUUID.Value, err = uuid.Parse(uuidPay.String())
	order.PaymentMethod.Value = orderRepoConv.PaymentMethodToRepo(paymentMethod)
	log.Println("=== Статус оплаты должен быть ===", s.orders)
	s.mu.Unlock()

	resp := order.TransactionUUID.Value.String()
	// resp := &orderRepoModel.PayOrderResponse{TransactionUUID: order.TransactionUUID.Value }

	return &resp, err

}
