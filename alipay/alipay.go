package alipay

import (
	"net/url"
	"time"
	"sort"
	"github.com/kafrax/chaos"
)

type AliApi = string

const (
	AliApiPay   AliApi = "alipay.trade.app.pay"
	AliApiQuery AliApi = "alipay.trade.query"
)

type config struct {
	domain     string
	AppId      string
	privateKey string
	format     string
	charset    string
	signType   string
	publicKey  string
	version    string
	values     *url.Values
}

var AliPay *config

func init() {
	AliPay = &config{
		domain:   "https://openapi.alipay.com/gateway.do",
		format:   "JSON",
		charset:  "uft-8",
		signType: "RSA2",
		version:  "1.0",
	}
}

// user demo:
// alipay/Alipay.AppId="12345678"
// Alipay.UrlValuesEncode(AliPayer)
func (c *config) UrlValuesEncode(a AliPayer) (string, error) {
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
		return "", err
	}
	u.Add("sign", sig)
	return u.Encode(), nil
}

func allKeys(u *url.Values) (ret []string) {
	for k, _ := range *u {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return
}
