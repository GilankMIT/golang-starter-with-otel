package infra

import (
	payment_service2 "go-otel/app/integration/payment_service"
	shared2 "go-otel/app/integration/shared"
	payment_service3 "go-otel/app/service/payment_service"
	shared3 "go-otel/app/service/shared"
)

func GetOptions() (shared3.ServiceOptions, error) {
	clientOption, err := GetIntegration()
	if err != nil {
		return shared3.ServiceOptions{}, err
	}
	return shared3.ServiceOptions{
		payment_service3.NewPaymentService(clientOption.PaymentServiceClient),
	}, nil
}

func GetIntegration() (shared2.ClientOption, error) {
	return shared2.ClientOption{
		payment_service2.NewPaymentServiceClient(),
	}, nil
}
