package order

import (
	"context"
	"log"

	// orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CreateOrder(ctx context.Context, request *orderServiceModel.CreateOrderRequest) (orderServiceModel.CreateOrderResponse, error) {
	if request.UserUUID == uuid.Nil || len(request.UserUUID) == 0 {
		return orderServiceModel.CreateOrderResponse{}, status.Errorf(codes.Internal, "inventory error")
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
		return orderServiceModel.CreateOrderResponse{}, status.Errorf(codes.Internal, "inventory error: %v", err)
	}

	log.Printf("Inventory ListParts response: %+v", partsResp)

	// 2. Проверяем, что все детали найдены
	if len(partsResp.Parts) != len(request.PartUuids) {
		log.Printf("Not all parts found: expected=%d, got=%d", len(request.PartUuids), len(partsResp.Parts))
		return orderServiceModel.CreateOrderResponse{}, status.Error(codes.NotFound, "some parts not found")
	}

	// 3. Считаем цену
	var total_price float64
	for _, p := range partsResp.Parts {
		log.Printf("Part: UUID=%s, Price=%f", p.Uuid, p.Price)
		total_price += p.Price
	}
	log.Printf("Total price calculated: %f", total_price)

	orderUUID := uuid.New()

	order := &orderServiceModel.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   request.UserUUID,
		PartUuids:  request.PartUuids,
		TotalPrice: total_price,
		Status:     orderServiceModel.OrderStatusPENDINGPAYMENT,
	}

	// ans := orderRepoConv.OrderDtoToService(order)
	s.orderService.CreateOrder(orderRepoConv.OrderDtoToRepo(order))

	return orderServiceModel.CreateOrderResponse{
		OrderUUID:  order.OrderUUID,
		TotalPrice: total_price,
	}, nil

}
