package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"

	common_http "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/http"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/application"
	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/domain/orders"
)

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

type ordersResource struct {
	service application.OrdersService

	repository orders.Repository
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}

	cmd := application.PlaceOrderCommand{
		OrderID:   orders.ID(uuid.NewV1().String()),
		ProductID: req.ProductID,
		Address:   application.PlaceOrderCommandAddress(req.Address), // 内存布局相同的类型可以直接强转
	}
	if err := o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, PostOrdersResponse{
		OrderID: string(cmd.OrderID),
	})
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   PostOrderAddress `json:"address"`
}

type PostOrdersResponse struct {
	OrderID string
}

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

func (o ordersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	// 使用 service 应该也没什么问题，现在能看到的唯一区别： repository 返回指针，而 service 返回复制值。但后面是使用只读方法，应该没有区别的。
	//order, err := o.service.OrderByID(orders.ID(chi.URLParam(r, "id")))
	order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}

	render.Respond(w, r, OrderPaidView{string(order.ID()), order.Paid()})
}
