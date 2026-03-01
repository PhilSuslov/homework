package kafka

import "github.com/PhilSuslov/homework/order/internal/model"

type OrderRecordedDecoder interface {
	Decode(data []byte) (model.OrderRecordedEvent, error)
}