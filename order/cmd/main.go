package main

import (
	"context"
	"sync"
	"time"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"


	"github.com/google/uuid"
)

const (
	httpPort = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout = 10 * time.Second
)

type OrderStorage struct{
	mu sync.RWMutex
	orders map[string]*orderV1.OrderDto
}

type OrderHandler struct {
	storage *OrderStorage
}

type CreateOrderOK struct {
	OrderUUID  string  `json:"order_uuid"`
	TotalPrice float64 `json:"total_price"`
}

func (CreateOrderOK) CreateOrderRes() {}

func NewOrderStorage() *OrderStorage{
	return &OrderStorage{
		orders: make(map[string]*orderV1.OrderDto),
	}
}

func NewOrderHandler(storage *OrderStorage) *OrderHandler{
	return &OrderHandler{
		storage: storage,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, request *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error){
	h.storage.mu.Lock()
	defer h.storage.mu.Unlock()

	orderUUID := uuid.New()

	order := &orderV1.OrderDto{
		OrderUUID: orderUUID,
		UserUUID: request.UserUUID,
		PartUuids: request.PartUuids,
		TotalPrice: 1443, //request., InventoryService
		// TransactionUUID:, Тут пока хз где брать информацию
    	// PaymentMethod: ,
    	// Status: ,
	}

	if _, ok := h.storage.orders[orderUUID.String()]; ok{
		return &orderV1.BadRequestError{
			Code: 400, Message: "Bad request",
		}, nil
	}
	h.storage.orders[orderUUID.String()] = order
	// return &CreateOrderOK{
	// 	OrderUUID: orderUUID,
	// 	TotalPrice: order.TotalPrice,
	// }, nil
}

func (o *OrderStorage) GetOrderByUUID(uuid string) *orderV1.OrderDto{
	o.mu.RLock()
	defer o.mu.RUnlock()

	order,ok := o.orders[uuid]
	if !ok{
		return nil
	}
	return order
}


func main(){
	storage := NewOrderStorage()

	orderHandler := NewOrderHandler(storage)

	orderService, err := orderV1.NewServer(orderHandler)
}