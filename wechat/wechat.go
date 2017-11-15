package wechat

//todo Detecting the error of return system to avoid repeated payment

type WechatApi = string

const (
	WechatApiUnifyOrder WechatApi = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	WechatApiTransfer   WechatApi = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
)

type common struct {
	charset string `json:"-"`
	AppId   string `json:"appid"`
	MchId   string `json:"mch_id"`
}

func (c *common) SignAndMarshal(m WePayer) (s string, err error) {
	data := m.GetMapData()
}
