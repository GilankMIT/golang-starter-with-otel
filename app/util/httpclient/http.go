package httpclient

import (
	"context"
	"time"
)

type ExchangeRequest struct {
	Host           string `json:"host"`
	URI            string `json:"uri"`
	Method         string `json:"method"`
	Payload        []byte `json:"payload"`
	ContextTimeout bool   `json:"context_timeout"`
	//Applied only if ContextTimeout false
	Timeout time.Duration `json:"timeout"`
}

type ExchangeResponse struct {
	Code         int       `json:"code"`
	Payload      []byte    `json:"payload"`
	RequestTime  time.Time `json:"request_time"`
	ResponseTime time.Time `json:"response_time"`
}

func Exchange(ctx context.Context, req ExchangeRequest) (ExchangeResponse, error) {

}
