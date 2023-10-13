package payment_service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	shared2 "go-otel/app/service/shared"
	"net/http"
)

type paymentServiceHttpController struct {
	r             *mux.Router
	serviceOption *shared2.ServiceOptions
}

func NewPaymentServiceHTTPController(r *mux.Router, serviceOption *shared2.ServiceOptions) {
	httpController := &paymentServiceHttpController{
		r:             r,
		serviceOption: serviceOption,
	}

	httpController.r.HandleFunc("/pay", httpController.pay).Methods(http.MethodGet)
}

func (ctrl paymentServiceHttpController) pay(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	resp, err := ctrl.serviceOption.PaymentService.Pay(ctx, "12345")
	if err != nil {
		json.NewEncoder(writer).Encode(map[string]any{"success": false, "resp": nil})
	}

	json.NewEncoder(writer).Encode(map[string]any{"success": true, "resp": resp})
}
