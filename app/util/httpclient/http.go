package httpclient

import (
	"bytes"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
	"time"
)

var (
	WITH_PROPAGATION = true
)

type ExchangeRequest struct {
	Host    string        `json:"host"`
	URI     string        `json:"uri"`
	Method  string        `json:"method"`
	Payload []byte        `json:"payload"`
	Header  http.Header   `json:"header"`
	Timeout time.Duration `json:"timeout"`
}

type ExchangeResponse struct {
	Code         int         `json:"code"`
	Payload      []byte      `json:"payload"`
	Header       http.Header `json:"header"`
	RequestTime  time.Time   `json:"request_time"`
	ResponseTime time.Time   `json:"response_time"`
}

func Exchange(ctx context.Context, req ExchangeRequest) (resp ExchangeResponse, err error) {

	if req.Header == nil {
		req.Header = map[string][]string{}
	}

	if WITH_PROPAGATION {
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	}

	resp.RequestTime = time.Now()

	var statusCode int
	var respBody []byte

	switch req.Method {
	case http.MethodGet:
		statusCode, respBody, err = get(ctx, req)
	case http.MethodPost:
		statusCode, respBody, err = post(ctx, req)
	}

	resp.ResponseTime = time.Now()

	if err != nil {
		return
	}

	resp.Code = statusCode
	resp.Payload = respBody
	return
}

func get(ctx context.Context, req ExchangeRequest) (statusCode int, respBody []byte, err error) {
	var client = &http.Client{}

	ctx, cancel := context.WithTimeout(ctx, req.Timeout)
	defer cancel()

	var payload = bytes.NewBuffer(req.Payload)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, req.Host+req.URI, payload)
	if err != nil {
		return 0, nil, err
	}
	request.Header = req.Header

	response, err := client.Do(request)
	if err != nil {
		return 0, nil, err
	}

	defer response.Body.Close()

	respBody, err = io.ReadAll(response.Body)
	statusCode = response.StatusCode

	return
}

func post(ctx context.Context, req ExchangeRequest) (statusCode int, respBody []byte, err error) {
	var client = &http.Client{}

	ctx, cancel := context.WithTimeout(ctx, req.Timeout)
	defer cancel()

	var payload = bytes.NewBuffer(req.Payload)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, req.Host+req.URI, payload)
	if err != nil {
		return 0, nil, err
	}
	request.Header = req.Header

	response, err := client.Do(request)
	if err != nil {
		return 0, nil, err
	}

	defer response.Body.Close()

	respBody, err = io.ReadAll(response.Body)
	statusCode = response.StatusCode

	return
}
