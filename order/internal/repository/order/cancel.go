// В оригинале был update.go, но я не понял как его использовать и для реализации интерфейса нам
// нужен метод Cancel

package order

import (
	"context"
	"log"

	// orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

func (s *OrderRepo) CancelOrder(ctx context.Context, orderUUID uuid.UUID) (bool) {
	_,err := s.conn.Exec(ctx, "UPDATE orders SET status = 'CANCELLED' WHERE order_uuid = $1", orderUUID.String())
	if err != nil{
		log.Printf("failed to scan in CancelOrder: %v\n", err)
		return false
	}
	return true
}
