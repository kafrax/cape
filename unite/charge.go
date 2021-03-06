package unite

import (
	"github.com/kafrax/papyrus"
	"github.com/kafrax/papyrus/alipay"
)

type ChargeRequest struct {
	OrderNo  string `json:"order_no"`
	Channel  string `json:"channel"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	ClientIP string `json:"client_ip"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}


func chooseAdapter(c string,pap *Papyrus)(ChargeAdapter,error){
	switch c {
	case "alipay.app":
		return alipay.New(pap)
	}

	return nil,payrus.ErrorsNew("no channel",payrus.ErrorChannel)
}

func (c *ChargeRequest) Validator(charge ChargeRequest) bool {
	panic("implement me")
}

func (c *ChargeRequest) Convert2SP(charge ChargeRequest) ChargeAdapter {
	panic("implement me")
}

func (c *ChargeRequest) Pay() {
	panic("implement me")
}

func (c *ChargeRequest) Convert2C(ChargeAdapter) ChargeResponse {
	panic("implement me")
}


//convert ChargeRequest to ChargeResponse
func New(request *ChargeRequest, pap *Papyrus) (c *ChargeResponse,err error) {
	//chooseAdapter
	ad,err:=chooseAdapter(request.Channel,pap)
	if err != nil {

	}
	//validator ChargeRequest is correct
	//convert ChargeRequest to ChargeAdapter
	//pay action
	//convert ChargeAdapter to ChargeResponse

}

//for pay
type ChargeResponse struct {
	ID           string `json:"id"`
	Object       string `json:"object"`
	Created      int    `json:"created"`
	Livemode     bool   `json:"livemode"`
	Paid         bool   `json:"paid"`
	Refunded     bool   `json:"refunded"`
	Reversed     bool   `json:"reversed"`
	App          string `json:"app"`
	Channel      string `json:"channel"`
	OrderNo      string `json:"order_no"`
	ClientIP     string `json:"client_ip"`
	Amount       int    `json:"amount"`
	AmountSettle int    `json:"amount_settle"`
	Currency     string `json:"currency"`
	Subject      string `json:"subject"`
	Body         string `json:"body"`
	Extra struct {
	} `json:"extra"`
	TimePaid      interface{} `json:"time_paid"`
	TimeExpire    int         `json:"time_expire"`
	TimeSettle    interface{} `json:"time_settle"`
	TransactionNo interface{} `json:"transaction_no"`
	Refunds struct {
		Object  string        `json:"object"`
		URL     string        `json:"url"`
		HasMore bool          `json:"has_more"`
		Data    []interface{} `json:"data"`
	} `json:"refunds"`
	AmountRefunded int         `json:"amount_refunded"`
	FailureCode    interface{} `json:"failure_code"`
	FailureMsg     interface{} `json:"failure_msg"`
	Credential struct {
		Object string `json:"object"`
		Upacp struct {
			Tn   string `json:"tn"`
			Mode string `json:"mode"`
		} `json:"upacp"`
	} `json:"credential"`
	Description interface{} `json:"description"`
}
