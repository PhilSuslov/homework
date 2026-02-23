package app

import (
	"context"
	"fmt"

	api "github.com/PhilSuslov/homework/inventory/internal/api/inventory/v1"
	"github.com/PhilSuslov/homework/inventory/internal/config"
	"github.com/PhilSuslov/homework/inventory/internal/repository"
	partRepo "github.com/PhilSuslov/homework/inventory/internal/repository/part"
	"github.com/PhilSuslov/homework/inventory/internal/service"
	partService "github.com/PhilSuslov/homework/inventory/internal/service/part"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type diContainer struct {
	inventoryV1API      inventoryV1.InventoryServiceServer
	inventoryService    service.InventoryService
	inventoryRepository repository.InventoryRepository

	mongoDBClient *mongo.Client
	mongoDBHandle *mongo.Database
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) InventoryV1API(ctx context.Context) inventoryV1.InventoryServiceServer {
	if d.inventoryV1API == nil {
		d.inventoryV1API = api.NewAPI(d.PartService(ctx))
	}
	return d.inventoryV1API
}

func (d *diContainer) PartService(ctx context.Context) service.InventoryService {
	if d.inventoryService == nil {
		d.inventoryService = partService.NewService(d.PartRepository(ctx))
	}
	return d.inventoryService
}

func (d *diContainer) PartRepository(ctx context.Context) repository.InventoryRepository {
	if d.inventoryRepository == nil {
		d.inventoryRepository = partRepo.NewNoteRepository(d.MongoDBHandle(ctx))
	}
	return d.inventoryRepository
}

func (d *diContainer) MongoDBClient(ctx context.Context) *mongo.Client {
	if d.mongoDBClient == nil {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to MongoDB: %s\n", err.Error()))
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(fmt.Sprintf("failed to ping MongoDB: %v\n", err))
		}

		closer.AddNamed("MongoDB client", func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})
		d.mongoDBClient = client
	}
	return d.mongoDBClient
}

func (d *diContainer) MongoDBHandle(ctx context.Context) *mongo.Database {
	if d.mongoDBHandle == nil {
		d.mongoDBHandle = d.MongoDBClient(ctx).Database(config.AppConfig().Mongo.DatabaseName())
	}
	return d.mongoDBHandle
}
