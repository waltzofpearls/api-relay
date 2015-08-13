package rapi

import "net/http"

type Endpoint struct {
	proxy    *Proxy
	method   string
	endpoint string
}

func NewEndpoint(proxy *Proxy, prefix, method, endpoint string) *Endpoint {
	ep := &Endpoint{
		proxy:    proxy,
		method:   method,
		endpoint: endpoint,
	}

	proxy.router.
		HandleFunc(prefix+endpoint, ep.Handler).
		Methods(method)

	return ep
}

func (ep *Endpoint) Handler(w http.ResponseWriter, r *http.Request) {
	ep.proxy.Request(ep, w, r)
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
