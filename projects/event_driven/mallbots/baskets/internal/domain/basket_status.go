package domain

type BasketStatus string

const (
	BasketUnknown      BasketStatus = ""
	BasketIsOpen       BasketStatus = "open"
	BasketIsCanceled   BasketStatus = "canceled"
	BasketIsCheckedOut BasketStatus = "checked_out"
)

func (s BasketStatus) String() string {
	switch s {
	case BasketIsOpen, BasketIsCanceled, BasketIsCheckedOut, BasketUnknown:
		return string(s)
	default:
		return ""
	}
}

func ToBasketStatus(status string) BasketStatus {
	bstatus := BasketStatus(status)
	switch bstatus {
	case BasketIsOpen, BasketIsCanceled, BasketIsCheckedOut, BasketUnknown:
		return bstatus
	default:
		return BasketUnknown
	}
}
