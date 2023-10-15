package infra

import (
	payment_service2 "go-otel/app/integration/payment_service"
	integrationShared "go-otel/app/integration/shared"
	payment_service3 "go-otel/app/service/order_service"
	serviceShared "go-otel/app/service/shared"
)

func GetOptions() (serviceShared.ServiceOptions, error) {
	clientOption, err := GetIntegration()
	if err != nil {
		return serviceShared.ServiceOptions{}, err
	}
	return serviceShared.ServiceOptions{
		payment_service3.NewOrderService(clientOption.PaymentServiceClient),
	}, nil
}

func GetIntegration() (integrationShared.ClientOption, error) {
	return integrationShared.ClientOption{
		payment_service2.NewPaymentServiceClient(),
	}, nil
}
