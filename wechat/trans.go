package wechat

//https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_1
type Transfer struct {
	MchAppId string `json:"mch_appid"`
	MchId string `json:"mchid"`
	NonceStr string `json:"nonce_str"`
	PartnerTradeNo string `json:"partner_trade_no"`
	Openid string `json:"openid"`
	CheckName string `json:"check_name"`
	ReUserName string `json:"re_user_name"`
}