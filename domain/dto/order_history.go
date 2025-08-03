package dto

import "order-service/constants"

type OrderHistoryRequest struct {
	OrderID string
	Status  constants.OrderStatusString
}