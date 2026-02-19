package order

import (
	"context"
	"log"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
)

func (r *OrderRepo) CreateOrder(ctx context.Context, order *orderRepoModel.OrderDto) {
	// r.mu.Lock()
	// defer r.mu.Unlock()
	// r.orders[order.OrderUUID.String()] = order
	err := r.conn.Ping(ctx)
	if err != nil {
		log.Println("failed to ping is order -> CreateOrder: ", err)
		return
	}
	res, err := r.conn.Exec(ctx, `INSERT INTO orders (order_uuid, user_uuid, part_uuids, total_price, 
	transaction_uuid, payment_method, status) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		order.OrderUUID, order.UserUUID, order.PartUuids, order.TotalPrice, order.TransactionUUID.Value,
		order.PaymentMethod.Value, order.Status)

	if err != nil {
		log.Printf("failed to CreateOrder: %v\n", err)
	}

	log.Printf("Create %d rows\n", res.RowsAffected())

}
