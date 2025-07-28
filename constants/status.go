package constants

type Orderstatus int
type OrderstatusString string

const (
	Pending        Orderstatus = 100
	PendingPayment Orderstatus = 200
	Payment        Orderstatus = 300
	Expire         Orderstatus = 400

	PendingString        OrderstatusString = "pending"
	PendingPaymentString OrderstatusString = "pending-payment"
	PaymentSuccessString OrderstatusString = "payment-success"
	ExpireString         OrderstatusString = "expired"
)

var mapStatusStringToInt = map[OrderstatusString]Orderstatus{
	PendingString:        Pending,
	PendingPaymentString: PendingPayment,
	PaymentSuccessString: Payment,
	ExpireString:         Expire,
}

var mapStatusIntToString = map[Orderstatus]OrderstatusString{
	Pending:        PendingString,
	PendingPayment: PendingPaymentString,
	Payment:        PaymentSuccessString,
	Expire:         ExpireString,
}

func (ps OrderstatusString) String() string {
	return string(ps)
}

func (ps Orderstatus) Int() int {
	return int(ps)
}

func (ps Orderstatus) GetStatusString() OrderstatusString {
	return mapStatusIntToString[ps]
}

func (ps OrderstatusString) GetStatusInt() Orderstatus {
	return mapStatusStringToInt[ps]
}
