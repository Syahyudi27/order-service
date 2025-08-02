package dto

import (
	"time"

	"github.com/google/uuid"
)

type PaymentRequest struct {
	OrderID        string          `json:"orderId"`
	ExpiredAt      time.Time       `json:"expiredAt"`
	Amount         float64         `json:"amount"`
	Description    *string         `json:"description"`
	CustomerDetail *CustomerDetail `json:"customerDetail"`
	ItemDetails    []ItemDetail    `json:"itemDetails"`
}

type CustomerDetail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ItemDetail struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Amount   int       `json:"price"`
	Quantity int       `json:"qty"`
}
