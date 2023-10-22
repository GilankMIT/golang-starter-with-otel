package payment_service

import (
	"context"
	"errors"
	"go-boilerplate/common/util/httpclient"
	"go-boilerplate/common/util/template"
	domainConstant "go-boilerplate/core/domain/constants"
	"net/http"
	"time"
)

const SERVICE_NAME = "integration.OrderService"

type PaymentServiceClient interface {
	Pay(ctx context.Context, req PayRequest) (resp PayResponse, err error)
}

type PaymentServiceClientImpl struct{}

func NewPaymentServiceClient() PaymentServiceClient {
	return &PaymentServiceClientImpl{}
}

func (p PaymentServiceClientImpl) Pay(ctx context.Context, req PayRequest) (resp PayResponse, err error) {

	err = template.IntegrationTemplateExec(ctx,
		SERVICE_NAME,
		req,

		//preCheck
		func(request any) error {
			return p.validatePay(req)
		},

		//process
		func(ctx context.Context, request any) (any, error) {
			exchangeRequest := p.buildExchangeRequest(request.(PayRequest))
			_, err = httpclient.Exchange(ctx, exchangeRequest)
			resp = PayResponse{
				TrxId:      req.TrxId,
				Status:     3,
				StatusDesc: domainConstant.StatusCode[3],
			}
			return resp, nil
		}, nil,
	)

	return
}

func (p PaymentServiceClientImpl) validatePay(req PayRequest) error {
	if req.TrxId == "" {
		return errors.New("trxId cannot be empty")
	}
	return nil
}

func (p PaymentServiceClientImpl) buildExchangeRequest(req PayRequest) httpclient.ExchangeRequest {
	return httpclient.ExchangeRequest{
		Host:    "http://example.com",
		URI:     "/api/v1/test",
		Method:  http.MethodGet,
		Payload: nil,
		Header:  nil,
		Timeout: time.Second * 10,
	}
}
