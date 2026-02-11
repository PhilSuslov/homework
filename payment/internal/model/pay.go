package model

type PayOrderRequest struct {
	OrderUuid     string                 
	UserUuid      string          
	PaymentMethod PaymentMethod          
}



// Доступные способы оплаты
type PaymentMethod int32

const (
	PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED PaymentMethod = 0 // Неизвестный способ оплаты
	PaymentMethod_PAYMENT_METHOD_CARD                PaymentMethod = 1 // Банковская карта
	PaymentMethod_PAYMENT_METHOD_SPB                 PaymentMethod = 2 // СБП
	PaymentMethod_PAYMENT_METHOD_CREDIT_CARD         PaymentMethod = 3 // Кредитная карта
	PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY      PaymentMethod = 4 // Деньги инвестора
)

// Enum value maps for PaymentMethod.
var (
	PaymentMethod_name = map[int32]string{
		0: "PAYMENT_METHOD_UNKNOWN_UNSPECIFIED",
		1: "PAYMENT_METHOD_CARD",
		2: "PAYMENT_METHOD_SPB",
		3: "PAYMENT_METHOD_CREDIT_CARD",
		4: "PAYMENT_METHOD_INVESTOR_MONEY",
	}
	PaymentMethod_value = map[string]int32{
		"PAYMENT_METHOD_UNKNOWN_UNSPECIFIED": 0,
		"PAYMENT_METHOD_CARD":                1,
		"PAYMENT_METHOD_SPB":                 2,
		"PAYMENT_METHOD_CREDIT_CARD":         3,
		"PAYMENT_METHOD_INVESTOR_MONEY":      4,
	}
)