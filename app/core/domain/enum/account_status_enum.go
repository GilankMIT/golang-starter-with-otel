package enum

var (
	ACCOUNT_STATUS_NORMAL = NewAccountStatusEnum("N", "account status normal")
	ACCOUNT_STATUS_FROZEN = NewAccountStatusEnum("F", "account status frozen")
)

type AccountStatusEnum struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func NewAccountStatusEnum(code, desc string) AccountStatusEnum {
	return AccountStatusEnum{
		Code: code,
		Desc: desc,
	}
}
