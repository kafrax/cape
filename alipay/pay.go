package alipay

import (
	"encoding/json"
	"github.com/kafrax/chaos"
)

type AppPay struct {
	NotifyUrl      string `json:"-"`
	Body           string `json:"body"`
	Subject        string `json:"subject"`
	OutTradeNo     string `json:"out_trade_no"`
	TotalAmount    string `json:"total_amount"`
	ProductCode    string `json:"product_code"`
	PassbackParams string `json:"passback_params"`
	TimeoutExpress string `json:"timeout_express"`
}

func (a *AppPay) ChooseApi() string {
	return AliApiPay
}

func (a *AppPay) SetExtraParam() (isNil bool, m map[string]string) {
	m = make(map[string]string)
	m["notify_url"] = a.NotifyUrl
	return
}

func (a *AppPay) BizContentKey() (isNil bool, key string) {
	return false, "biz_content"
}

func (a *AppPay) BizContentValue() string {
	b,err:=json.Marshal(a)
	if err != nil {
		return ""
	}
	return chaos.Byte2String(b)
}
