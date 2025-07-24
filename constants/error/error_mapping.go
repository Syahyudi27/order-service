package error

import (
	errPayment "order-service/constants/error/payment"
)

func ErrMapping(err error) bool {
	var (
		GeneralError = GeneralErrors
		PaymentErrors   = errPayment.PaymentErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralError...)
	allErrors = append(allErrors, PaymentErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
