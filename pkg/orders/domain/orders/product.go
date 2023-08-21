package orders

import (
	"errors"

	pkg_errors "github.com/pkg/errors"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/price"
)

type ProductID string

var errEmptyProductID = errors.New("empty product ID")

func NewErrEmptyProductID() error {
	return pkg_errors.WithStack(errEmptyProductID)
}

func IsErrEmptyProductID(err error) bool {
	return errors.Is(err, errEmptyProductID)
}

type Product struct {
	id    ProductID
	name  string
	price price.Price
}

func NewProduct(id ProductID, name string, price price.Price) (Product, error) {
	if len(id) == 0 {
		return Product{}, NewErrEmptyProductID()
	}

	return Product{id, name, price}, nil
}

func (p Product) ID() ProductID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}
