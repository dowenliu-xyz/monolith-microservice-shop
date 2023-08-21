package orders

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPClient struct {
	address string
}

func NewHTTPClient(address string) HTTPClient {
	return HTTPClient{address}
}

func (h HTTPClient) MarkOrderAsPaid(orderID string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders/%s/paid", h.address, orderID), nil)
	if err != nil {
		return errors.Wrap(err, "cannot create request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "request to orders failed")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.Errorf("request to orders failed with status %d", resp.StatusCode)
	}
	return nil
}
