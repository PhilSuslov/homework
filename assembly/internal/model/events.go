package model

type ShipAssembled struct {
	Event_uuid     string // Уникальный идентификатор события (для идемпотентности)
	Order_uuid     string // Идентификатор собранного заказа
	User_uuid      string // Идентификатор пользователя
	Build_time_sec int64 // Время (в секундах), потраченное на сборку корабля
}

