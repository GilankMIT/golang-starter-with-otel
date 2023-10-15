package shared

import (
	payment_service2 "go-otel/app/service/order_service"
)

type ServiceOptions struct {
	PaymentService payment_service2.OrderService
}
