package order

import (
	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
)


func (r *OrderRepo) CreateOrder(order *orderRepoModel.OrderDto){
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.OrderUUID.String()] = order
}

