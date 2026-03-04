package model

type OrderPaidEvent struct {
	Event_uuid       string
	Order_uuid       string
	User_uuid        string
	Payment_method   string
	Transaction_uuid string
}

type ShipAssembled struct {
	Event_uuid     string
	Order_uuid     string
	User_uuid      string
	Build_time_sec int64
}
