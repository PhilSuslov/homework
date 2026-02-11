package order

import (
	"context"
	"log"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CreateOrder(ctx context.Context, request *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	if request.UserUUID == uuid.Nil || len(request.UserUUID) == 0 {
		return nil, status.Errorf(codes.Internal, "inventory error")
	}

	strUUID := make([]string, len(request.PartUuids))
	for i, v := range request.PartUuids {
		strUUID[i] = v.String()
	}

	//1. Запрашиваем детали из Inventory
	partsResp, err := s.inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{
		Filter: &inventoryV1.PartsFilter{
			Uuids: strUUID,
		},
	})
	if err != nil {
		log.Printf("Inventory ListParts error: %v", err)
		return nil, status.Errorf(codes.Internal, "inventory error: %v", err)
	}

	log.Printf("Inventory ListParts response: %+v", partsResp)

	// 2. Проверяем, что все детали найдены
	if len(partsResp.Parts) != len(request.PartUuids) {
		log.Printf("Not all parts found: expected=%d, got=%d", len(request.PartUuids), len(partsResp.Parts))
		return nil, status.Error(codes.NotFound, "some parts not found")
	}

	// 3. Считаем цену
	var total_price float64
	for _, p := range partsResp.Parts {
		log.Printf("Part: UUID=%s, Price=%f", p.Uuid, p.Price)
		total_price += p.Price
	}
	log.Printf("Total price calculated: %f", total_price)

	orderUUID := uuid.New()

	order := &orderV1.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   request.UserUUID,
		PartUuids:  request.PartUuids,
		TotalPrice: total_price,
		Status:     orderV1.OrderStatusPENDINGPAYMENT,
	}

	s.orderService.CreateOrder(order)

	return &orderV1.CreateOrderResponse{
		OrderUUID:  order.OrderUUID,
		TotalPrice: order.TotalPrice,
	}, nil

}
