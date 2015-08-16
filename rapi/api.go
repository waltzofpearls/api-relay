package rapi

type TransformCb func() (err error)

type Api struct {
	prefix string
	proxy  *Proxy
	config *Config
}

func New(prefix string, config *Config) *Api {
	a := new(Api)
	a.config = config
	a.prefix = prefix
	a.proxy = NewProxy(a.config)
	return a
}

func (a *Api) Run() {
	a.proxy.Serve()
}

func (a *Api) NewEndpoint(method, endpoint string) *Endpoint {
	return NewEndpoint(a, a.prefix, method, endpoint)
}
