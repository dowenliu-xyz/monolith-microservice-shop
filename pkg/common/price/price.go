package price

import (
	"errors"

	pkg_errors "github.com/pkg/errors"
)

var (
	errPriceTooLow     = errors.New("price must be greater than 0")
	errInvalidCurrency = errors.New("invalid currency")
)

func NewErrPriceTooLow() error {
	return pkg_errors.WithStack(errPriceTooLow)
}

func IsErrPriceTooLow(err error) bool {
	return errors.Is(err, errPriceTooLow)
}

func NewErrInvalidCurrency() error {
	return pkg_errors.WithStack(errInvalidCurrency)
}

func IsErrInvalidCurrency(err error) bool {
	return errors.Is(err, errInvalidCurrency)
}

type Price struct {
	cents    uint
	currency string
}

func NewPrice(cents uint, currency string) (Price, error) {
	if cents <= 0 {
		return Price{}, NewErrPriceTooLow()
	}
	if len(currency) != 3 {
		return Price{}, NewErrInvalidCurrency()
	}

	return Price{cents, currency}, nil
}

// NewPriceP works as NewPrice, but on error it will panic instead of returning error.
func NewPriceP(cents uint, currency string) Price {
	p, err := NewPrice(cents, currency)
	if err != nil {
		panic(err)
	}

	return p
}

func (p Price) Cents() uint {
	return p.cents
}

func (p Price) Currency() string {
	return p.currency
}
