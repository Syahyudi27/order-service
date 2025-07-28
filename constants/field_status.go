package constants

type FieldStatusSting string

const (
	AvailableStatus FieldStatusSting = "pending"
	BookedStatus FieldStatusSting = "settlement"
)

func (p FieldStatusSting) String()string{
	return string(p)
}
