package model

type OrderRecordedEvent struct{
	Event_uuid string
	Order_uuid string
	User_uuid string
	Payment_method string
	Transaction_uuid string
}