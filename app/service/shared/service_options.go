package shared

import (
	payment_service2 "go-otel/app/service/payment_service"
)

type ServiceOptions struct {
	PaymentService payment_service2.PaymentService
}
