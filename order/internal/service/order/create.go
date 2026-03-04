package order

import (
	"context"
	"log"

	"github.com/PhilSuslov/homework/order/internal/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CreateOrder(ctx context.Context, request *model.CreateOrderRequest) (model.CreateOrderResponse, error) {
	if request.UserUUID == uuid.Nil || len(request.UserUUID) == 0 {
		logger.Error(ctx, "failed to CreateOrder. UserUUID == Nil or len == 0",
			zap.String("UserUUID", request.UserUUID.String()))
		return model.CreateOrderResponse{}, status.Errorf(codes.Internal, "inventory error")
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
		return model.CreateOrderResponse{}, status.Errorf(codes.Internal, "inventory error: %v", err)
	}

	// 2. Проверяем, что все детали найдены
	if len(partsResp.Parts) != len(request.PartUuids) {
		logger.Error(ctx, "Not all parts found: ", zap.Int("expected= ", len(request.PartUuids)),
			zap.Int("got", len(partsResp.Parts)))
		return model.CreateOrderResponse{}, status.Error(codes.NotFound, "some parts not found")
	}

	// 3. Считаем цену
	var total_price float64
	for _, p := range partsResp.Parts {
		logger.Info(ctx, "Parts: ", zap.String("UUID= ", p.Uuid),
			zap.Float64("Price= ", p.Price))
		total_price += p.Price
	}
	logger.Info(ctx, "Total price calculated: ", zap.Float64("total_price", total_price))

	orderUUID := uuid.New()

	order := &model.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   request.UserUUID,
		PartUuids:  request.PartUuids,
		TotalPrice: total_price,
		Status:     model.OrderStatusPENDINGPAYMENT,
	}
	logger.Info(ctx, "order is CreateOrder Service:", zap.Any("order", order))
	
	err = s.orderService.CreateOrder(ctx, orderRepoConv.OrderDtoToRepo(order))
	if err != nil{
		log.Print("Ошибка в OrderService.CreateOrder", err)
	}

	event := model.OrderRecordedEvent{
		Event_uuid: uuid.NewString(),
		Order_uuid: orderUUID.String(),
		User_uuid: uuid.NewString(),
		Payment_method: "PENDING_PAYMENT",
		Transaction_uuid: uuid.NewString(),
	}

	err = s.orderProducerService.ProducerOrderRecorded(ctx, event)

	if err != nil{
		log.Printf("Ошибка в OrderProducerService", err)
		return model.CreateOrderResponse{}, err
	}

	log.Printf("Нет ошибки в OrderProducerService")

	// if err := telegramService.SendUFONotification(ctx, uuid, info); err != nil {
	// 	return "", err
	// }
	
	return model.CreateOrderResponse{
		OrderUUID:  order.OrderUUID,
		TotalPrice: total_price,
	}, nil
}
