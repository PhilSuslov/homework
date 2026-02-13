package converter

import (
	"github.com/google/uuid"

	orderModel "github.com/PhilSuslov/homework/order/internal/model"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

func PayOrderRequestToModel(ogen *orderV1.PayOrderRequest) *orderModel.PayOrderRequest {
	return &orderModel.PayOrderRequest{
		OrderUUID:     ogen.OrderUUID,
		PaymentMethod: orderModel.PaymentMethod(ogen.PaymentMethod),
	}
}

func PayOrderParamsToModel(ogen orderV1.PayOrderParams) uuid.UUID {
	return ogen.OrderUUID
}

func PayOrderResponseToOgen(service orderModel.PayOrderResponse) orderV1.PayOrderResponse {
	return orderV1.PayOrderResponse{
		TransactionUUID: service.TransactionUUID,
	}
}

func GetOrderByUUIDParamsToModel(ogen orderV1.GetOrderByUUIDParams) uuid.UUID {
	return ogen.OrderUUID
}

func OrderDtoToOgen(service orderModel.OrderDto) orderV1.OrderDto {
	return orderV1.OrderDto{
		OrderUUID:       service.OrderUUID,
		UserUUID:        service.UserUUID,
		PartUuids:       service.PartUuids,
		TotalPrice:      service.TotalPrice,
		TransactionUUID: orderV1.OptNilString(service.TransactionUUID),
		PaymentMethod:   OptPaymentMethodToOgen(service.PaymentMethod),
		Status:          orderV1.OrderStatus(service.Status),
	}
}

func PaymentMethodToOgen(service orderModel.PaymentMethod) orderV1.PaymentMethod {
	return orderV1.PaymentMethod(service)
}

func OptPaymentMethodToOgen(service orderModel.OptPaymentMethod) orderV1.OptPaymentMethod {
	return orderV1.OptPaymentMethod{
		Value: orderV1.PaymentMethod(service.Value),
		Set:   service.Set,
	}
}

func CreateOrderRequestToOgen(service *orderModel.CreateOrderRequest) *orderV1.CreateOrderRequest {
	return &orderV1.CreateOrderRequest{
		UserUUID:  service.UserUUID,
		PartUuids: service.PartUuids,
	}
}

func CreateOrderRequestToModel(ogen *orderV1.CreateOrderRequest) *orderModel.CreateOrderRequest {
	return &orderModel.CreateOrderRequest{
		UserUUID:  ogen.UserUUID,
		PartUuids: ogen.PartUuids,
	}
}

func CreateOrderResponseToOgen(service *orderModel.CreateOrderResponse) *orderV1.CreateOrderResponse {
	return &orderV1.CreateOrderResponse{
		OrderUUID:  service.OrderUUID,
		TotalPrice: service.TotalPrice,
	}
}

func CreateOrderResponseToModel(ogen *orderV1.CreateOrderResponse) *orderModel.CreateOrderResponse {
	return &orderModel.CreateOrderResponse{
		OrderUUID:  ogen.OrderUUID,
		TotalPrice: ogen.TotalPrice,
	}
}

func NewErrToOgen(service *orderModel.GenericErrorStatusCode) *orderV1.GenericErrorStatusCode {
	return &orderV1.GenericErrorStatusCode{
		StatusCode: service.StatusCode,
		Response:   GenericErrToOgen(service.Response),
	}
}

func GenericErrToOgen(service orderModel.GenericError) orderV1.GenericError {
	return orderV1.GenericError{
		Code:    orderV1.OptInt(service.Code),
		Message: orderV1.OptString(service.Message),
	}
}
