package app

import (
	"context"
	"fmt"
	"log"
	"sync"

	api "github.com/PhilSuslov/homework/inventory/internal/api/inventory/v1"
	"github.com/PhilSuslov/homework/inventory/internal/config"
	"github.com/PhilSuslov/homework/inventory/internal/repository"
	partRepo "github.com/PhilSuslov/homework/inventory/internal/repository/part"
	"github.com/PhilSuslov/homework/inventory/internal/service"
	partService "github.com/PhilSuslov/homework/inventory/internal/service/part"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"go.mongodb.org/mongo-driver/bson"
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

	mu   sync.Mutex
	once sync.Once
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

// ----------------- MongoDB -----------------

func (d *diContainer) MongoDBClient(ctx context.Context) (*mongo.Client, error) {
	var err error
	d.once.Do(func() {
		client, e := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
		if e != nil {
			err = fmt.Errorf("failed to connect to MongoDB: %w", e)
			return
		}

		if e := client.Ping(ctx, readpref.Primary()); e != nil {
			err = fmt.Errorf("failed to ping MongoDB: %w", e)
			return
		}

		d.mongoDBClient = client
		closer.AddNamed("MongoDB client", func(ctx context.Context) error {
			return d.mongoDBClient.Disconnect(ctx)
		})
	})
	return d.mongoDBClient, err
}

func (d *diContainer) MongoDBHandle(ctx context.Context) (*mongo.Database, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.mongoDBHandle != nil {
		return d.mongoDBHandle, nil
	}

	client, err := d.MongoDBClient(ctx)
	if err != nil {
		return nil, err
	}

	d.mongoDBHandle = client.Database(config.AppConfig().Mongo.DatabaseName())

	// Создаем индексы
	collection := d.mongoDBHandle.Collection("note")
	fmt.Printf("Inserting fake note with UUID: %s\n", collection.Name())
		indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "body.name", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
	}
	if _, err := collection.Indexes().CreateMany(ctx, indexModels); err != nil {
		log.Panicf("failed to create indexes: %v", err)
	}

	return d.mongoDBHandle, nil
}

// ----------------- Repository -----------------

func (d *diContainer) PartRepository(ctx context.Context) (repository.InventoryRepository, error) {
	if d.inventoryRepository != nil {
		return d.inventoryRepository, nil
	}

	db, err := d.MongoDBHandle(ctx)
	if err != nil {
		return nil, err
	}

	d.inventoryRepository = partRepo.NewNoteRepository(db)
	return d.inventoryRepository, nil
}

// ----------------- Service -----------------

func (d *diContainer) PartService(ctx context.Context) (service.InventoryService, error) {
	if d.inventoryService != nil {
		return d.inventoryService, nil
	}

	repo, err := d.PartRepository(ctx)
	if err != nil {
		return nil, err
	}

	d.inventoryService = partService.NewService(repo)
	return d.inventoryService, nil
}

// ----------------- API -----------------

func (d *diContainer) InventoryV1API(ctx context.Context) (inventoryV1.InventoryServiceServer, error) {
	if d.inventoryV1API != nil {
		return d.inventoryV1API, nil
	}

	svc, err := d.PartService(ctx)
	if err != nil {
		return nil, err
	}

	d.inventoryV1API = api.NewAPI(svc)
	return d.inventoryV1API, nil
}

// ----------------- Инициализация фейковых данных -----------------

func (d *diContainer) InitFakeData(ctx context.Context) error {
	db, err := d.MongoDBHandle(ctx)
	if err != nil {
		return err
	}

	collection := db.Collection("note")

	// Создаём начальный объект
	note, err := partRepo.CreateInitialInventoryOrder(ctx) // <-- правильно используем пакет partRepo
	if err != nil {
		log.Printf("⚠️ Failed to create initial inventory order: %v", err)
		return err
	}

	// Проверяем, есть ли уже запись с таким OrderUUID
	count, err := collection.CountDocuments(ctx, bson.M{"orderuuid": note.OrderUUID})
	if err != nil {
		return fmt.Errorf("failed to check existing initial note: %w", err)
	}

	if count == 0 {
		if _, err := collection.InsertOne(ctx, note); err != nil {
			return fmt.Errorf("failed to insert initial note: %w", err)
		}
		log.Printf("✅ Initial note inserted: %v", note.OrderUUID)
	} else {
		log.Printf("⚠️ Initial note already exists: %v", note.OrderUUID)
	}

	return nil
}