package alipay

import (
	"encoding/json"
	"github.com/kafrax/chaos"
)

//https://docs.open.alipay.com/204/105465/

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

type AppPayResponse struct {
	AppId             string `json:"app_id"`
	AuthAppId         string `json:"auth_app_id"`
	NotifyId          string `json:"notify_id"`
	NotifyType        string `json:"notify_type"`
	NotifyTime        string `json:"notify_time"`
	TradeNo           string `json:"trade_no"`
	TradeStatus       string `json:"trade_status"`
	TotalAmount       string `json:"total_amount"`
	ReceiptAmount     string `json:"receipt_amount"`
	InvoiceAmount     string `json:"invoice_amount"`
	BuyerPayAmount    string `json:"buyer_pay_amount"`
	SellerId          string `json:"seller_id"`
	SellerEmail       string `json:"seller_email"`
	BuyerId           string `json:"buyer_id"`
	BuyerLogonId      string `json:"buyer_logon_id"`
	FundBillList      string `json:"fund_bill_list"`
	Charset           string `json:"charset"`
	PointAmount       string `json:"point_amount"`
	OutTradeNo        string `json:"out_trade_no"`
	OutBizNo          string `json:"out_biz_no"`
	GmtCreate         string `json:"gmt_create"`
	GmtPayment        string `json:"gmt_payment"`
	GmtRefund         string `json:"gmt_refund"`
	GmtClose          string `json:"gmt_close"`
	Subject           string `json:"subject"`
	Body              string `json:"body"`
	RefundFee         string `json:"refund_fee"`
	Version           string `json:"version"`
	SignType          string `json:"sign_type"`
	Sign              string `json:"sign"`
	PassbackParams    string `json:"passback_params"`
	VoucherDetailList string `json:"voucher_detail_list"`
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
	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return chaos.Byte2String(b)
}
