package alipay

import (
	"encoding/json"

	"github.com/kafrax/chaos"
)

//https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query/

type TransferQuery struct {
	OutBizNo string `json:"out_biz_no"`
	OrderId  string `json:"order_id"`
}

func (*TransferQuery) ChooseApi() string {
	return AliApiTransferQuery
}

func (*TransferQuery) SetExtraParam() (isNil bool, m map[string]string) {
	return true, nil
}

func (*TransferQuery) BizContentKey() (isNil bool, key string) {
	return false, "biz_content"
}

func (t *TransferQuery) BizContentValue() string {
	b, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return chaos.Byte2String(b)
}

type TransferQueryResponse struct {
	AliPayFundTransOrderQueryResponse struct {
		ArrivalTimeEnd string `json:"arrival_time_end"`
		Code           string `json:"code"`
		ErrorCode      string `json:"error_code"`
		FailReason     string `json:"fail_reason"`
		Msg            string `json:"msg"`
		OrderFee       string `json:"order_fee"`
		OrderID        string `json:"order_id"`
		SubCode        string `json:"sub_code"`
		SubMsg         string `json:"sub_msg"`
		OutBizNo       string `json:"out_biz_no"`
		PayDate        string `json:"pay_date"`
		Status         string `json:"status"`
	} `json:"alipay_fund_trans_order_query_response"`
	Sign string `json:"sign"`
}
