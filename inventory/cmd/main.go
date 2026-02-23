package main

import (
	"context"
	// "log"
	// "net"
	"os/signal"
	"syscall"
	"fmt"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	// inventoryAPI "github.com/PhilSuslov/homework/inventory/internal/api/inventory/v1"
	"github.com/PhilSuslov/homework/inventory/internal/config"
	// "github.com/PhilSuslov/homework/inventory/internal/model"
	"github.com/PhilSuslov/homework/inventory/internal/app"
	// inventoryRepo "github.com/PhilSuslov/homework/inventory/internal/repository/part"
	// inventoryService "github.com/PhilSuslov/homework/inventory/internal/service/part"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	// inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

const configPath = "../../deploy/compose/inventory/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil{
		logger.Error(appCtx, "❌ Не удалось создать приложение",zap.Error(err))
		return
	}

	err = a.Run(appCtx)
	if err != nil{
		logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
		return
	}

// 	
// 
// 	clientOpts := options.Client().ApplyURI(config.AppConfig().Mongo.URI())
// 	client, err := mongo.Connect(ctx, clientOpts)
// 	if err != nil {
// 		logger.Error(ctx, "Ошибка подключения к MongoDB", zap.Error(err))
// 		panic(err)
// 	}
// 
// 	if err := client.Ping(ctx, nil); err != nil {
// 		logger.Error(ctx, "❌ Не удалось пингануть MongoDB", zap.Error(err))
// 		panic(err)
// 	}
// 
// 	db := client.Database(config.AppConfig().Mongo.DatabaseName())
// 	collection := db.Collection("notes") // По идее note тоже передать через config
// 
// 	indexModels := []mongo.IndexModel{
// 		{
// 			Keys:    bson.D{{Key: "body.name", Value: 1}},
// 			Options: options.Index().SetUnique(false),
// 		},
// 	}
// 
// 	_, err = collection.Indexes().CreateMany(ctx, indexModels)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 
// 	lis, err := net.Listen("tcp", config.AppConfig().InventoryGRPC.Address())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 
// 	defer func() {
// 		if cerr := lis.Close(); cerr != nil {
// 			log.Printf("Failed to closer listener: %v\n", cerr)
// 		}
// 	}()
// 
// 	note := model.Note{
// 		OrderUUID: "8f1c1f5a-2d5f-4c2f-8e13-123456789abc",
// 		Body: model.Part{
// 			Uuid:          "2f1c3f5a-2d5f-4c5f-8e13-123456749abc",
// 			Name:          "Test Part",
// 			Price:         100,
// 			StockQuantity: 10,
// 			Category:      1,
// 			Manufacturer:  model.Manufacturer{Name: "Acme", Country: "USA"},
// 			CreatedAt:     time.Now(),
// 			UpdatedAt:     time.Now(),
// 		},
// 	}
// 	_, err = collection.InsertOne(context.Background(), note)
// 	if err != nil {
// 		log.Printf("Note with OrderUUID: %v already insert\n", note.OrderUUID)
// 	}
// 
// 	//------------------------------------------------------------------
// 
// 	grpcServer := grpc.NewServer()
// 
// 	// Регистрируем наш сервис
// 	repo := inventoryRepo.NewNoteRepository(db)
// 	service := inventoryService.NewService(repo)
// 	api := inventoryAPI.NewAPI(service)
// 
// 	inventory_v1.RegisterInventoryServiceServer(grpcServer, api)
// 	reflection.Register(grpcServer)
// 
// 	log.Printf("📦 Inventory service started on %s\v", config.AppConfig().InventoryGRPC.Address())
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatal(err)
// 	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second )
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil{
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}