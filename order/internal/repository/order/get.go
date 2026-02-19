package order

import (
	"context"
	"log"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

func (s *OrderRepo) GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool) {
	// order, ok := s.db.[orderUUID.String()]
	var order orderRepoModel.OrderDto
	rows, err := s.conn.Query(ctx, `SELECT order_uuid, user_uuid, part_uuids, total_price,
	transaction_uuid, payment_method, status FROM orders WHERE order_uuid = $1`, orderUUID.String())
	if err != nil{
		log.Printf("failed to select orders by uuid: %v. Error: %v\n", orderUUID, err)
		return nil, false 
	}
	defer rows.Close()

	for rows.Next(){
		// log.Println(rows)
		err = rows.Scan(&order.OrderUUID, &order.UserUUID, &order.PartUuids, &order.TotalPrice, 
			&order.TransactionUUID.Value, &order.PaymentMethod.Value, &order.Status)
		if err != nil{
			log.Printf("failed to scan order: %v\n", err)
			return &order, false
		}
		log.Println(order.OrderUUID, order.UserUUID, order.PartUuids, order.TotalPrice, 
			order.TransactionUUID.Value, order.PaymentMethod.Value, order.Status)
	}
	return &order, true
}
