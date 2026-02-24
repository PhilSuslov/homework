package part

import (
	"context"
	"math/rand/v2"
	"time"

	model "github.com/PhilSuslov/homework/inventory/internal/repository/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func CreateInitialInventoryOrder (ctx context.Context) (model.Note, error) {
		note := model.Note{
		OrderUUID: uuid.New().String(),
		Body: model.Part{
			Uuid:          uuid.New().String(),
			Name:          gofakeit.Name(),
			Price:         rand.Float64(),
			StockQuantity: int64(rand.IntN(20)),
			Category:      1,
			Manufacturer:  model.Manufacturer{Name: gofakeit.Name(), Country: gofakeit.Country()},
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	return note, nil


}