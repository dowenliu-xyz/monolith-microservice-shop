package price_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/price"
)

func TestNewPrice(t *testing.T) {
	testCases := []struct {
		Name        string
		Cents       uint
		Currency    string
		ExpectedErr assert.ErrorAssertionFunc
	}{
		{
			Name:        "valid",
			Cents:       10,
			Currency:    "EUR",
			ExpectedErr: assert.NoError,
		},
		{
			Name:     "invalid_cents",
			Cents:    0,
			Currency: "EUR",
			ExpectedErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.True(t, price.IsErrPriceTooLow(err), i...)
			},
		},
		{
			Name:     "empty_currency",
			Cents:    10,
			Currency: "",
			ExpectedErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.True(t, price.IsErrInvalidCurrency(err), i...)
			},
		},
		{
			Name:     "invalid_currency_length",
			Cents:    10,
			Currency: "US",
			ExpectedErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.True(t, price.IsErrInvalidCurrency(err), i...)
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			_, err := price.NewPrice(c.Cents, c.Currency)
			c.ExpectedErr(t, err)
		})
	}
}
