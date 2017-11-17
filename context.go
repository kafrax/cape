package payrus

import "math"

const (
	abortIndex uint8 = math.MaxUint8 / 2
)

// when done papyrus action,do next like update order
type NextFunc func(context *Context)
type NextList []NextFunc

// Context is all of the pay channel to use for execute operation like pay or transfer etc.
type Context struct {
	next       NextList
	errors     Error
	abortIndex uint8 //now handlerFunc's index.
	params     map[string]string
}

func (c *Context) Reset() {

}

func (c *Context) Clone() *Context {
	var cp = *c
	cp.next = nil
	cp.errors = Error{}
	cp.abortIndex = -1
	return &cp
}

func (c *Context) Next() {

}

func (c *Context) Abort() {
	c.abortIndex = abortIndex
}

func (c *Context) AbortWithErr(msg string, t ErrorType, err error) Error {
	c.Abort()
	return ErrorsNew(msg, t, err)
}

func (c *Context) IsAbort() bool {
	return c.abortIndex >= abortIndex
}

func (c *Context) NextLen() int {
	return len(c.next)
}

//set other param for action
func (c *Context) SetParam(key, value string) {

}

//get other param from pay channel
func (c *Context) GetParam(key string) string {
	return ""
}
