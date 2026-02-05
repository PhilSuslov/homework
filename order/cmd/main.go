package main

import (
	"sync"
	"time"
	"github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

const (
	httpPort = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout = 10 * time.Second
)

type OrderStorage struct{
	mu sync.RWMutex
	Orders map[string]*orderV1.OrderDto
}

func NewOrderStorage() *OrderStorage{
	return &OrderStorage{
		Orders: make(map[string]*orderV1.OrderDto),
	}
}

func (o *OrderStorage) GetOrderByUUID(uuid string) *orderV1.OrderDto{
	o.mu.RLock()
	defer o.mu.RUnlock()

	order,ok := o.Orders[uuid]
	if !ok{
		return nil
	}
	return order
}

func main(){
	GetOrderByUUID()
}