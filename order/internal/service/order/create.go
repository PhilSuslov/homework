package order

import (
	"context"

	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CreateOrder(ctx context.Context, request *orderServiceModel.CreateOrderRequest) (orderServiceModel.CreateOrderResponse, error) {
	if request.UserUUID == uuid.Nil || len(request.UserUUID) == 0 {
		logger.Error(ctx, "failed to CreateOrder. UserUUID == Nil or len == 0",
			zap.String("UserUUID", request.UserUUID.String()))
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
		logger.Error(ctx, "Inventory ListParts error", zap.Error(err))
		return orderServiceModel.CreateOrderResponse{}, status.Errorf(codes.Internal, "inventory error: %v", err)
	}

	// log.Printf("Inventory ListParts response: %+v", partsResp)

	// 2. Проверяем, что все детали найдены
	if len(partsResp.Parts) != len(request.PartUuids) {
		logger.Error(ctx, "Not all parts found: ", zap.Int("expected= ", len(request.PartUuids)),
			zap.Int("got", len(partsResp.Parts)))
		// log.Printf("Not all parts found: expected=%d, got=%d", len(request.PartUuids), len(partsResp.Parts))
		return orderServiceModel.CreateOrderResponse{}, status.Error(codes.NotFound, "some parts not found")
	}

	// 3. Считаем цену
	var total_price float64
	for _, p := range partsResp.Parts {
		logger.Info(ctx, "Parts: ", zap.String("UUID= ", p.Uuid),
			zap.Float64("Price= ", p.Price))
		// log.Printf("Part: UUID=%s, Price=%f", p.Uuid, p.Price)
		total_price += p.Price
	}
	logger.Info(ctx, "Total price calculated: ", zap.Float64("total_price", total_price))
	// log.Printf("Total price calculated: %f", total_price)

	orderUUID := uuid.New()

	order := &orderServiceModel.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   request.UserUUID,
		PartUuids:  request.PartUuids,
		TotalPrice: total_price,
		Status:     orderServiceModel.OrderStatusPENDINGPAYMENT,
	}
	logger.Info(ctx, "order is CreateOrder Service:", zap.Any("order", order))
	// log.Println("order is CreateOrder Service",order)

	s.orderService.CreateOrder(ctx, orderRepoConv.OrderDtoToRepo(order))

	return orderServiceModel.CreateOrderResponse{
		OrderUUID:  order.OrderUUID,
		TotalPrice: total_price,
	}, nil

}
