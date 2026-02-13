package order

import (
	"context"

	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
)

func (s *OrderService) NewError(ctx context.Context, err error) *orderServiceModel.GenericErrorStatusCode {
	// Тут можно возвращать любую реализацию ErrorRes, например:
	var Err orderServiceModel.GenericErrorStatusCode
	Err.StatusCode = 500
	Err.Response.Code.Value = 500
	Err.Response.Message.Value = err.Error()
	return &orderServiceModel.GenericErrorStatusCode{
		StatusCode: Err.StatusCode,
		Response:   Err.Response,
	}
}
