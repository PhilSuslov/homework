package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/order"
	orderServ "github.com/PhilSuslov/homework/order/internal/service/order"
	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort = "8080"
	// readHeaderTimeout = 5 * time.Second
	shutdownTimeout = 10 * time.Second
)

//
// func (s *OrderService) CreateOrder(ctx context.Context, request *orderV1.CreateOrderRequest) (orderV1.CreateOrderResponse, error) {
// 	log.Printf("CreateOrder called, req: %+v", request)
//
// 	if request.UserUUID == uuid.Nil || len(request.UserUUID) == 0 {
// 		return orderV1.CreateOrderResponse{}, status.Error(codes.InvalidArgument, "user_uuid and part_uuid are required")
// 	}
//
// 	strUUID := make([]string, len(request.PartUuids))
// 	for i, v := range request.PartUuids {
// 		strUUID[i] = v.String()
// 	}
//
// 	//1. Запрашиваем детали из Inventory
// 	partsResp, err := s.inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{
// 		Filter: &inventoryV1.PartsFilter{
// 			Uuids: strUUID,
// 		},
// 	})
// 	if err != nil {
// 		log.Printf("Inventory ListParts error: %v", err)
// 		return &orderV1.CreateOrderResponse{}, status.Errorf(codes.Internal, "inventory error: %v", err)
// 	}
//
// 	log.Printf("Inventory ListParts response: %+v", partsResp)
//
// 	// 2. Проверяем, что все детали найдены
// 	if len(partsResp.Parts) != len(request.PartUuids) {
// 		log.Printf("Not all parts found: expected=%d, got=%d", len(request.PartUuids), len(partsResp.Parts))
// 		return &orderV1.CreateOrderResponse{}, status.Error(codes.NotFound, "some parts not found")
// 	}
//
// 	// 3. Считаем цену
// 	var total_price float64
// 	for _, p := range partsResp.Parts {
// 		log.Printf("Part: UUID=%s, Price=%f", p.Uuid, p.Price)
// 		total_price += p.Price
// 	}
// 	log.Printf("Total price calculated: %f", total_price)
//
// 	orderUUID := uuid.New()
//
// 	order := &orderV1.OrderDto{
// 		OrderUUID:  orderUUID,
// 		UserUUID:   request.UserUUID,
// 		PartUuids:  request.PartUuids,
// 		TotalPrice: total_price,
// 		Status:     orderV1.OrderStatusPENDINGPAYMENT,
// 	}
//
// 	s.mu.Lock()
// 	s.orders[orderUUID.String()] = order
// 	log.Printf("Order saved: UUID=%s, map keys: %v", order.OrderUUID.String(), s.orders)
// 	defer s.mu.Unlock()
//
// 	return &orderV1.CreateOrderResponse{
// 		OrderUUID:  order.OrderUUID,
// 		TotalPrice: order.TotalPrice,
// 	}, nil
//
// }

// --------------------------------------------------
// -- PAY ORDER
// --------------------------------------------------
//
// func (s *OrderService) PayOrder(ctx context.Context,
// 	request *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
// 	log.Println(params)
//
// 	s.mu.Lock()
// 	order, ok := s.orders[params.OrderUUID.String()]
// 	if !ok {
// 		s.mu.Unlock()
// 		return nil, status.Error(codes.NotFound, "order not found")
// 	}
//
// 	if order.Status == orderV1.OrderStatusPAID {
// 		s.mu.Unlock()
// 		return nil, status.Error(codes.Canceled, "order already paid")
// 	}
//
// 	if order.Status == orderV1.OrderStatusCANCELLED {
// 		s.mu.Unlock()
// 		return nil, status.Error(codes.Canceled, "order cancelled")
// 	}
//
// 	// log.Println("+++++++++++++++++++++++++++++")
// 	s.mu.Unlock()
//
// 	//Проверка метода оплаты
// 	var pm paymentV1.PaymentMethod
// 	switch request.PaymentMethod {
// 	case "CARD":
// 		pm = 1
// 	case "SBP":
// 		pm = 2
// 	case "CREDIT_CARD":
// 		pm = 3
// 	case "INVESTOR_MONEY":
// 		pm = 4
// 	default:
// 		pm = 0
// 	}
//
// 	// Вызываем PaymentService
// 	payResp, err := s.paymentClient.PayOrder(ctx, &paymentV1.PayOrderRequest{
// 		OrderUuid:     request.OrderUUID.String(),
// 		UserUuid:      order.UserUUID.String(),
// 		PaymentMethod: pm,
// 	})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "payment error: %v", err)
// 	}
// 	transactionUUID := payResp.TransactionUuid
// 	paymentMethod := request.PaymentMethod
//
// 	transactionuuid, _ := uuid.Parse(transactionUUID)
//
// 	s.mu.Lock()
// 	order.Status = orderV1.OrderStatusPAID
// 	order.TransactionUUID.Value = transactionuuid
// 	order.PaymentMethod.Value = paymentMethod
// 	log.Println("=== Статус оплаты должен быть ===", s.orders)
// 	s.mu.Unlock()
//
// 	resp := &orderV1.PayOrderResponse{TransactionUUID: transactionuuid}
//
// 	return resp, nil
//
// }

