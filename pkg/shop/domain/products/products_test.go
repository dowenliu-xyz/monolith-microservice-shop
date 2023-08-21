package products_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/price"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/shop/domain/products"
)

func TestNewProduct(t *testing.T) {
	testPrice, err := price.NewPrice(42, "USD")
	assert.NoError(t, err)

	testCases := []struct {
		TestName string

		ID          products.ID
		Name        string
		Description string
		Price       price.Price

		ExpectedErr assert.ErrorAssertionFunc
	}{
		{
			TestName:    "valid",
			ID:          "1",
			Name:        "foo",
			Description: "bar",
			Price:       testPrice,
			ExpectedErr: assert.NoError,
		},
		{
			TestName:    "empty_id",
			ID:          "",
			Name:        "foo",
			Description: "bar",
			Price:       testPrice,

			ExpectedErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.True(t, products.IsErrEmptyID(err), i...)
			},
		},
		{
			TestName:    "empty_name",
			ID:          "1",
			Name:        "",
			Description: "bar",
			Price:       testPrice,

			ExpectedErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.True(t, products.IsErrEmptyName(err), i...)
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := products.NewProduct(c.ID, c.Name, c.Description, c.Price)
			c.ExpectedErr(t, err)
		})
	}
}
