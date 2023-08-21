package products

import (
	"errors"

	pkg_errors "github.com/pkg/errors"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/price"
)

type ID string

var (
	errEmptyID   = errors.New("empty product ID")
	errEmptyName = errors.New("empty product name")
)

func NewErrEmptyID() error {
	return pkg_errors.WithStack(errEmptyID)
}

func IsErrEmptyID(err error) bool {
	return errors.Is(err, errEmptyID)
}

func NewErrEmptyName() error {
	return pkg_errors.WithStack(errEmptyName)
}

func IsErrEmptyName(err error) bool {
	return errors.Is(err, errEmptyName)
}

type Product struct {
	id ID

	name        string
	description string

	price price.Price
}

func NewProduct(id ID, name string, description string, price price.Price) (*Product, error) {
	if len(id) == 0 {
		return nil, NewErrEmptyID()
	}
	if len(name) == 0 {
		return nil, NewErrEmptyName()
	}

	return &Product{id, name, description, price}, nil
}

func (p Product) ID() ID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Price() price.Price {
	return p.price
}
