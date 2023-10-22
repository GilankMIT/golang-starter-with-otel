package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"go-boilerplate/common/util/constants"
	"go-boilerplate/common/util/logutil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
	"strings"
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

func (e ExchangeRequest) GetFullURI() string {
	return e.Host + e.URI
}

type ExchangeResponse struct {
	Code         int         `json:"code"`
	Payload      []byte      `json:"payload"`
	Header       http.Header `json:"header"`
	RequestTime  time.Time   `json:"request_time"`
	ResponseTime time.Time   `json:"response_time"`
}

func Exchange(ctx context.Context, req ExchangeRequest) (resp ExchangeResponse, err error) {

	ctx = context.WithValue(ctx, constants.LOG_APPENDER_CTX_KEY, logutil.HTTP_APPENDER)

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
		statusCode, respBody, err = httpExchangeByMethod(ctx, http.MethodGet, req)
	case http.MethodPost:
		statusCode, respBody, err = httpExchangeByMethod(ctx, http.MethodPost, req)
	}

	resp.ResponseTime = time.Now()

	if err != nil {
		return
	}

	resp.Code = statusCode
	resp.Payload = respBody
	return
}

func httpExchangeByMethod(ctx context.Context, method string, req ExchangeRequest) (statusCode int, respBody []byte, err error) {
	var client = &http.Client{}

	ctx, cancel := context.WithTimeout(ctx, req.Timeout)
	defer cancel()

	var payload = bytes.NewBuffer(req.Payload)
	request, err := http.NewRequestWithContext(ctx, method, req.Host+req.URI, payload)
	if err != nil {
		return 0, nil, err
	}
	request.Header = req.Header

	timeStart := time.Now()
	logutil.LogInfo(ctx, fmt.Sprintf("http %s to %s invoke: %s", strings.ToUpper(req.Method),
		req.GetFullURI(), string(req.Payload)))

	response, err := client.Do(request)
	if err != nil {
		timeCost := time.Now().Sub(timeStart).Milliseconds()
		logutil.LogInfo(ctx, fmt.Sprintf("exception in http %s request to %s : %s, tc: %dms", strings.ToUpper(req.Method),
			req.GetFullURI(), err.Error(), timeCost))
		return 0, nil, err
	}

	defer response.Body.Close()

	respBody, err = io.ReadAll(response.Body)
	if respBody != nil {
		timeCost := time.Now().Sub(timeStart).Milliseconds()
		logutil.LogInfo(ctx, fmt.Sprintf("http %s from %s result: %s, tc: %dms", strings.ToUpper(req.Method),
			req.GetFullURI(), string(respBody), timeCost))
	}
	statusCode = response.StatusCode

	return
}
