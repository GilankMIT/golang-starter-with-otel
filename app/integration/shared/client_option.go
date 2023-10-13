package shared

import (
	"go-otel/app/integration/payment_service"
)

type ClientOption struct {
	payment_service.PaymentServiceClient
}
