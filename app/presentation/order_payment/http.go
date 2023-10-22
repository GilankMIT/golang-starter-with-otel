package order_payment

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-boilerplate/service/shared"
	"net/http"
)

type orderPaymentHttpController struct {
	r             *mux.Router
	serviceOption *shared.ServiceOptions
}

func NewOrderPaymentHTTPController(r *mux.Router, serviceOption *shared.ServiceOptions) {
	httpController := &orderPaymentHttpController{
		r:             r,
		serviceOption: serviceOption,
	}

	httpController.r.HandleFunc("/pay", httpController.pay).Methods(http.MethodGet)
}

func (ctrl orderPaymentHttpController) pay(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	resp := ctrl.serviceOption.OrderService.Pay(ctx, "12345")
	json.NewEncoder(writer).Encode(resp)
}
