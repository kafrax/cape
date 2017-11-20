package payrus

type ErrorType = uint64

const (
	ErrorServer ErrorType = 1 << 63
	ErrorChannel ErrorType = 1 << 63-1
	ErrorAlipayParams ErrorType = 1 << 63-2
)

type Error struct {
	ErrType ErrorType
	ErrMsg  string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func ErrorsNew(msg string, t ErrorType) error {
	return &Error{t, msg}
}
