package order_service

import "go-boilerplate/core/shared/model"

type OrderPayResponse struct {
	Status string `json:"status"`
	model.BaseServiceResponse
}
