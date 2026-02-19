package order
// 
// import (
// 	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
// 	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
// 	"github.com/stretchr/testify/mock"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// 
// 	"github.com/google/uuid"
// )
// 
// func (s *ServiceSuite) TestCancelOrder_Success() {
// 	orderUUID := uuid.New()
// 
// 	order := &orderRepoModel.OrderDto{
// 		OrderUUID: orderUUID,
// 		Status:    orderRepoModel.OrderStatusCANCELLED,
// 	}
// 
// 	s.serviceOrder.
// 		On("CancelOrder", mock.Anything, orderUUID).
// 		Return(order, true)
// 
// 	err := s.service.CancelOrder(s.ctx, orderUUID)
// 
// 	s.Require().NoError(err)
// 	s.Equal(
// 		orderRepoModel.OrderStatus(orderServiceModel.OrderStatusCANCELLED),
// 		order.Status,
// 	)
// }
// func (s *ServiceSuite) TestCancelOrder_NotFound() {
// 	orderUUID := uuid.New()
// 
// 	s.serviceOrder.
// 		On("CancelOrder", mock.Anything, orderUUID).
// 		Return((*orderRepoModel.OrderDto)(nil), false)
// 
// 	err := s.service.CancelOrder(s.ctx, orderUUID)
// 
// 	s.Require().Error(err)
// 
// 	st, _ := status.FromError(err)
// 	s.Equal(codes.NotFound, st.Code())
// }
// 
// func (s *ServiceSuite) TestCancelOrder_PaidOrder() {
// 	orderUUID := uuid.New()
// 
// 	order := &orderRepoModel.OrderDto{
// 		OrderUUID: orderUUID,
// 		Status:    orderRepoModel.OrderStatus(orderServiceModel.OrderStatusPAID),
// 	}
// 
// 	s.serviceOrder.
// 		On("CancelOrder", mock.Anything, orderUUID).
// 		Return(order, true)
// 
// 	err := s.service.CancelOrder(s.ctx, orderUUID)
// 
// 	s.Require().Error(err)
// 
// 	st, _ := status.FromError(err)
// 	s.Equal(codes.Unknown, st.Code())
// }
