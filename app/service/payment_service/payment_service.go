package payment_service

import (
	"context"
	"errors"
	payment_service2 "go-otel/app/integration/payment_service"
	template2 "go-otel/app/util/template"
)

const SERVICE_NAME = "service.PaymentService"

type PaymentService interface {
	Pay(ctx context.Context, trxId string) (resp string, err error)
}

type PaymentServiceImpl struct {
	paymentServiceClient payment_service2.PaymentServiceClient
}

func NewPaymentService(client payment_service2.PaymentServiceClient) PaymentService {
	return &PaymentServiceImpl{
		client,
	}
}

func (p PaymentServiceImpl) Pay(ctx context.Context, trxId string) (resp string, err error) {

	err = template2.ServiceExec(ctx,
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

func (p PaymentServiceImpl) payValidate(trxId string) error {
	if trxId == "" {
		return errors.New("trxId cannot be empty")
	}
	return nil
}

func (p PaymentServiceImpl) payInnerServiceProcess(ctx context.Context, trxId string) (resp string, err error) {
	payRequest := payment_service2.PayRequest{trxId}
	paymentResp, err := p.paymentServiceClient.Pay(ctx, payRequest)
	if err != nil {
		return "", err
	}
	resp = paymentResp.StatusDesc

	return resp, nil
}
