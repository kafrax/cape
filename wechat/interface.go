package wechat

type WePayer interface {
	ChooseHost() string
	SetExtraParam(key, value string)
	SetActionParam()
	DoRequest(body []byte) ([]byte, error)
	GetMapData() map[string]string
}
