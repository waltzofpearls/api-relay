package rapi

type TransformCb func() (err error)

type Api struct {
	prefix string
	proxy  Proxy
}

func New(prefix string) *Api {
	return &Api{
		prefix: prefix,
		proxy:  *NewProxy(),
	}
}

func (a *Api) NewEndpoint(method, endpoint string) *Endpoint {
	return NewEndpoint(&a.proxy, a.prefix, method, endpoint)
}

func (a *Api) Run() {
	a.proxy.Serve()
}
