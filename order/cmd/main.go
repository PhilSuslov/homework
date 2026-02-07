package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

const (
	httpPort          = "8080"
	// readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type OrderService struct {
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient

	mu     sync.RWMutex
	orders map[string]*orderV1.OrderDto
}

func NewOrderService(
	inventoryClient inventoryV1.InventoryServiceClient,
	paymentClient paymentV1.PaymentServiceClient) *OrderService {
	return &OrderService{
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
		orders:          make(map[string]*orderV1.OrderDto),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, request *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {

	if request.UserUUID == uuid.Nil || len(request.UserUUID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_uuid and part_uuid are required")
	}

	strUUID := make([]string, len(request.PartUuids))
	for i,v := range request.PartUuids{
		strUUID[i] = v.String()
	}

	//1. Запрашиваем детали из Inventory
	partsResp, err := s.inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{
		Filter: &inventoryV1.PartsFilter{
			Uuids: strUUID,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inventory error: %v", err)
	}

	// 2. Проверяем, что все детали найдены
	if len(partsResp.Parts) != len(request.PartUuids) {
		return nil, status.Error(codes.NotFound, "some parts not found")
	}

	// 3. Считаем цену
	var total_price float64
	for _, p := range partsResp.Parts {
		total_price += p.Price
	}

	orderUUID := uuid.New()

	order := &orderV1.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   request.UserUUID,
		PartUuids:  request.PartUuids,
		TotalPrice: total_price,
		Status:     orderV1.OrderStatusPENDINGPAYMENT,
	}

	s.mu.Lock()
	s.orders[orderUUID.String()] = order
	defer s.mu.Unlock()

	return &orderV1.CreateOrderResponse{
		OrderUUID:  order.OrderUUID,
		TotalPrice: order.TotalPrice,
	}, nil

}

// --------------------------------------------------
// -- PAY ORDER
// --------------------------------------------------

func (s *OrderService) PayOrder(ctx context.Context,
	request *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {

	s.mu.Lock()
	order, ok := s.orders[request.OrderUUID.String()]
	if !ok {
		s.mu.Unlock()
		return nil, status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderV1.OrderStatusPAID {
		s.mu.Unlock()
		return nil, status.Error(codes.Canceled, "order already paid")
	}

	if order.Status == orderV1.OrderStatusCANCELLED {
		s.mu.Unlock()
		return nil, status.Error(codes.Canceled, "order cancelled")
	}

	s.mu.Unlock()

	//Проверка метода оплаты
	var pm paymentV1.PaymentMethod
	switch request.PaymentMethod {
	case "CARD":
		pm = 1
	case "SBP":
		pm = 2
	case "CREDIT_CARD":
		pm = 3
	case "INVESTOR_MONEY":
		pm = 4
	default:
		pm = 0
	}
	// Вызываем PaymentService
	payResp, err := s.paymentClient.PayOrder(ctx, &paymentV1.PayOrderRequest{
		OrderUuid:     request.OrderUUID.String(),
		UserUuid:      order.UserUUID.String(),
		PaymentMethod: pm,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "payment error: %v", err)
	}

	transactionUUID := payResp.TransactionUuid
	paymentMethod := request.PaymentMethod

	transactionuuid, _ := uuid.Parse(transactionUUID)

	s.mu.Lock()
	order.Status = orderV1.OrderStatusPAID
	order.TransactionUUID.Value = transactionuuid
	order.PaymentMethod.Value = paymentMethod
	s.mu.Unlock()

	resp := &orderV1.PayOrderResponse{TransactionUUID: transactionuuid}

	return resp, nil

}

// --------------------------------------------------
// -- GET ORDER
// --------------------------------------------------

func (s *OrderService) GetOrderByUUID(ctx context.Context, params orderV1.GetOrderByUUIDParams) (orderV1.GetOrderByUUIDRes, error) {
	s.mu.Lock()
	order, ok := s.orders[params.OrderUUID.String()]
	s.mu.Unlock()

	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &orderV1.OrderDto{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          order.Status,
	}, nil
}

// --------------------------------------------------
// -- CANCEL ORDER
// --------------------------------------------------

func (s *OrderService) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	s.mu.Lock()
	order, ok := s.orders[params.OrderUUID.String()]
	if !ok {
		s.mu.Unlock()
		return nil, status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderV1.OrderStatusPAID {
		s.mu.Unlock()
		return nil, status.Error(codes.Unknown, "Conflict")
	}

	order.Status = orderV1.OrderStatusCANCELLED
	s.mu.Unlock()
	return nil, status.Error(codes.Canceled, "No content")
}

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

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ---------------- gRPC connections ----------------

	invConn, err := grpc.NewClient(
		"dns:///localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to inventory: %v", err)
	}
	defer invConn.Close()

	payConn, err := grpc.NewClient(
		"dns:///localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to payment: %v", err)
	}
	defer payConn.Close()

	inventoryClient := inventoryV1.NewInventoryServiceClient(invConn)
	paymentClient := paymentV1.NewPaymentServiceClient(payConn)

	// ---------------- Order service ----------------

	orderService := NewOrderService(inventoryClient, paymentClient)

	// ---------------- OGEN server ----------------

// 	ogenServer, err := orderV1.NewServer(orderService)
// 	if err != nil {
// 		log.Fatalf("failed to create ogen server: %v", err)
// 	}
// 

	ogenServer, err := orderV1.NewServer(orderService)
	if err != nil {
		log.Fatalf("failed to create ogen server: %v", err)
	}

	httpServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: ogenServer, // <- здесь всё верно
	}

	// ---------------- Run HTTP server ----------------

	go func() {
		log.Printf("OrderService started on :%s", httpPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %v", err)
		}
	}()

	// ---------------- Graceful shutdown ----------------

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	log.Println("bye 👋")
}
