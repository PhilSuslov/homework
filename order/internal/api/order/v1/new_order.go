package v1

import (
	"context"
	orderService "github.com/PhilSuslov/homework/order/internal/service"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	orderConv "github.com/PhilSuslov/homework/order/internal/converter"
)

type OrderHandler struct {
    service orderService.OrderService
}


func NewOrderHandler(s orderService.OrderService) *OrderHandler {
    return &OrderHandler{service: s}
}

func convertHandlerToService (handler *OrderHandler) orderV1.Handler{
	return orderV1.UnimplementedHandler{}
}

func (s *OrderHandler) NewError(ctx context.Context, err error) *orderV1.GenericErrorStatusCode{
	res := s.service.NewError(ctx, err)
	return orderConv.NewErrToOgen(res)
}