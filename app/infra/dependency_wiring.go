package infra

import (
	paymentServiceIntegration "go-otel-example/app/integration/payment_service"
	integrationShared "go-otel-example/app/integration/shared"
	orderService "go-otel-example/app/service/order_service"
	serviceShared "go-otel-example/app/service/shared"
)

func GetOptions() (serviceShared.ServiceOptions, error) {
	clientOption, err := GetIntegration()
	if err != nil {
		return serviceShared.ServiceOptions{}, err
	}
	return serviceShared.ServiceOptions{
		orderService.NewOrderService(clientOption.PaymentServiceClient),
	}, nil
}

func GetIntegration() (integrationShared.ClientOption, error) {
	return integrationShared.ClientOption{
		paymentServiceIntegration.NewPaymentServiceClient(),
	}, nil
}
