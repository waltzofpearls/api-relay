package rapi

type TransformCb func() (err error)

type Api struct {
	endpoint string
	proxy    Proxy
}

func New(endpoint string) *Api {
	return &Api{
		endpoint: "/api" + endpoint,
		proxy:    &Proxy{},
	}
}

func (a *Api) NewEndpoint(method, endpoint string) *Api {
	return NewEndpoint(method, endpoint)
}
