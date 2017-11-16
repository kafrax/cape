package wechat

// WePayer wechat action interface
type WePayer interface {
	ChooseHost() string
	SetExtraParam(key, value string)
	SetActionParam()
	GetMapData() map[string]string
}

// recipient for all of wechat return msg
type Recipient interface {
	Convert2Map(ReturnMsg) map[string]string
	CheckSign(map[string]string) bool
	Unmarshaler([]byte, *ReturnMsg)
}
