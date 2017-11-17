package payrus

type IPapyrus interface {
}

type Papyrus struct {
	middleware []Endpoint
}

func (p *Papyrus) Add(endpoint Endpoint) {
	p.middleware = append(p.middleware, endpoint)
}

func (p *Papyrus) Next(nextFunc NextFunc) {

}

func annotate(endpoint Endpoint) Middleware {
	return func(next Endpoint) Endpoint {
		return func(ctx *Context, request interface{}) (interface{}, error) {
			resp, err := endpoint(ctx, request)
			if err != nil {
				return nil, nil
			}
			return next(ctx, resp)
		}
	}
}

// Endpoint is the middleware function
type Endpoint = func(ctx *Context, request interface{}) (response interface{}, err error)

// _Nop Useful for tests.
func _Nop(*Context, interface{}) (interface{}, error) { return struct{}{}, nil }

// to add middleware what you want to do when pay channel action will doing before.
type Middleware = func(Endpoint) Endpoint

func (p *Papyrus) chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Endpoint) Endpoint {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
