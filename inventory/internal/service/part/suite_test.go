package part

// import (
// 	"context"
// 	"testing"
// 
// 	"github.com/stretchr/testify/suite"
// 
// 	mock "github.com/PhilSuslov/homework/inventory/internal/repository/mocks"
// )
// 
// type ServiceSuite struct {
// 	suite.Suite
// 
// 	ctx context.Context
// 
// 	serviceInventory *mock.InventoryService
// 
// 	service *service
// }
// 
// func (s *ServiceSuite) SetupTest() {
// 	s.ctx = context.Background()
// 
// 	s.serviceInventory = mock.NewInventoryService(s.T())
// 
// 	s.service = NewService(s.serviceInventory)
// }
// 
// func (s *ServiceSuite) TearDownTest() {
// }
// 
// func TestServiceIntegration(t *testing.T) {
// 	suite.Run(t, new(ServiceSuite))
// }
