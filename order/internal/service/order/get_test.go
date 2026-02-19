package order
// 
// import (
// 	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
// 	"github.com/google/uuid"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )
// 
// // ✅ Успешный кейс
// func (s *ServiceSuite) TestGetOrderByUUID_Success() {
// 	orderUUID := uuid.New()
// 	order := &orderRepoModel.OrderDto{
// 		OrderUUID: orderUUID,
// 		UserUUID:  uuid.New(),
// 		// можно добавить другие поля, если нужно
// 	}
// 
// 	// Настраиваем мок: возвращаем order и true
// 	s.serviceOrder.
// 		On("GetOrderByUUID", s.ctx, orderUUID).
// 		Return(order, true)
// 
// 	// Вызываем сервис
// 	res, err := s.service.GetOrderByUUID(s.ctx, orderUUID)
// 
// 	// Проверки
// 	s.Require().NoError(err)
// 	s.Require().NotNil(res)
// 	s.Equal(order.OrderUUID, res.OrderUUID)
// 	s.Equal(order.UserUUID, res.UserUUID)
// }
// 
// // ❌ Ошибка: order не найден
// func (s *ServiceSuite) TestGetOrderByUUID_NotFound() {
// 	orderUUID := uuid.New()
// 
// 	// Настраиваем мок: order отсутствует
// 	s.serviceOrder.
// 		On("GetOrderByUUID", s.ctx, orderUUID).
// 		Return((*orderRepoModel.OrderDto)(nil), false)
// 
// 	res, err := s.service.GetOrderByUUID(s.ctx, orderUUID)
// 
// 	s.Nil(res)
// 	s.Error(err)
// 
// 	st, ok := status.FromError(err)
// 	s.True(ok)
// 	s.Equal(codes.NotFound, st.Code())
// 	s.Equal("order not found", st.Message())
// }