package kafka

import "github.com/PhilSuslov/homework/assembly/internal/model"

type AssemblyDecoder interface{
	Decode (data []byte) (model.ShipAssembled, error)
}