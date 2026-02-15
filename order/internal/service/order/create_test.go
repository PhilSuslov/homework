package order

import (
	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
	"github.com/stretchr/testify/mock"

	"github.com/google/uuid"
)
func (s *ServiceSuite) TestCreateOrder_WithoutInventory() {
	userUUID := uuid.New()
	partUUIDs := []uuid.UUID{uuid.New(), uuid.New()}

	req := &orderServiceModel.CreateOrderRequest{
		UserUUID:  userUUID,
		PartUuids: partUUIDs,
	}

	// Подменяем inventoryClient на заглушку
	s.service.inventoryClient = &fakeInventoryClient{}

	// Мок репозитория: возвращаем заранее подготовленный OrderDto
	createdOrder := &orderServiceModel.OrderDto{
		OrderUUID: uuid.New(),
		UserUUID:  userUUID,
		PartUuids: partUUIDs,
		Status:    orderServiceModel.OrderStatusPENDINGPAYMENT,
		TotalPrice: 3.0, // сумма фейковых частей
	}

	s.serviceOrder.On("CreateOrder", mock.Anything).Return(createdOrder, nil)

	resp, err := s.service.CreateOrder(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.NotEqual(uuid.Nil, resp.OrderUUID)

	// Проверка созданного заказа
	s.NotNil(createdOrder)
	s.Equal(userUUID, createdOrder.UserUUID)
	s.Equal(partUUIDs, createdOrder.PartUuids)
	s.Equal(orderServiceModel.OrderStatusPENDINGPAYMENT, createdOrder.Status)
	s.Equal(3.0, createdOrder.TotalPrice)
}