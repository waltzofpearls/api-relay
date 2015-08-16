package rapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type TransformCb func() (err error)

type Endpoint struct {
	config       *Config
	path         string
	reqExtStruct interface{}
	reqIntStruct interface{}
	resExtStruct interface{}
	resIntStruct interface{}
	transformer  Transformable
}

func NewEndpoint(a *Api, method, path string) *Endpoint {
	ep := &Endpoint{
		config:      a.config,
		path:        path,
		transformer: a.transformer,
	}

	a.Route(method, ep.config.Listener.Prefix+path, ep)

	return ep
}

func (ep *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tr := http.DefaultTransport

	r.URL.Host = ep.config.Backend.Address
	r.URL.Scheme = "http"
	r.URL.Path = ep.config.Backend.Prefix + ep.path

	if ep.reqExtStruct != nil {
		ep.transformer.TransformRequest(r, ep.reqExtStruct)
	}

	res, resErr := tr.RoundTrip(r)
	if resErr != nil {
		panic(fmt.Sprintf("Response error: %s", resErr))
	} else {
		defer res.Body.Close()
	}

	if ep.resExtStruct != nil {
		ep.transformer.TransformResponse(res, ep.resExtStruct)
	}

	w.WriteHeader(res.StatusCode)
	_, ioErr := io.Copy(w, res.Body)

	if ioErr != nil {
		log.Printf("Error writting response: %s", ioErr)
	}
}

func (ep *Endpoint) InternalPath(path string) *Endpoint {
	ep.path = path
	return ep
}

func (ep *Endpoint) TransformRequest(external, internal interface{}) *Endpoint {
	ep.reqExtStruct = external
	ep.reqIntStruct = internal
	return ep
}

func (ep *Endpoint) TransformResponse(external, internal interface{}) *Endpoint {
	ep.resExtStruct = external
	ep.resIntStruct = internal
	return ep
}

// To be done
func (ep *Endpoint) TransformRequestCb(cb TransformCb) *Endpoint {
	err := cb()
	if err != nil {
		log.Print("Something went wrong")
	}
	return ep
}

// To be done
func (ep *Endpoint) TransformResponseCb(cb TransformCb) *Endpoint {
	err := cb()
	if err != nil {
		log.Print("Something went wrong")
	}
	return ep
}