// --------------------------------------------------
// -- GET ORDER
// --------------------------------------------------

// func (s *OrderService) GetOrderByUUID(ctx context.Context, params orderV1.GetOrderByUUIDParams) (orderV1.GetOrderByUUIDRes, error) {
// 	s.mu.Lock()
// 	order, ok := s.orders[params.OrderUUID.String()]
// 	// log.Printf("GetOrderByUUID called: looking for %s, found: %v", params.OrderUUID.String(), ok)
// 	log.Printf("GetOrderByUUID: param=%s, map keys=%v", params.OrderUUID.String(), s.orders)
// 	s.mu.Unlock()
//
// 	if !ok {
// 		log.Printf("s.orders[params.OrderUUID.String()] - %v. Order not found in map!", s.orders[params.OrderUUID.String()])
// 		return nil, status.Error(codes.NotFound, "order not found")
// 	}
//
// 	return &orderV1.OrderDto{
// 		OrderUUID:       order.OrderUUID,
// 		UserUUID:        order.UserUUID,
// 		PartUuids:       order.PartUuids,
// 		TotalPrice:      order.TotalPrice,
// 		TransactionUUID: order.TransactionUUID,
// 		PaymentMethod:   order.PaymentMethod,
// 		Status:          order.Status,
// 	}, nil
// }

// --------------------------------------------------
// -- CANCEL ORDER
// --------------------------------------------------
//
// func (s *OrderService) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
// 	s.mu.Lock()
// 	order, ok := s.orders[params.OrderUUID.String()]
// 	if !ok {
// 		s.mu.Unlock()
// 		return nil, status.Error(codes.NotFound, "order not found")
// 	}
//
// 	if order.Status == orderV1.OrderStatusPAID {
// 		s.mu.Unlock()
// 		return nil, status.Error(codes.Unknown, "Conflict")
// 	}
//
// 	order.Status = orderV1.OrderStatusCANCELLED
// 	s.mu.Unlock()
// 	return nil, status.Error(codes.Canceled, "No content")
// }

// func (s *OrderService) NewError(ctx context.Context, err error) *orderV1.GenericErrorStatusCode {
// 	// Тут можно возвращать любую реализацию ErrorRes, например:
// 	var Err orderV1.GenericErrorStatusCode
// 	Err.StatusCode = 500
// 	Err.Response.Code.Value = 500
// 	Err.Response.Message.Value = err.Error()
// 	return &orderV1.GenericErrorStatusCode{
// 		StatusCode: Err.StatusCode,
// 		Response:   Err.Response,
// 	}
// }
//
// type loggingResponseWriter struct {
// 	http.ResponseWriter
// 	statusCode int
// }
//
// func (lrw *loggingResponseWriter) WriteHeader(code int) {
// 	lrw.statusCode = code
// 	lrw.ResponseWriter.WriteHeader(code)
// }

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
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

	repo := orderRepo.NewOrderRepo()
	orderService := orderServ.NewOrderService(inventoryClient, paymentClient, repo)

	// ---------------- OGEN server ----------------

	ogenServer, err := orderV1.NewServer(orderService)
	if err != nil {
		log.Fatalf("failed to create ogen server: %v", err)
	}

	// 	httpServer := &http.Server{
	//     Addr:    ":" + httpPort,
	//     Handler: ogenServer,
	// }

	httpServer := &http.Server{
		Addr: ":" + httpPort,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("REQUEST: %s %s — вход в OGEN", r.Method, r.URL.Path)

			// Создаем обертку ResponseWriter, чтобы логировать статус-код
			lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
			ogenServer.ServeHTTP(lrw, r)

			body, _ := io.ReadAll(r.Body)
			log.Printf("Body: %s, Query: %s", string(body), r.URL.RawQuery)
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			log.Printf("REQUEST: %s %s — обработан OGEN с кодом %d", r.Method, r.URL.Path, lrw.statusCode)
		}),
	}

	log.Println("===== CHECK ROUTES =====")
	paths := []string{
		"/api/v1/order/",
		"/api/v1/order/123e4567-e89b-12d3-a456-426614174000",
	}
	methods := []string{"GET", "POST"}

	for _, p := range paths {
		for _, m := range methods {
			_, ok := ogenServer.FindRoute(m, p)
			log.Println(m, p, "->", ok)
		}
	}

	// 	ops := []string{
	//     "/api/v1/order/",
	//     "/api/v1/order",
	//     "/api/v1/order",
	//     "/api/v1/order/",
	//     "/api/v1/order/123e4567-e89b-12d3-a456-426614174000", // тестовый UUID
	// }
	//
	// 	mtd := []string{"GET", "POST", "PUT", "DELETE"}
	//
	// 	for _, op := range ops {
	// 		for _, mt := range mtd {
	// 			if route, ok := ogenServer.FindRoute(mt, op); ok {
	// 				log.Printf("FOUND: %s %s -> %s\n", mt, op, route.OperationID)
	// 			} else {
	// 				log.Printf("NOT FOUND: %s %s\n", mt, op)
	// 			}
	// 		}
	// }

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
