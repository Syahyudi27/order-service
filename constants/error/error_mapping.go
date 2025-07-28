package error

import (
	errOrder "order-service/constants/error/Order"
)

func ErrMapping(err error) bool {
	var (
		GeneralError = GeneralErrors
		OrderErrors   = errOrder.OrderErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralError...)
	allErrors = append(allErrors, OrderErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
