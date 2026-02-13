package order

import (
	"context"
	"log"

	orderModel "github.com/PhilSuslov/homework/order/internal/model"

	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/google/uuid"
)

func (s *OrderService) GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderModel.OrderDto, error) {

	order, ok := s.orderService.GetOrderByUUID(ctx, orderUUID) 
	if !ok {
		log.Printf("order.OrderUUID - %v. Order not found in map!", order.OrderUUID)
		return nil, status.Error(codes.NotFound, "order not found")
	}

	ans := orderRepoConv.OrderDtoToService(order)

	return ans, nil

}
