package payment_service

type PayRequest struct {
	TrxId string `json:"trx_id"`
}

type PayResponse struct {
	TrxId      string `json:"trx_id"`
	Status     int    `json:"status"`
	StatusDesc string `json:"status_desc"`
}
