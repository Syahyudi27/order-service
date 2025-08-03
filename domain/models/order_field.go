package models

import (
	"order-service/constants"
	"time"
)

type OrderField struct {
	ID              uint                        `gorm:"primaryKey:autoIncrement"`
	OrderID         uint                        `gorm:"type:bigint;not null"`
	FieldScheduleID constants.OrderStatusString `gorm:"type:varchar(30);not null"`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
