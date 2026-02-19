package order

import (
	"context"
	"log"

	// orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"

	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"

	"github.com/google/uuid"
)

func (s *OrderRepo) PayOrderCreate(ctx context.Context, req *orderRepoModel.PayOrderRequest, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool) {
	var order orderRepoModel.OrderDto
	// s.mu.Lock()
	// order, ok := s.orders[orderUUID.String()]
	// s.mu.Unlock()
	row := s.conn.QueryRow(ctx, `SELECT order_uuid, user_uuid, part_uuids, total_price,
	transaction_uuid, payment_method, status FROM orders WHERE order_uuid = $1`, orderUUID.String())
	// if err != nil{
	// 	log.Printf("failed to connect PayOrderCreate: %v\n", err)
	// }
	err := row.Scan(&order.OrderUUID, &order.UserUUID, &order.PartUuids, &order.TotalPrice,
		&order.TransactionUUID.Value, &order.PaymentMethod.Value, &order.Status)

	if err != nil {
		log.Printf("failed to scan order: %v\n", err)
		return &order, false
	}
	return &order, true
}

func (s *OrderRepo) PayOrder(ctx context.Context, orderUUID uuid.UUID, uuidPay uuid.UUID, paymentMethod string) (*string, error) {
	var order orderRepoModel.OrderDto
	log.Println("--------------------------------------------")
	row := s.conn.QueryRow(ctx, `SELECT order_uuid, user_uuid, part_uuids, total_price,
	transaction_uuid, payment_method, status FROM orders WHERE order_uuid = $1`, orderUUID.String())
	err := row.Scan(&order.OrderUUID, &order.UserUUID, &order.PartUuids, &order.TotalPrice,
		&order.TransactionUUID.Value, &order.PaymentMethod.Value, &order.Status)

	if err != nil {
		log.Printf("failed to scan order: %v\n", err)
		// resp := order.TransactionUUID.Value.String()
		return nil, err
	}
	order.Status = orderRepoModel.OrderStatusPAID
	log.Println("--------------------------------------------")
	_, err = s.conn.Exec(ctx, "UPDATE orders SET status = 'PAID' WHERE order_uuid = $1", orderUUID.String())
	if err != nil{
		log.Printf("failed to update status paid: %v\n", err)
		return nil, err
	}
	log.Println("--------------------------------------------")
	order.TransactionUUID.Value, err = uuid.Parse(uuidPay.String())
	order.PaymentMethod.Value = orderRepoConv.PaymentMethodToRepo(paymentMethod)
	resp := order.TransactionUUID.Value.String()
	return &resp, err

	// 	s.mu.Lock()
	// 	order, ok := s.orders[orderUUID.String()]
	// 	if !ok {
	// 		return nil, orderErr.ErrNotFound
	// 	}
	// 	var err error
	//
	// 	order.Status = orderRepoModel.OrderStatusPAID
	// 	order.TransactionUUID.Value, err = uuid.Parse(uuidPay.String())
	// 	order.PaymentMethod.Value = orderRepoConv.PaymentMethodToRepo(paymentMethod)
	// 	log.Println("=== Статус оплаты должен быть ===", s.orders)
	// 	s.mu.Unlock()
	//
	// 	resp := order.TransactionUUID.Value.String()
	//
	// 	return &resp, err

}
