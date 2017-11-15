package wechat

type WePayer interface {
	ChooseHost() string
	SetExtraParam(key, value string)
	SetActionParam()
	GetMapData() map[string]string
}
