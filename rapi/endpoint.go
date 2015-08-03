package rapi

type Endpoint struct {
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
	return
}

func (ep *Endpoint) TransformRequest(internal, external interface{}) *Endpoint {
	return
}

func (ep *Endpoint) TransformResponse(internal, external interface{}) *Endpoint {
	return
}

func (ep *Endpoint) TransformRequestCb(callback TransformCb) *Endpoint {
	err := callback()
	return
}

func (ep *Endpoint) TransformResponseCb(callback TransformCb) *Endpoint {
	err := callback()
	return
}
