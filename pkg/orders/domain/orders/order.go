package orders

import (
	"errors"

	pkg_errors "github.com/pkg/errors"
)

type ID string

var errEmptyOrderID = errors.New("empty order id")

func IsErrEmptyOrderID(err error) bool {
	return errors.Is(err, errEmptyOrderID)
}

type Order struct {
	id      ID
	product Product
	address Address

	paid bool
}

func (o *Order) ID() ID {
	return o.id
}

func (o *Order) Product() Product {
	return o.product
}

func (o *Order) Address() Address {
	return o.address
}

func (o *Order) Paid() bool {
	return o.paid
}

func (o *Order) MarkAsPaid() {
	o.paid = true
}

func NewOrder(id ID, product Product, address Address) (*Order, error) {
	if len(id) == 0 {
		return nil, pkg_errors.WithStack(errEmptyOrderID)
	}

	return &Order{id, product, address, false}, nil
}
