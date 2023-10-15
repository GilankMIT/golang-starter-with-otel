package order_service

import (
	"context"
	"errors"
	"go-otel/app/integration/payment_service"
	"go-otel/app/util/template"
)

const SERVICE_NAME = "service.OrderService"

type OrderService interface {
	Pay(ctx context.Context, trxId string) (resp string, err error)
}

type orderServiceImpl struct {
	paymentServiceClient payment_service.PaymentServiceClient
}

func NewOrderService(client payment_service.PaymentServiceClient) OrderService {
	return &orderServiceImpl{
		client,
	}
}

func (p orderServiceImpl) Pay(ctx context.Context, trxId string) (resp string, err error) {

	err = template.ServiceTemplateExec(ctx,
		SERVICE_NAME,
		trxId,

		//preCheck
		func(request any) error {
			return p.payValidate(trxId)
		},

		//process
		func(ctx context.Context, request any) (any, error) {
			resp, err = p.payInnerServiceProcess(ctx, trxId)
			return resp, err
		}, nil,
	)

	return
}

func (p orderServiceImpl) payValidate(trxId string) error {
	if trxId == "" {
		return errors.New("trxId cannot be empty")
	}
	return nil
}

func (p orderServiceImpl) payInnerServiceProcess(ctx context.Context, trxId string) (resp string, err error) {
	payRequest := payment_service.PayRequest{trxId}
	paymentResp, err := p.paymentServiceClient.Pay(ctx, payRequest)
	if err != nil {
		return "", err
	}
	resp = paymentResp.StatusDesc

	return resp, nil
}
