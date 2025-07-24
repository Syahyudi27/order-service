package constants

type Paymentstatus int
type PaymentstatusString string

const (
	Intial     Paymentstatus = 0
	Pending    Paymentstatus = 100
	Settlement Paymentstatus = 200
	Expire     Paymentstatus = 300

	InitialString    PaymentstatusString = "initial"
	PendingString    PaymentstatusString = "pending"
	SettlementString PaymentstatusString = "settlement"
	ExpireString     PaymentstatusString = "fail"
)

var mapStatusStringToInt = map[PaymentstatusString]Paymentstatus{
	InitialString:    Intial,
	PendingString:    Pending,
	SettlementString: Settlement,
	ExpireString:     Expire,
}

var mapStatusIntToString = map[Paymentstatus]PaymentstatusString{
	Intial:     InitialString,
	Pending:    PendingString,
	Settlement: SettlementString,
	Expire:     ExpireString,
}

func (ps PaymentstatusString) String() string {
	return string(ps)
}

func (ps Paymentstatus) Int() int {
	return int(ps)
}

func (ps Paymentstatus) GetStatusString() PaymentstatusString {
	return mapStatusIntToString[ps]
}

func (ps PaymentstatusString) GetStatusInt() Paymentstatus {
	return mapStatusStringToInt[ps]
}
