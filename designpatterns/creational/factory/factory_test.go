package creational

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(CASH)
	if err != nil {
		t.Error(err)
	}
	msg := payment.Pay(10.2)
	if !strings.Contains(msg, "cash") {
		t.Error("Payment for CASH aint cash")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebit(t *testing.T) {
	payment, err := GetPaymentMethod(DEBIT)
	if err != nil {
		t.Error(err)
	}
	msg := payment.Pay(11.2)
	if !strings.Contains(msg, "debit") {
		t.Error("Payment for DEBIT aint debit")
	}
	t.Log("LOG:", msg)
}
