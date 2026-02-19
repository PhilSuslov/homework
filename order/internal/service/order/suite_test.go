package order

// import (
// 	"context"
// 	"testing"
// 
// 	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"
// 	mockRepo "github.com/PhilSuslov/homework/order/internal/repository/mocks"
// 	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
// 	// payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/google/uuid"
// 	grpc "google.golang.org/grpc"
// )
// 
// type ServiceSuite struct {
// 	suite.Suite
// 	ctx           context.Context
// 	serviceOrder  *mockRepo.OrderRepository
// 	service       *OrderService
// 	inventoryClient *fakeInventoryClient
// 	paymentClient   *fakePaymentClient
// 	fakeOrderSvc *fakeOrderService
// 	fakePayment *fakePaymentClient
// }
// 
// func (s *ServiceSuite) SetupTest() {
//     s.ctx = context.Background()
// 
//     // Мок репозитория
//     s.serviceOrder = new(mockRepo.OrderRepository)
//     s.serviceOrder.On("CreateOrder", mock.AnythingOfType("*model.OrderDto")).Return(nil)
// 
//     // Фейки
//     s.inventoryClient = &fakeInventoryClient{}
//     s.paymentClient = &fakePaymentClient{}
// 
//     // Прямо передаем их в сервис
//     s.service = &OrderService{
//         inventoryClient: s.inventoryClient,
//         // paymentClient:   s.paymentClient,
//         orderService:      s.serviceOrder, // убедиться, что поле совпадает с конструктором
//     }
// }
// 
// func TestServiceIntegration(t *testing.T) {
// 	suite.Run(t, new(ServiceSuite))
// }
// 
// // -----------------------
// // fake Inventory
// // -----------------------
// type fakeInventoryClient struct{}
// 
// func (f *fakeInventoryClient) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest, opts ...grpc.CallOption) (*inventoryV1.ListPartsResponse, error) {
// 	parts := make([]*inventoryV1.Part, len(req.Filter.Uuids))
// 	for i, u := range req.Filter.Uuids {
// 		parts[i] = &inventoryV1.Part{
// 			Uuid:  u,
// 			Price: float64(i + 1),
// 		}
// 	}
// 	return &inventoryV1.ListPartsResponse{Parts: parts}, nil
// }
// 
// func (f *fakeInventoryClient) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest, opts ...grpc.CallOption) (*inventoryV1.GetPartResponse, error) {
// 	return &inventoryV1.GetPartResponse{
// 		Part: &inventoryV1.Part{
// 			Uuid:  req.Uuid,
// 			Price: 100,
// 		},
// 	}, nil
// }
// 
// // -----------------------
// // fake Payment
// // -----------------------
// // type fakePaymentClient struct{}
// // 
// // func (f *fakePaymentClient) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest, opts ...grpc.CallOption) (*payment_v1.PayOrderResponse, error) {
// // 	return &payment_v1.PayOrderResponse{
// // 		TransactionUuid: "fake-tx-uuid",
// // 	}, nil
// // }
// 
// // // Фейковые клиенты
// // type fakePaymentClient struct{}
// // 
// // func (f *fakePaymentClient) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest, opts ...interface{}) (*payment_v1.PayOrderResponse, error) {
// // 	return &payment_v1.PayOrderResponse{
// // 		TransactionUuid: "11111111-1111-1111-1111-111111111111",
// // 	}, nil
// // }
// 
// type fakeOrderService struct {
// 	mock.Mock
// }
// 
// func (f *fakeOrderService) PayOrderCreate(ctx context.Context, req interface{}, orderUUID uuid.UUID) (*orderRepoModel.OrderDto, bool) {
// 	args := f.Called(ctx, req, orderUUID)
// 	return args.Get(0).(*orderRepoModel.OrderDto), args.Bool(1)
// }
// 
// func (f *fakeOrderService) PayOrder(orderUUID, transactionUUID uuid.UUID, paymentMethod string) (*orderRepoModel.PayOrderResponse, error) {
// 	args := f.Called(orderUUID, transactionUUID, paymentMethod)
// 	return args.Get(0).(*orderRepoModel.PayOrderResponse), args.Error(1)
// }
// 
