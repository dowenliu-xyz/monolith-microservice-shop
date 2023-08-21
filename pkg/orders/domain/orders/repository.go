package orders

import (
	"errors"

	pkg_errors "github.com/pkg/errors"
)

var errNotFound = errors.New("order not found")

func NewErrNotFound() error {
	return pkg_errors.WithStack(errNotFound)
}

func IsErrNotFound(err error) bool {
	return errors.Is(err, errNotFound)
}

type Repository interface {
	Save(*Order) error
	ByID(ID) (*Order, error)
}
