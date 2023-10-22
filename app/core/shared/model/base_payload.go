package model

type BaseServiceResponse struct {
	IsSuccess    bool   `json:"is_success"`
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
	ResponseTime int64  `json:"response_time"`
}
