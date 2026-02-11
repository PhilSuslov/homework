package order

import (
	"sync"

	def "github.com/PhilSuslov/homework/order/internal/repository"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"

)

var _ def.OrderRepository = (*OrderRepo)(nil)

type OrderRepo struct {
	mu     sync.RWMutex
	orders map[string]*orderV1.OrderDto
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orders:          make(map[string]*orderV1.OrderDto),
	}
}
