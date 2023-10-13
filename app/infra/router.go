package infra

import (
	"github.com/gorilla/mux"
	payment_service2 "go-otel/app/delivery/payment_service"
	shared2 "go-otel/app/service/shared"
	middleware2 "go-otel/app/util/middleware"
	"log"
	"net/http"
	"time"
)

func HandleBasicRouter() (*http.Server, error) {

	r := mux.NewRouter()

	serviceOption, err := GetOptions()
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
	r.Use(middleware2.WithOtelTrace)
}

func registerRoutes(r *mux.Router, options *shared2.ServiceOptions) {
	payment_service2.NewPaymentServiceHTTPController(r, options)
}
