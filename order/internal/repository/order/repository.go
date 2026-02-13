package order

import (
	"sync"

	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	def "github.com/PhilSuslov/homework/order/internal/repository"

)

var _ def.OrderRepository = (*OrderRepo)(nil)

type OrderRepo struct {
	mu     sync.RWMutex
	orders map[string]*orderRepoModel.OrderDto
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orders:          make(map[string]*orderRepoModel.OrderDto),
	}
}
