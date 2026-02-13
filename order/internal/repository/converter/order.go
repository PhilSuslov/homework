package converter

import (
	// "github.com/google/uuid"
	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
	"github.com/google/uuid"
)

func PayOrderRequestToService(repo *orderRepoModel.PayOrderRequest) *orderServiceModel.PayOrderRequest {
	return &orderServiceModel.PayOrderRequest{
		OrderUUID:     repo.OrderUUID,
		PaymentMethod: orderServiceModel.PaymentMethod(repo.PaymentMethod),
	}
}

func PayOrderRequestToRepo(service *orderServiceModel.PayOrderRequest) *orderRepoModel.PayOrderRequest {
	return &orderRepoModel.PayOrderRequest{
		OrderUUID:     service.OrderUUID,
		PaymentMethod: orderRepoModel.PaymentMethod(service.PaymentMethod),
	}
}

func PayOrderParamsToService(repo *orderRepoModel.PayOrderParams) *orderServiceModel.PayOrderParams {
	return &orderServiceModel.PayOrderParams{
		OrderUUID: repo.OrderUUID,
	}
}

func PayOrderParamsToRepo(service uuid.UUID) orderRepoModel.PayOrderParams {
	return orderRepoModel.PayOrderParams{
		OrderUUID: service,
	}
}

func OrderStatusToRepo(service orderServiceModel.OrderStatus) orderRepoModel.OrderStatus{
	return orderRepoModel.OrderStatus(service)
}

func OrderStatusToService(repo orderRepoModel.OrderStatus) orderServiceModel.OrderStatus{
	return orderServiceModel.OrderStatus(repo)
}

func PaymentMethodToRepo(service string) orderRepoModel.PaymentMethod{
	return orderRepoModel.PaymentMethod(service)
}

func PaymentMethodToService(repo string) orderServiceModel.PaymentMethod{
	return orderServiceModel.PaymentMethod(repo)
}

func PayOrderResponseToService(repo string) orderServiceModel.PayOrderResponse{
	repoUUID, _ := uuid.Parse(repo)
	return orderServiceModel.PayOrderResponse{
		TransactionUUID: repoUUID,
	}
}

func PayOrderResponseToRepo(service *orderServiceModel.PayOrderResponse) orderRepoModel.PayOrderResponse{
	return orderRepoModel.PayOrderResponse{
		TransactionUUID: service.TransactionUUID,
	}
}

func GetOrderByUUIDToService(repo *orderRepoModel.GetOrderByUUIDParams) orderServiceModel.GetOrderByUUIDParams {
	return orderServiceModel.GetOrderByUUIDParams{
		OrderUUID: repo.OrderUUID,
	}
}

func GetOrderByUUIDToRepo(service *orderServiceModel.GetOrderByUUIDParams) orderRepoModel.GetOrderByUUIDParams {
	return orderRepoModel.GetOrderByUUIDParams{
		OrderUUID: service.OrderUUID,
	}
}

func OrderDtoToRepo(service *orderServiceModel.OrderDto) *orderRepoModel.OrderDto {
	return &orderRepoModel.OrderDto{
		OrderUUID:       service.OrderUUID,
		UserUUID:        service.UserUUID,
		PartUuids:       service.PartUuids,
		TotalPrice:      service.TotalPrice,
		TransactionUUID: orderRepoModel.OptNilString(service.TransactionUUID),
		PaymentMethod:   OptPaymentMethodToRepo(service.PaymentMethod),
		Status:          orderRepoModel.OrderStatus(service.Status),
	}
}

func OptPaymentMethodToRepo(service orderServiceModel.OptPaymentMethod) orderRepoModel.OptPaymentMethod{
	return orderRepoModel.OptPaymentMethod{
		Value: orderRepoModel.PaymentMethod(service.Value),
		Set: service.Set,
	}
}

func OrderDtoToService(repo *orderRepoModel.OrderDto) *orderServiceModel.OrderDto {
	return &orderServiceModel.OrderDto{
		OrderUUID:       repo.OrderUUID,
		UserUUID:        repo.UserUUID,
		PartUuids:       repo.PartUuids,
		TotalPrice:      repo.TotalPrice,
		TransactionUUID: orderServiceModel.OptNilString(repo.TransactionUUID),
		PaymentMethod:   OptPaymentMethodToService(repo.PaymentMethod),
		Status:          orderServiceModel.OrderStatus(repo.Status),
	}
}

func OptPaymentMethodToService(repo orderRepoModel.OptPaymentMethod) orderServiceModel.OptPaymentMethod{
	return orderServiceModel.OptPaymentMethod{
		Value: orderServiceModel.PaymentMethod(repo.Value),
		Set: repo.Set,
	}
}
