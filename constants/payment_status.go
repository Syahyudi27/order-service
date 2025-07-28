package constants

type PaymentStatusSting string

const (
	PendingPaymentStatus    PaymentStatusSting = "pending"
	SettlementPaymentStatus PaymentStatusSting = "settlement"
	ExpiredPaymentStatus    PaymentStatusSting = "expire"
)
