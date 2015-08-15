package rapi

import "net/http"

type Endpoint struct {
	api          *Api
	externalPath string
	internalPath string
	method       string
}

func NewEndpoint(a *Api, prefix, method, path string) *Endpoint {
	ep := Endpoint{
		api:          a,
		externalPath: path,
		internalPath: path,
		method:       method,
	}

	a.proxy.router.
		HandleFunc(prefix+path, ep.Handler).
		Methods(method)

	return &ep
}

func (ep *Endpoint) Handler(w http.ResponseWriter, r *http.Request) {
	ep.api.proxy.Request(ep, w, r)
}

func (ep *Endpoint) InternalPath(path string) *Endpoint {
	ep.internalPath = path
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
