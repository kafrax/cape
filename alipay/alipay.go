package alipay

import (
	"encoding/json"
	"net/url"
	"sort"
	"time"

	"github.com/kafrax/chaos"
	"github.com/kafrax/netask"
	"github.com/kafrax/cape/unite"
)

//todo Detecting the error of return system to avoid repeated payment
type AliApi = string

const (
	AliApiPay           AliApi = "alipay.trade.app.pay"
	AliApiQuery         AliApi = "alipay.trade.query"
	AliApiTransfer      AliApi = "alipay.fund.trans.toaccount.transfer"
	AliApiTransferQuery AliApi = "alipay.fund.trans.order.query"
)

type Alipay struct {
	AppId      string
	privateKey string
	publicKey  string
	domain     string
	format     string
	charset    string
	signType   string
	version    string
	values     *url.Values
}

func New(pap *unite.Papyrus) *Alipay {
	return &Alipay{
		domain:     "https://openapi.alipay.com/gateway.do",
		format:     "JSON",
		charset:    "uft-8",
		signType:   "RSA2",
		version:    "1.0",
		privateKey: pap.PrivateKey,
		publicKey:  pap.PublicKey,
		AppId:      pap.AppID,
	}
}

// user demo:
// alipay/Alipay.AppId="12345678"
// Alipay.UrlValuesEncode(AliPayer)
func (c *Alipay) UrlValuesEncode(a AliPayer) string {
	u := &url.Values{}
	u.Add("app_id", c.AppId)
	u.Add("method", a.ChooseApi())
	u.Add("format", c.format)
	u.Add("charset", c.charset)
	u.Add("sign_type", c.signType)
	u.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	u.Add("version", c.version)
	if ok, v := a.BizContentKey(); !ok {
		u.Add(v, a.BizContentValue())
	}
	if ok, v := a.SetExtraParam(); !ok {
		for key, value := range v {
			u.Add(key, value)
		}
	}
	sig, err := SignRsa2(allKeys(u), u, chaos.String2Byte(c.privateKey))
	if err != nil {
		return ""
	}
	u.Add("sign", sig)
	return u.Encode()
}

func allKeys(u *url.Values) (ret []string) {
	for k, _ := range *u {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return
}

//for http post
func (c *Alipay) Post(method string, obj AliPayer, result interface{}) (err error) {
	resp, err := netask.Post(
		c.domain,
		"application/x-www-form-urlencoded;charset=utf-8",
		false, chaos.String2Byte(c.UrlValuesEncode(obj)),
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, result)
	if err != nil {
		return err
	}
	return
}
