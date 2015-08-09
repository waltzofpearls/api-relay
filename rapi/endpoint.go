package rapi

type Endpoint struct {
	proxy    *Proxy
	method   string
	endpoint string
}

func NewEndpoint(proxy *Proxy, method, endpoint string) *Endpoint {
	return &Endpoint{
		proxy:    proxy,
		method:   method,
		endpoint: endpoint,
	}
}

func (ep *Endpoint) InternalPath(endpoint string) *Endpoint {
	return ep
}

func (ep *Endpoint) TransformRequest(internal, external interface{}) *Endpoint {
	return ep
}

func (ep *Endpoint) TransformResponse(internal, external interface{}) *Endpoint {
	return ep
}

func (ep *Endpoint) TransformRequestCb(callback TransformCb) *Endpoint {
	err := callback()
	if err != nil {
		panic("Something went wrong")
	}
	return ep
}

func (ep *Endpoint) TransformResponseCb(callback TransformCb) *Endpoint {
	err := callback()
	if err != nil {
		panic("Something went wrong")
	}
	return ep
}
