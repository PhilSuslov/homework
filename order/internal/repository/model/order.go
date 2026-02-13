package model

import "github.com/google/uuid"

type OrderDto struct {
    // Уникальный идентификатор заказа.
    OrderUUID uuid.UUID `json:"order_uuid"`
    // UUID пользователя.
    UserUUID  uuid.UUID   `json:"user_uuid"`
    PartUuids []uuid.UUID `json:"part_uuids"`
    // Итоговая стоимость.
    TotalPrice float64 `json:"total_price"`
    // UUID транзакции (если оплачен).
    TransactionUUID OptNilString     `json:"transaction_uuid"`
    PaymentMethod   OptPaymentMethod `json:"payment_method"`
    Status          OrderStatus      `json:"status"`
}

type PayOrderRequest struct {
    OrderUUID     uuid.UUID     `json:"order_uuid"`
    PaymentMethod PaymentMethod `json:"payment_method"`
}

type PayOrderParams struct {
    // UUID пользователя, для которого запрашиваются или
    // обновляются данные.
    OrderUUID uuid.UUID
}

type GetOrderByUUIDParams struct {
    // UUID пользователя, для которого запрашиваются или
    // обновляются данные.
    OrderUUID uuid.UUID
}

type OptNilString struct {
    Value uuid.UUID
    Set   bool
    Null  bool
}

type OptPaymentMethod struct {
    Value PaymentMethod
    Set   bool
}

type CancelOrderParams struct {
    // UUID пользователя, для которого запрашиваются или
    // обновляются данные.
    OrderUUID uuid.UUID
}

type PayOrderResponse struct {
    // Уникальный идентификатор транзакции.
    TransactionUUID uuid.UUID `json:"transaction_uuid"`
}

type PaymentMethod string

const (
	PaymentMethodUNKNOWN    PaymentMethod = "UNKNOWN"
	PaymentMethodCARD      PaymentMethod = "CARD"
	PaymentMethodSBP        PaymentMethod = "SBP"
	PaymentMethodCREDITCARD PaymentMethod = "CREDIT_CARD"
)

type OrderStatus string

const (
	OrderStatusPENDINGPAYMENT OrderStatus = "PENDING_PAYMENT"
	OrderStatusPAID           OrderStatus = "PAID"
	OrderStatusCANCELLED      OrderStatus = "CANCELLED"
)

type CreateOrderRes interface {
    createOrderRes()
}

type PayOrderRes interface {
    payOrderRes()
}

type GetOrderByUUIDRes interface {
    getOrderByUUIDRes()
}

type CancelOrderRes interface {
    cancelOrderRes()
}



func (PayOrderResponse) payOrderRes() {}
func (GetOrderByUUIDParams) getOrderByUUIDRes() {}
func (OrderDto) getOrderByUUIDRes() {}