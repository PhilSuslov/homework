package service
import (
	"context"

	"github.com/PhilSuslov/homework/notification/internal/model"


)

type TelegramPaidService interface {
	SendPaidNotification(ctx context.Context, uuid string, paid model.OrderPaidEvent) error

}

type TelegramAssembledService interface {
	SendAssembledNotification(ctx context.Context, uuid string, assembly model.ShipAssembled) error

}

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

