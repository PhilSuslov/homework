package kafka

import "github.com/PhilSuslov/homework/notification/internal/model"

type OrderAssembledDecoder interface {
	Decode(data []byte) (model.ShipAssembled, error)
}

type OrderPaidDecoder interface {
	Decode(data []byte) (model.OrderPaidEvent, error)
}