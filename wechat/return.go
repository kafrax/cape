package wechat

type ReturnMsg struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	Recipient
}

func AddRecipient(sub Recipient) ReturnMsg {
	msg := ReturnMsg{}
	msg.Recipient = sub
	return msg
}

type NotifyReturnMsg struct {
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      string `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       string `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

func (*NotifyReturnMsg) Convert2Map(ReturnMsg) map[string]string {
	panic("implement me")
}

func (*NotifyReturnMsg) CheckSign(map[string]string) bool {
	panic("implement me")
}

func (*NotifyReturnMsg) Unmarshaler([]byte, *ReturnMsg) {
	panic("implement me")
}

type UnifyReturnMsg struct {
	PrepayId  string `xml:"prepay_id"`
	TradeType string `xml:"trade_type"`
}

func (*UnifyReturnMsg) Convert2Map(ReturnMsg) map[string]string {
	panic("implement me")
}

func (*UnifyReturnMsg) CheckSign(map[string]string) bool {
	panic("implement me")
}

func (*UnifyReturnMsg) Unmarshaler([]byte, *ReturnMsg) {
	panic("implement me")
}

type TransferReturnMsg struct {
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentNo      string `xml:"payment_no"`
	PaymentTime    string `xml:"payment_time"`
}

func (*TransferReturnMsg) Convert2Map(ReturnMsg) map[string]string {
	panic("implement me")
}

func (*TransferReturnMsg) CheckSign(map[string]string) bool {
	panic("implement me")
}

func (*TransferReturnMsg) Unmarshaler([]byte, *ReturnMsg) {
	panic("implement me")
}

type Write2Wechat struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}
