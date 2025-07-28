package error

import "errors"

var (
	ErrOrderNotFound = errors.New("Order not found")
	ErrFieldAlreadyBooked = errors.New("field schedule already booked")
)

var OrderErrors = []error{
	ErrOrderNotFound,
	ErrFieldAlreadyBooked,
}
