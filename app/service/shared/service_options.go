package shared

import (
	orderService "go-boilerplate/service/order_service"
)

type ServiceOptions struct {
	OrderService orderService.OrderService
}
