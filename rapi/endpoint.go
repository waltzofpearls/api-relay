package rapi

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
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

	ep.CopyUrlVars(r)

	r.URL.Host = ep.config.Backend.Address
	r.URL.Scheme = "http"
	r.URL.Path = ep.config.Backend.Prefix + ep.path

	if ep.reqIntStruct != nil && ep.reqExtStruct != nil {
		ep.transformer.TransformRequest(r, &ep.reqExtStruct, &ep.reqIntStruct)
	}

	res, resErr := tr.RoundTrip(r)
	if resErr != nil {
		panic(fmt.Sprintf("Response error: %s", resErr))
	} else {
		defer res.Body.Close()
	}

	if ep.resIntStruct != nil && ep.resExtStruct != nil {
		ep.transformer.TransformResponse(res, &ep.resIntStruct, &ep.resExtStruct)
	}

	w.WriteHeader(res.StatusCode)
	_, ioErr := io.Copy(w, res.Body)

	if ioErr != nil {
		log.Printf("Error writting response: %s", ioErr)
	}
}

func (ep *Endpoint) CopyUrlVars(r *http.Request) {
	var path bytes.Buffer

	vars := mux.Vars(r)
	if len(vars) == 0 {
		return
	}

	t, err := template.New("path").Parse(ep.path)
	if err != nil {
		return
	}

	t.Execute(&path, vars)
	ep.path = path.String()
}

func (ep *Endpoint) InternalPath(path string) *Endpoint {
	ep.path = path
	return ep
}

func (ep *Endpoint) TransformRequest(ex, in interface{}) *Endpoint {
	ep.reqExtStruct = ex
	ep.reqIntStruct = in
	return ep
}

func (ep *Endpoint) TransformResponse(in, ex interface{}) *Endpoint {
	ep.resIntStruct = in
	ep.resExtStruct = ex
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
