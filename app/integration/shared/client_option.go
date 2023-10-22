package shared

import (
	paymentServiceIntegration "go-otel-example/app/integration/payment_service"
)

type ClientOption struct {
	paymentServiceIntegration.PaymentServiceClient
}
