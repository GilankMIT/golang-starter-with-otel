package infra

import (
	"go-boilerplate/common/integration/payment_service"
	"go-boilerplate/common/integration/shared"
	orderService "go-boilerplate/service/order_service"
	serviceShared "go-boilerplate/service/shared"
)

func GetOptions(infraOption InfraOption) (serviceShared.ServiceOptions, error) {
	clientOption, err := GetIntegration()
	if err != nil {
		return serviceShared.ServiceOptions{}, err
	}
	return serviceShared.ServiceOptions{
		orderService.NewOrderService(clientOption.PaymentServiceClient, infraOption.CacheService),
	}, nil
}

func GetIntegration() (shared.ClientOption, error) {
	return shared.ClientOption{
		payment_service.NewPaymentServiceClient(),
	}, nil
}
