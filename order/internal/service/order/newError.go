package order

import (
	"context"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

func (s *OrderService) NewError(ctx context.Context, err error) *orderV1.GenericErrorStatusCode {
	// Тут можно возвращать любую реализацию ErrorRes, например:
	var Err orderV1.GenericErrorStatusCode
	Err.StatusCode = 500
	Err.Response.Code.Value = 500
	Err.Response.Message.Value = err.Error()
	return &orderV1.GenericErrorStatusCode{
		StatusCode: Err.StatusCode,
		Response:   Err.Response,
	}
}
