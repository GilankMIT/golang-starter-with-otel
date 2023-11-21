package enum

var (
	ASSET_TYPE_BALANCE      = NewAssetTypeEnum("BALANCE", "prepaid balance of user")
	ASSET_TYPE_INNERACCOUNT = NewAssetTypeEnum("INNERACCOUNT", "internal account")
)

type AssetTypeEnum struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func NewAssetTypeEnum(code, desc string) AssetTypeEnum {
	return AssetTypeEnum{
		Code: code,
		Desc: desc,
	}
}
