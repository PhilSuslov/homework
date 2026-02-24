package order

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

func (r *OrderRepo) CreateOrder(ctx context.Context, order *orderRepoModel.OrderDto) error {
	if r.conn == nil {
		return fmt.Errorf("pgx pool is nil")
	}

	if order.OrderUUID == uuid.Nil {
		order.OrderUUID = uuid.New()
	}
	if order.UserUUID == uuid.Nil {
		order.UserUUID = uuid.New()
	}
	if len(order.PartUuids) == 0 {
		order.PartUuids = []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	}
	if order.TotalPrice == 0 {
		order.TotalPrice = rand.Float64()
	}

	// Проверка соединения
	if err := r.conn.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	// Конвертируем массив UUID для pgx

	res, err := r.conn.Exec(ctx, `
        INSERT INTO orders (order_uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `, order.OrderUUID, order.UserUUID, order.PartUuids, order.TotalPrice,
		order.TransactionUUID.Value, order.PaymentMethod.Value, order.Status)

	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	log.Printf("Created %d rows\n", res.RowsAffected())
	return nil
}
