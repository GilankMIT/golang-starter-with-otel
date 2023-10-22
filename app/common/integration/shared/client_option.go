package shared

import (
	"go-boilerplate/common/integration/payment_service"
)

type ClientOption struct {
	payment_service.PaymentServiceClient
}
