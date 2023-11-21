package fund_direction

type DirectionEnum struct {
	Code string
	Desc string
}

func (DirectionEnum) NewDirectionEnum(code, desc string) DirectionEnum {
	return DirectionEnum{
		Code: code,
		Desc: desc,
	}
}
