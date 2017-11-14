package alipay

type AliPayer interface {
	ChooseApi() string
	SetExtraParam() (isNil bool, m map[string]string)
	BizContentKey() (isNil bool, key string)
	BizContentValue() string
}
