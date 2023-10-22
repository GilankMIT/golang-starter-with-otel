package payment_service

import (
	"context"
	"errors"
	"go-otel-example/app/util/httpclient"
	"go-otel-example/app/util/template"
	"net/http"
	"time"
)

const SERVICE_NAME = "integration.OrderService"

type PayRequest struct {
	TrxId string `json:"trx_id"`
}

type PayResponse struct {
	TrxId      string `json:"trx_id"`
	Status     int    `json:"status"`
	StatusDesc string `json:"status_desc"`
}

var StatusCode = map[int]string{
	1:  "INIT",
	2:  "PENDING",
	3:  "SUCCESS",
	99: "FAILED",
}

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
			httpclient.Exchange(ctx, exchangeRequest)
			resp = PayResponse{
				TrxId:      req.TrxId,
				Status:     3,
				StatusDesc: StatusCode[3],
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
		URI:     "api/v1/test",
		Method:  http.MethodGet,
		Payload: nil,
		Header:  nil,
		Timeout: time.Second * 10,
	}
}
