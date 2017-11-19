package unite

type IPapyrus interface {
}

type Papyrus struct {
	AppID      string
	MchID      string
	PrivateKey string
	PublicKey  string
	AppKey     string
	AppSecret  string
}