package order_service

import (
	"context"
	"errors"
	paymentServiceIntegration "go-boilerplate/common/integration/payment_service"
	"go-boilerplate/common/util/logutil"
	"go-boilerplate/common/util/template"
	"go-boilerplate/core/service/cache"
	coreModel "go-boilerplate/core/shared/model"
	orderServiceValidation "go-boilerplate/core/validation/order_service_validation"
)

const SERVICE_NAME = "service.OrderService"

type OrderService interface {
	Pay(ctx context.Context, trxId string) (resp OrderPayResponse)
}

type orderServiceImpl struct {
	paymentServiceClient paymentServiceIntegration.PaymentServiceClient
	cacheClient          cache.CacheService
}

func NewOrderService(paymentServiceClient paymentServiceIntegration.PaymentServiceClient,
	cacheService cache.CacheService) OrderService {
	return &orderServiceImpl{
		paymentServiceClient,
		cacheService,
	}
}

func (p orderServiceImpl) Pay(ctx context.Context, trxId string) (resp OrderPayResponse) {

	var err error

	err = template.ServiceTemplateExec(ctx,
		SERVICE_NAME,
		trxId,

		//preCheck
		func(request any) error {
			return p.payValidate(trxId)
		},

		//process
		func(ctx context.Context, request any) (any, error) {
			resp.Status, err = p.payInnerServiceProcess(ctx, trxId)
			if err != nil {
				logutil.LogError(ctx, "failed when doing pay", err.Error())
				return resp, err
			}

			p.cacheClient.SetNamespace(SERVICE_NAME).SetString(ctx, "payResult", resp.Status)
			return resp, err
		},

		//post result
		func(request, result any) {
			coreModel.ResolveAppError(err, &resp.BaseServiceResponse)
			p.cacheClient.SetNamespace(SERVICE_NAME).Clean(ctx)
		},
	)

	return
}

func (p orderServiceImpl) payValidate(trxId string) (err error) {
	if trxId == "" {
		return errors.New("trxId cannot be empty")
	}

	err = orderServiceValidation.TrxIdIsPayable(trxId)

	return err
}

func (p orderServiceImpl) payInnerServiceProcess(ctx context.Context, trxId string) (resp string, err error) {
	payRequest := paymentServiceIntegration.PayRequest{TrxId: trxId}
	paymentResp, err := p.paymentServiceClient.Pay(ctx, payRequest)
	if err != nil {
		return "", err
	}
	resp = paymentResp.StatusDesc

	return resp, nil
}
