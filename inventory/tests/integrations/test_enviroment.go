package integrations

import (
	"context"
	"os"
	"time"

	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (env *TestEnvironment) InsertTestNote(ctx context.Context) (string, error) {
	noteUUID := gofakeit.UUID()
	now := time.Now()

	noteDoc := bson.M{
		"_id": noteUUID,
		"body": bson.M{
			"Uuid":          gofakeit.UUID(),
			"Name":          gofakeit.Name(),
			"Description":   gofakeit.Sentence(),
			"Price":         gofakeit.Float64(),
			"StockQuantity": gofakeit.Int64(),
			"Category":      1,
			"CreatedAt":     primitive.NewDateTimeFromTime(now),
			"UpdatedAt":     primitive.NewDateTimeFromTime(now),
		},
	}
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory-service"
	}

	_, err := env.Mongo.Client().Database(databaseName).
		Collection(notiesCollectionName).InsertOne(ctx, noteDoc)
	if err != nil {
		return "", err
	}
	return noteUUID, nil
}

func (env *TestEnvironment) GetListPartsNoteInfo() *inventory_v1.ListPartsRequest {
	var filters *inventory_v1.PartsFilter
	filters = &inventory_v1.PartsFilter{
		Uuids:      []string{gofakeit.UUID(), gofakeit.UUID()},
		Names:      []string{gofakeit.Name(), gofakeit.Name()},
		Categories: []inventory_v1.Category{1, 2, 3},
	}
	return &inventory_v1.ListPartsRequest{
		Filter: filters,
	}
}

func (env *TestEnvironment) ClearNoteCollection(ctx context.Context) error {
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory-service"
	}

	_, err := env.Mongo.Client().Database(databaseName).
		Collection(notiesCollectionName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}
	return nil
}
