package infra

import (
	"github.com/gorilla/mux"
	"go-boilerplate/common/util/middleware"
	orderPayment "go-boilerplate/presentation/order_payment"
	"go-boilerplate/service/shared"
	"log"
	"net/http"
	"time"
)

func HandleBasicRouter(infraOption InfraOption) (*http.Server, error) {

	r := mux.NewRouter()

	serviceOption, err := GetOptions(infraOption)
	if err != nil {
		return nil, err
	}

	registerMiddleware(r)
	registerRoutes(r, &serviceOption)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	return srv, nil
}

func registerMiddleware(r *mux.Router) {
	r.Use(middleware.WithOtelTrace)
}

func registerRoutes(r *mux.Router, options *shared.ServiceOptions) {
	orderPayment.NewOrderPaymentHTTPController(r, options)
}
