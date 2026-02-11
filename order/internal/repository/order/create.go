package order

import order_v1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"

func (r *OrderRepo) CreateOrder(order *order_v1.OrderDto){
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.OrderUUID.String()] = order
}

