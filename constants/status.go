package constants

type OrderStatus int
type OrderStatusString string

const (
	Pending        OrderStatus = 100
	PendingPayment OrderStatus = 200
	Payment        OrderStatus = 300
	Expire         OrderStatus = 400

	PendingString       OrderStatusString = "pending"
	PendingPaymentString = "pending-payment"
	PaymentSuccessString = "payment-success"
	ExpireString        OrderStatusString = "expired"
)

var mapStatusStringToInt = map[OrderStatusString]OrderStatus{
	PendingString:        Pending,
	PendingPaymentString: PendingPayment,
	PaymentSuccessString: Payment,
	ExpireString:         Expire,
}

var mapStatusIntToString = map[OrderStatus]OrderStatusString{
	Pending:        PendingString,
	PendingPayment: PendingPaymentString,
	Payment:        PaymentSuccessString,
	Expire:         ExpireString,
}

func (ps OrderStatusString) String() string {
	return string(ps)
}

func (ps OrderStatus) Int() int {
	return int(ps)
}

func (ps OrderStatus) GetStatusString()OrderStatusString {
	return mapStatusIntToString[ps]
}

func (ps OrderStatusString) GetStatusInt() OrderStatus {
	return mapStatusStringToInt[ps]
}
