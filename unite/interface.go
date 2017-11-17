package unite

type ChargeAdapter interface {
	Validator(charge ChargeRequest) bool           //valid msg from C
	Convert2SP(charge ChargeRequest) ChargeAdapter //convert msg C2Server provider
	Pay()                                          //pay action
	Convert2C(ChargeAdapter) ChargeResponse        //return msg and convert to C
}

type TransferAdapter interface {
	
}