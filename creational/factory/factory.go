package factory

import "errors"

const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	return nil, errors.New("Not implemented yet")
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return ""
}

type DebitCardPM struct{}

func (d *DebitCardPM) Pay(amount float32) string {
	return ""
}
