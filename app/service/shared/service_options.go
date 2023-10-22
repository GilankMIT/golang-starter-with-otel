package shared

import (
	orderService "go-otel-example/app/service/order_service"
)

type ServiceOptions struct {
	OrderService orderService.OrderService
}
