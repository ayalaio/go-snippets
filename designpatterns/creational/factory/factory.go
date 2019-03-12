package creational

import "errors"

const (
	CASH  = 1
	DEBIT = 2
)

type PaymentMethod interface {
	Pay(float32) string
}

func GetPaymentMethod(paymentType int) (pm PaymentMethod, err error) {
	switch paymentType {
	case CASH:
		pm, err = &CashPaymentMethod{}, nil
	case DEBIT:
		pm, err = &DebitPaymentMethod{}, nil
	default:
		pm, err = nil, errors.New("Not implemented")
	}
	return
}

type CashPaymentMethod struct{}

type DebitPaymentMethod struct{}

func (_ *CashPaymentMethod) Pay(_ float32) string {
	return "Payed with cash"
}

func (_ *DebitPaymentMethod) Pay(_ float32) string {
	return "Payed with debit"
}
