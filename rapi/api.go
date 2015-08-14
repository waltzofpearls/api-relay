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

func (a *Api) setListenProtocol(protocol string) *Api {
	a.proxy.setListenProtocol(protocol)
	return a
}

func (a *Api) setListenAddr(addr string) *Api {
	a.proxy.setListenAddr(addr)
	return a
}

func (a *Api) setDownstreamProtocol(protocol string) *Api {
	a.proxy.setDownstreamProtocol(protocol)
	return a
}

func (a *Api) setDownstreamAddr(addr string) *Api {
	a.proxy.setDownstreamAddr(addr)
	return a
}

func (a *Api) NewEndpoint(method, endpoint string) *Endpoint {
	return NewEndpoint(&a.proxy, a.prefix, method, endpoint)
}

func (a *Api) Run() {
	a.proxy.Serve()
}
