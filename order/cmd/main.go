package main

import (
	"sync"
	"time"
	""
)

const (
	httpPort = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout = 10 * time.Second
)

type OrderStorage struct{
	mu sync.RWMutex
	Orders map[string]OrderDto
}

func NewOrderStorage() *OrderStorage{
	return &OrderStorage{
		Orders: make(map[string]*),
	}
}

func main(){
	GetOrderByUUID()
}