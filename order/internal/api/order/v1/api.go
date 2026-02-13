package v1

import (
	orderService "github.com/PhilSuslov/homework/order/internal/service"
)

type api struct {
	
	service orderService.OrderService
}

func NewAPI(apiService orderService.OrderService) *api {
	return &api{
		service: apiService,
	}
}
