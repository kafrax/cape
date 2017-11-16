package unite

type ErrorType uint64

const (
	ErrorServer ErrorType = 1 << 63
)

type Error struct {
	Err     error
	ErrType ErrorType
	ErrMsg  string
}

func (e *Error) Error() string {
	return e.ErrMsg
}