package service

import (
	"context"
	"github.com/PhilSuslov/homework/assembly/internal/model"
)

type AssemblyProducerService interface {
	ProducerAssembledRecorded(ctx context.Context, event model.ShipAssembled) error
}

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}