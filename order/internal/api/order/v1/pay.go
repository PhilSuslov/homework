package v1

import (
	"context"
	"errors"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	orderConv "github.com/PhilSuslov/homework/order/internal/converter"

)

func (s *OrderHandler) PayOrder(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
	res, err := s.service.PayOrder(ctx, orderConv.PayOrderRequestToModel(req), orderConv.PayOrderParamsToModel(params))
	if err != nil {
		return nil, errors.New("Ошибка в PayOrder в API")
	}
	resConv := orderConv.PayOrderResponseToOgen(res)
	return &resConv, err 
}
