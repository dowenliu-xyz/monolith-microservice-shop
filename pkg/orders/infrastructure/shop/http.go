package shop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/price"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/domain/orders"
	shop_http "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/shop/interfaces/private/http"
)

type HTTPClient struct {
	address string
}

func NewHTTPClient(address string) HTTPClient {
	return HTTPClient{address}
}

func (h HTTPClient) ProductByID(id orders.ProductID) (orders.Product, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", h.address, id))
	if err != nil {
		return orders.Product{}, errors.Wrap(err, "request to shop failed")
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return orders.Product{}, errors.Wrap(err, "cannot read response")
	}

	productView := shop_http.ProductView{}
	if err := json.Unmarshal(b, &productView); err != nil {
		return orders.Product{}, errors.Wrapf(err, "cannot decode response: %s", b)
	}

	return OrderProductFromHTTP(productView)
}

func OrderProductFromHTTP(shopProduct shop_http.ProductView) (orders.Product, error) {
	productPrice, err := OrderProductPriceFromHTTP(shopProduct.Price)
	if err != nil {
		return orders.Product{}, errors.Wrap(err, "cannot decode price")
	}

	return orders.NewProduct(orders.ProductID(shopProduct.ID), shopProduct.Name, productPrice)
}

func OrderProductPriceFromHTTP(priceView shop_http.PriceView) (price.Price, error) {
	return price.NewPrice(priceView.Cents, priceView.Currency)
}
