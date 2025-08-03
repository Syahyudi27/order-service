package models

import "time"

type OrderHistory struct {
	ID        uint   `gorm:"primaryKey:autoIncrement"`
	OrderID   uint   `gorm:"type:bigint;not null"`
	Status    string `gorm:"type:varchar(30);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}