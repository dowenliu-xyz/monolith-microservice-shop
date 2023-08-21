package application

import (
	"github.com/pkg/errors"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/price"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/shop/domain/products"
)

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService(repo products.Repository, readModel productReadModel) ProductsService {
	return ProductsService{repo, readModel}
}

func (s ProductsService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()
}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	productPrice, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.WithMessage(err, "invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, productPrice)
	if err != nil {
		return errors.WithMessage(err, "cannot create product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.WithMessage(err, "cannot save product")
	}

	return nil
}
