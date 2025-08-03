package dto

import "order-service/constants"

type OrderFieldRequest struct {
	OrderID         uint
	FieldScheduleID constants.OrderStatusString
}