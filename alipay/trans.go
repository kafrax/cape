package alipay

import (
	"encoding/json"
	"github.com/kafrax/chaos"
)

//https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer

type Transfer struct {
	OutBizNo     string `json:"out_biz_no"`
	PayeeType    string `json:"payee_type"`
	PayeeAccount string `json:"payee_account"`
	Amount       string `json:"amount"`
}

type TransferResponse struct {
	AliPayFundTransToAccountTransferResponse struct {
		Code     string `json:"code"`
		Msg      string `json:"msg"`
		OrderID  string `json:"order_id"`
		SubCode  string `json:"sub_code"`
		SubMsg   string `json:"sub_msg"`
		OutBizNo string `json:"out_biz_no"`
		PayDate  string `json:"pay_date"`
	} `json:"alipay_fund_trans_toaccount_transfer_response"`
	Sign string `json:"sign"`
}

func (w *Transfer) ChooseApi() string {
	return AliApiTransfer
}

func (w *Transfer) SetExtraParam() (isNil bool, m map[string]string) {
	return true, nil
}

func (w *Transfer) BizContentKey() (isNil bool, key string) {
	return false, "biz_content"
}

func (w *Transfer) BizContentValue() string {
	b, err := json.Marshal(w)
	if err != nil {
		return ""
	}
	return chaos.Byte2String(b)
}
